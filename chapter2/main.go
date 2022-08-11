package main

import "eason/chapter2/school"

func main() {
	// 如果这个类型不在当前包，那么你要调用的话，应该是：包名.类型名 / 包名.函数名()
	// 如果这个类型希望其它包也可以使用，那么你的类型名称和属性名称的首字母，必须大写
	var p school.Student
	p.Name = "橙橙"
	p.Age = 11
	p.Number = 27
	p.Weight = 28.5
	p.MathScores = 98.5
	p.EnglishScores = 95.8

	// 定义变量
	// := 1：定义变量；2：赋值
	p2 := school.Student{
		Name:          "老六",
		Age:           11,
		Number:        27,
		Weight:        28.5,
		MathScores:    98.5,
		EnglishScores: 95.8,
	}

	school.WoAmi(p2)
	school.WoAmi(p2)
}
