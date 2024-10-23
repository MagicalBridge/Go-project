package main

import (
	"fmt"
	"time"
)

// 返回值可以返回单个值，也可以返回多个值 返回值类型也需要指定
func addResult(num1 int, num2 int) int {
	var sum int
	sum = num1 + num2
	return sum
}

// GetResult 函数还可以返回多个值，写法如下：
func GetResult() (int, int) {
	var num1 = 10
	var num2 = 20
	return num1, num2
}

// Callback 1、回调函数的应用场景 定义一个回调函数类型
type Callback func()

// 模拟一个异步操作，完成后调用回调函数
func asyncOperation(callback Callback) {
	// 模拟异步操作
	time.Sleep(2 * time.Second)
	// 操作完成后调用回调函数
	callback()
}

// Strategy 2、定义一个策略函数类型 策略模式 根据不同的策略（即不同的函数）来执行不同的行为
type Strategy func(int, int) int

// 加法策略
func addStrategy(a, b int) int {
	return a + b
}

// 减法策略
func subtract(a, b int) int {
	return a - b
}

// 使用策略的函数
func executeStrategy(strategy Strategy, a, b int) int {
	return strategy(a, b)
}

// FilterFunc 3、过滤器和映射器，定义一个过滤函数类型
type FilterFunc func(int) bool

// MapFunc 定义一个映射函数类型
type MapFunc func(int) int

// 过滤函数
func filter(numbers []int, f FilterFunc) []int {
	var result []int
	for _, num := range numbers {
		if f(num) {
			result = append(result, num)
		}
	}
	return result
}

// 映射函数
func mapFunc(numbers []int, f MapFunc) []int {
	var result []int
	for _, num := range numbers {
		result = append(result, f(num))
	}
	return result
}

func main() {
	var s = addResult(1, 3)
	fmt.Println(s)

	var s1, s2 = GetResult()
	fmt.Println(s1)
	fmt.Println(s2)

	// 1、回调函数的演示场景
	fmt.Println("开始异步操作")
	asyncOperation(func() {
		fmt.Println("异步操作完成")
	})
	fmt.Println("异步操作已启动")

	// 2、策略模式
	result1 := executeStrategy(addStrategy, 10, 5)
	fmt.Println("10 + 5 =", result1) // 输出: 10 + 5 = 15

	result2 := executeStrategy(subtract, 10, 5)
	fmt.Println("10 - 5 =", result2) // 输出: 10 - 5 = 5

	// 3、过滤器
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// 过滤出偶数
	evenNumbers := filter(numbers, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println("偶数:", evenNumbers) // 输出: 偶数: [2 4 6 8 10]
	// 将每个数平方
	squaredNumbers := mapFunc(numbers, func(n int) int {
		return n * n
	})
	fmt.Println("平方数:", squaredNumbers) // 输出: 平方数: [1 4 9 16 25 36 49 64 81 100]
}
