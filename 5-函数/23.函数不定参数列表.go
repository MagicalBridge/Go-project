package main

import "fmt"

// 不定参数的形式更加灵活，相当于一个集合，会给参数添加一个编号，是从0开始算的

func TestSum(args ...int) {
	fmt.Println(args[0])
	fmt.Println(args[1])
	fmt.Println(args[2])
}

// 通过循环的方式，将参数取出来, 可以使用len函数来计算参数存储的数据的长度

func TestSum2(args ...int) {
	for i := 0; i < len(args); i++ {
		fmt.Println(args[i])
	}

	// 数组、切片的遍历，可以使用for range
	for index, value := range args {
		fmt.Println("index", index)
		fmt.Println("value", value)
	}

	// 有时候，我们并不需要第一个index变量，可以使用匿名变量_, 这样程序就知道我舍弃了什么
	for _, value := range args {
		fmt.Println("value", value)
	}

}

// 不定参数可以和确定参数一起使用，需要注意的是确定的参数一定放在前面

func TestSum3(num1 int, args ...int) {
	fmt.Println("num1", num1)
	// 有时候，我们并不需要第一个index变量，可以使用匿名变量_, 这样程序就知道我舍弃了什么
	for _, value := range args {
		fmt.Println("value", value)
	}
}

func main() {
	//TestSum(1, 2, 4)
	//TestSum2(2, 3, 5)
	TestSum3(1, 2, 3)
}

/**
总结：
可以通过下标方式获取参数值
通过循环的方式获取参数值
固定参数放在前面，不定参数放在后面
在对函数进行调用的时候，固定参数必须传值，不定参数可以根据需要来决定是否要传值
*/
