package consts

var RegisterStatusMap = map[int]string{
	0: "成功",
	1: "用户名或密码为空",
	2: "用户名已注册",
	3: "请稍后再试",
}

var LoginStatusMap = map[int]string{
	0: "成功",
	1: "用户名为空",
	2: "密码为空",
	3: "用户名未注册",
	4: "用户名或密码错误",
	5: "请稍后再试",
}

var ChangePassworStatusMap = map[int]string{
	0: "成功",
	1: "用户名为空",
	2: "旧密码为空",
	3: "新密码为空",
	4: "与当前密码相同",
}
