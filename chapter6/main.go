package main

import "fmt"

// 实现一个用户注册功能
// 要求1：需要填写：用户名、密码、确认密码、年龄、学号
// 要求2：需要判断所有的资料不能为空
// 要求3：需要判断两次密码必须一样
// 要求4：密码长度必须大于6位
// 要求5：注册成功后，将注册的信息返回出来，然后打印：恭喜xxx，注册成功！
func main() {
	name, _, _, _ := Reg("eason", "1234", "1234", 10, 27)
	fmt.Printf("恭喜" + name + "注册成功！")
}

// Reg 注册功能
func Reg(name string, pwd string, pwd2 string, age int, number int) (string, string, int, int) {
	//var name string
	//var pwd string
	//var pwd2 string
	//var age int
	//var number int
	if name == "" {
		fmt.Print("资料不能为空")
	}

	if pwd == "" {
		fmt.Printf("资料不能为空")
	}

	if pwd2 == "" {
		fmt.Printf("资料不能为空")
	}

	if age < 1 || age > 100 {
		fmt.Printf("资料不能为空")
	}

	if number < 0 || number > 52 {
		fmt.Printf("资料不能为空")
	}

	if pwd != pwd2 {
		fmt.Printf("两次密码必须一样")
	}
	if len(pwd) <= 6 {
		fmt.Printf("密码长度必须大于6位")
	}

	// 保存到数据库
	saveDatabase(name, pwd, age, number)

	return name, pwd, age, number
}

// 保存到数据库
func saveDatabase(arg ...any) bool {
	return true
}
