package main

import "fmt"

// 运行期间不能改变的量
/**
常量命名规范:
- 只能由数字，字母，(下划线)组成
- 不能以数字开头
- 大写字母和小写字母是不同的，但是建议全部大写
- 不能是关键字
- 见名知意,多个单词组成的名称，使用全大写字母书写，中间使用 (下划线)分隔
例如: const APP_VER="0.7 Beta"
*/
func main() {
	const PI = 3.14
	fmt.Println(PI)
}
