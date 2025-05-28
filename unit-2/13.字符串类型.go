package main

import "fmt"

// 用双引号引用的
// 在go语言中一个汉字占3个字符 打印中文的时候需要注意

/*
*
这个函数 `prefixEnvVars` 的作用是为给定的环境变量名称添加一个前缀，并返回一个包含该环境变量名称的切片（数组）。

### 详细解释：

 1. **函数签名**:
    ```go
    func prefixEnvVars(name string) []string
    ```
    - `prefixEnvVars` 是函数的名称。
    - `name` 是一个字符串参数，表示要添加前缀的环境变量名称。
    - 返回值是一个字符串切片 `[]string`，表示处理后的环境变量名称。

 2. **常量 `evnVarPrefix`**:
    ```go
    const evnVarPrefix = "SIGNATURE"
    ```
    - `evnVarPrefix` 是一个常量，值为 `"SIGNATURE"`。这个常量用作环境变量名称的前缀。

 3. **函数体**:
    ```go
    return []string{evnVarPrefix + "_" + name}
    ```
    - 这里使用了字符串拼接操作：`evnVarPrefix + "_" + name`。
    - `evnVarPrefix` 是 `"SIGNATURE"`，所以拼接后的结果是 `"SIGNATURE_"` 加上 `name` 的值。
    - 例如，如果 `name` 是 `"API_KEY"`，那么拼接后的结果就是 `"SIGNATURE_API_KEY"`。
    - 最终，这个拼接后的字符串被放入一个字符串切片 `[]string` 中，并作为函数的返回值。

### 示例：

假设你调用这个函数：
```go
result := prefixEnvVars("API_KEY")
```
那么 `result` 的值将是：
```go
[]string{"SIGNATURE_API_KEY"}
*/
func prefixEnvVars(name string) []string {
	// 一个应用实例
	const evnVarPrefix = "SIGNATURE"
	res := []string{evnVarPrefix + "_" + name}
	return res
}

func main() {
	var name string = "louis"
	fmt.Println(name)
	fmt.Println(len(name))

	res := prefixEnvVars("louis")
	fmt.Println(res)
}
