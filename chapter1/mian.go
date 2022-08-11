package main

import (
	"fmt"
)

// 程序运行的入口：包名=main，函数名=main
func main() {
	// 调用call uj函数  包名.函数名()
	uj()

	if true == true {
		woAmI("橙橙", 11, true)
	} else {
		woAmI("可馨", 11, false)
	}

}

func uj() {
	// fmt = format（格式化）
	// fmt.Println：换行，ln = line（行、线）
	/*
		111
		222333
	*/
	fmt.Println("111")
	fmt.Print("222")
	fmt.Println("333")
}

// 两个单词相连，第2个单词首字母要用大写
func woAmI(name string, age int, isBoy bool) {
	if isBoy == true {
		// 函数名首字母，必须是大写，才能让，其他包的代码来调用。
		// 跳转到println的代码实现
		fmt.Printf("大家好，我是%s，我今年%d岁，我是男孩。", name, age)
	} else {
		fmt.Printf("大家好，我是%s，我今年%d岁，我是女孩。", name, age)
	}
}
