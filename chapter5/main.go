package main

import "fmt"

func main() {
	result := Add(1, 5)
	result = Add(result, result)
	fmt.Printf("结果等于：%d\n", result)

	result2 := Compare(5, 8)
	fmt.Print(result2)
}

// 1 + 1 = ?
func Add(left int, right int) int {
	result := left + right
	return result
}

func Compare(left int, right int) bool {
	result := left > right
	return result
}
