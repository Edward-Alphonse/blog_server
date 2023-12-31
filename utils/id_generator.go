package utils

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// 因为snowFlake目的是解决分布式下生成唯一id 所以ID中是包含集群和节点编号在内的
// https://segmentfault.com/a/1190000014767902
const (
	numberBits uint8 = 12 // 表示每个集群下的每个节点，1毫秒内可生成的id序号的二进制位 对应上图中的最后一段
	workerBits uint8 = 10 // 每台机器(节点)的ID位数 10位最大可以有2^10=1024个节点数 即每毫秒可生成 2^12-1=4096个唯一ID 对应上图中的倒数第二段
	// 这里求最大值使用了位运算，-1 的二进制表示为 1 的补码，感兴趣的同学可以自己算算试试 -1 ^ (-1 << nodeBits) 这里是不是等于 1023
	workerMax   int64 = -1 ^ (-1 << workerBits) // 节点ID的最大值，用于防止溢出
	numberMax   int64 = -1 ^ (-1 << numberBits) // 同上，用来表示生成id序号的最大值
	timeShift   uint8 = workerBits + numberBits // 时间戳向左的偏移量
	workerShift uint8 = numberBits              // 节点ID向左的偏移量
	// 41位字节作为时间戳数值的话，大约68年就会用完
	// 假如你2010年1月1日开始开发系统 如果不减去2010年1月1日的时间戳 那么白白浪费40年的时间戳啊！
	// 这个一旦定义且开始生成ID后千万不要改了 不然可能会生成相同的ID
	epoch int64 = 1525705533000 // 这个是我在写epoch这个常量时的时间戳(毫秒)
)

type Worker struct {
	mu        sync.Mutex
	timestamp int64
	id        int64
	number    int64
}

var worker *Worker

func init() {
	wr, err := NewWorker(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	worker = wr
}

func NewWorker(workerId int64) (*Worker, error) {
	if workerId < 0 || workerId > workerMax {
		return nil, errors.New(("Worker ID excess of quantity"))
	}
	return &Worker{
		id:        workerId,
		timestamp: 0,
		number:    0,
	}, nil
}

func (w *Worker) GenerateId() int64 {
	w.mu.Lock()
	defer w.mu.Unlock()
	now := time.Now().UnixNano() / 1e6
	if w.timestamp == now {
		w.number++
		if w.number > numberMax {
			for now <= w.timestamp {
				now = time.Now().UnixNano() / 1e6
			}
		}
	} else {
		w.timestamp = now
		w.number = 0
	}
	id := int64((now-epoch)<<timeShift | (w.id << workerShift) | (w.number))
	return id
}

func GenerateId() int64 {
	return worker.GenerateId()
}
