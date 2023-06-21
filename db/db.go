package db

import (
	"blog_server/models"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type DataBase struct {
	user     string
	password string
	network  string
	server   string
	port     int
	database string

	db *sql.DB
}

var Instance *DataBase

const tableName = "users"

func init() {
	Instance = NewDefaultDB()
	Instance.setupConnection()
}

func NewDateBase(user, password, network, server string, port int, database string) *DataBase {
	return &DataBase{
		user:     user,
		password: password,
		network:  network,
		server:   server,
		port:     port,
		database: database,
	}
}

func NewDefaultDB() *DataBase {
	return NewDateBase("foreverSun", "123456", "tcp", "localhost", 3306, "demo")
}

func (mysql *DataBase) setupConnection() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", mysql.user, mysql.password, mysql.network, mysql.server, mysql.port, mysql.database)
	DB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Open mysql failed,err:%v\n", err)
		return
	}
	DB.SetConnMaxLifetime(100 * time.Second) //最大连接周期，超过时间的连接就close
	DB.SetMaxOpenConns(100)                  //设置最大连接数
	DB.SetMaxIdleConns(16)                   //设置闲置连接数
	mysql.db = DB
}

func (mysql *DataBase) Query(sql string) (*models.User, error) {
	row := mysql.db.QueryRow(sql)
	user := models.User{}
	if err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email); err != nil {
		fmt.Println("scan failed, err: ", err)
		return nil, err
	}
	return &user, nil
}

func (mysql *DataBase) QueryCount(sql string) (int, error) {
	row := mysql.db.QueryRow(sql)
	count := 0
	if err := row.Scan(&count); err != nil {
		fmt.Println("scan failed, err: ", err)
		return 0, err
	}
	return count, nil
}

func (mysql *DataBase) QueryOne(user *models.User) error {
	sql := fmt.Sprintf("select * from %s", tableName)
	row := mysql.db.QueryRow(sql)
	if err := row.Scan(&user.Id, &user.Name, &user.Password, &user.Email); err != nil {
		fmt.Println("scan failed, err: ", err)
		return err
	}
	return nil
}

func (mysql *DataBase) QueryMulti() []*models.User {
	sql := fmt.Sprintf("select * from %s where id > ?", tableName)
	rows, err := mysql.db.Query(sql, 0)
	defer func() {
		if rows != nil {
			rows.Close()
		}
	}()
	users := make([]*models.User, 0)
	if err != nil {
		fmt.Printf("Query failed,err:%v", err)
		return users
	}
	for rows.Next() {
		user := &models.User{}
		if err = rows.Scan(user.Id, user.Name); err != nil {
			fmt.Println("")
			continue
		}
		users = append(users, user)
	}
	return users
}

func (mysql *DataBase) Insert(sql string) error {
	// sql := fmt.Sprintf("insert INTO %s values(?,?,?,?)", tableName)
	fmt.Println("----------", sql)
	result, err := mysql.db.Exec(sql)
	if err != nil {
		fmt.Printf("Insert failed, err:%v", err)
		return err
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return err
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected() //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return err
	}
	fmt.Println("RowsAffected:", rowsaffected)
	return nil
}

func (mysql *DataBase) InsertUser(user models.User) error {
	sql := fmt.Sprintf("insert INTO %s values(?,?,?,?)", tableName)
	result, err := mysql.db.Exec(sql, user.Id, user.Name, user.Password, user.Email)
	if err != nil {
		fmt.Printf("Insert failed,err:%v", err)
		return err
	}
	lastInsertID, err := result.LastInsertId() //插入数据的主键id
	if err != nil {
		fmt.Printf("Get lastInsertID failed,err:%v", err)
		return err
	}
	fmt.Println("LastInsertID:", lastInsertID)
	rowsaffected, err := result.RowsAffected() //影响行数
	if err != nil {
		fmt.Printf("Get RowsAffected failed,err:%v", err)
		return err
	}
	fmt.Println("RowsAffected:", rowsaffected)
	return nil
}
