package main

import (
	"fmt"
	"time"
)

func main() {
	// 数组的用途：可以把一组同样类型的变量，组合成一个列表
	var arts [3]News // 定义了一个News类型的数组，长度等于3

	arts[0] = News{
		title:    "新闵学校（中学部）图书馆升级改造项目成交公告",
		createAt: time.Date(2022, 7, 13, 0, 0, 0, 0, time.Local),
		content:  "上海市松江区新闵学校《中学部图书馆升级改造项目》项目综合评议采购（竞争性谈判），经过专家认真评审并经采购人确认，本次成交结果公布如下：\n1、成交日期：2022年7月13日\n2、成交供应商：上海普育教育设备有限公司\n3、成交金额：408000元",
		author:   "新闵",
		src:      "未知",
		count:    0,
	}
	arts[1] = News{
		title:    "2021学年第二学期班主任考核优秀公示",
		createAt: time.Date(2022, 7, 13, 0, 0, 0, 0, time.Local),
		content:  "经学校德育考核工作小组综合评估，拟定李文静、李馨佳、韦微、韩愈、胡欢欢、陈文彬、颜舒湘、叶雨、冯敬洋、袁媛、魏祥云、周佳妮、韩慧敏、陈莹、刘珍珍、王昱琪、汪韵婷、孙丽萍、李翠保、郝璇老师为新闵学校2021学年第二学期班主任考核优秀人员。特此公示！\n\n公示时间：2022.7.13-2022.7.17\n\n如有异议，请与校务委员联系反映情况！\n\n联系电话：57683395",
		author:   "学校",
		src:      "未知",
		count:    0,
	}

	arts[2] = News{
		title:    "新闵学校（中学部）门厅改造和团队室项目成交公告",
		createAt: time.Date(2022, 7, 8, 0, 0, 0, 0, time.Local),
		content:  "上海市松江区新闵学校《中学部门厅改造和团队室项目》 项目综合评议采购（竞争性谈判），经过专家认真评审并经采购人确认，本次成交结果公布如下：\n\n1、成交日期：2022年7月8日\n\n2、成交供应商：上海邑源建筑装饰工程有限公司\n\n3、成交金额：373000元",
		author:   "新闵-学校",
		src:      "未知",
		count:    0,
	}

	// 第一次执行i = 0
	// 第二次执行i = 1
	// 第三次执行i = 2
	for i := 0; i < len(arts); i++ {
		fmt.Printf("第%d篇文章：", i+1)
		ShowArticle(arts[i])
		fmt.Println("-------------------------------------")
	}
}

func ShowArticle(news News) {
	if news.author == "新闵学校" {
		return
	}
	fmt.Printf("标题：%s\n", news.title)
	fmt.Printf("发布时间：%s\n", news.createAt)
	fmt.Printf("正文：%s\n\n", news.content)
	fmt.Printf("文章作者：%s\n", news.author)
	fmt.Printf("文章出处：%s\n", news.src)
	fmt.Printf("文章浏览次数：%d\n", news.count)
}
