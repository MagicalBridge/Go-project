package main

import "fmt"

/**
泛型 Generic Type
泛型函数和泛型类型。
使用了类型参数的函数或类型称为泛型函数或泛型类型.

类型参数 Type Parameter
Go中通过类型参数来支持泛型的。类型参数需要配合类型和函数的定义来一起使用。

相关概念:
- 类型形参，Type parameter，定义时的类型参数。
- 类型约束，Type Constraint，定义类型参数时对参数做的类型上的约束。类型约束其实就是接口interface。
- 实例化，Instantiations，使用类型实参替换类型形参而得到具体类型的过程称为实例化。
- 类型实参，Type Argument，实例化时传递的实际参数

使用场景：
泛型类型，generic type，类型定义中带有类型参数的，称为泛型类型
泛型接收器，接收器类型为泛型类型的，称为泛型接收器。
泛型函数，generic function，函数定义中带有类型参数的，称为泛型函数
泛型接口，接口定义中带有类型参数的，称为泛型接口。
*/

func main() {
	// 这里定义了一个泛型，类型参数T，用于表示 mySlice 中元素的类型
	type mySlice[T int | string] []T

	// 进行泛型的实例化
	intSlice := mySlice[int]{1, 2, 3, 4}
	stringSlice := mySlice[string]{"one", "two"}

	fmt.Println(intSlice)
	fmt.Println(stringSlice)

	// 这里定义一个泛型map, 其中key是string类型，value是float32类型
	type myMap[K string, V float32] map[K]V
	myFloatMap := myMap[string, float32]{
		"one":   1.0,
		"two":   2.0,
		"three": 3.0,
	}
	fmt.Println(myFloatMap)
}
