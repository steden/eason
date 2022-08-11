package school

import "fmt"

func WoAmi(s Student) {
	fmt.Printf("大家好，我是：%s，我今年：%d岁，我的学号是：%d，我的体重是：%fkg，我的数学成绩是：%f，我的英语成绩是：%f。",
		s.Name, s.Age, s.Number, s.Weight, s.MathScores, s.EnglishScores)
}

// 结构类型
// 定义规则：type 类型名称 类型
type Student struct {
	// 名称
	Name string // string 默认""
	// 年龄
	Age int // int默认值：0
	// 学号
	Number int
	// 体重
	Weight float64
	// 数学成绩
	MathScores float64
	// 英语成绩
	EnglishScores float64
}
