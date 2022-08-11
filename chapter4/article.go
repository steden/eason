package main

import (
	"fmt"
	"time"
)

// 结构类型的作用：可以把一组相关联的属性组合在一起
type News struct {
	title    string    // 标题
	createAt time.Time // 发布时间
	content  string    // 内容
	author   string    // 作者
	src      string    // 来源
	count    int       // 浏览次数
}

// iphone
type iphone struct {
	name     string
	cpu      string
	CameraPx int // 单位：万
	Battery  int // 单位：毫安
}

var phone = iphone{
	name:     "苹果13",
	cpu:      "A15",
	CameraPx: 1200,
	Battery:  5000,
}

var name string
var cpu string
var CameraPx int // 单位：万
var Battery int  // 单位：毫安

//// ShowIphoneArg 显示苹果手机的参数
//func ShowIphoneArg1(name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int, name string, cpu string, CameraPx int, Battery int) {
//
//}

// ShowIphoneArg 显示苹果手机的参数
func ShowIphoneArg2(myPhone iphone) {
	// 逻辑：if true {}
	var arr1 []string
	arr1[0] = ""
	arr1[1] = ""
	arr1[2] = ""
	arr1[3] = ""

	var arr2 []News
	// 结构的值初始化：
	/*
		结构名称 {
			属性名称1:属性值1,
			属性名称2:属性值2,
			属性名称3:属性值3,
		}
	*/
	arr2[0] = News{
		title:    "某个标题",
		createAt: time.Time{},
		content:  "",
		author:   "",
		src:      "",
		count:    0,
	}

	var arr3 []int
	arr3[0] = 123

	// 变量名称 = 值
	// a : 变量名称
	// int：变量类型
	// 123：值
	var a int
	a = 123
	fmt.Println(a)
}
