package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hellogo"
	// Contains的方法判断字符串中是否包含指定子字符串，如果包含则返回true，否则返回false
	b := strings.Contains(str, "go")
	fmt.Println(b)

	strJoin := []string{"hello", "world"}
	// Join方法将字符串切片的各字符串连接成一个字符串，连接时使用sep字符串作为分隔符
	strJoinRes := strings.Join(strJoin, "|")
	fmt.Println(strJoinRes)

	strIndex := "helloWeb3"
	// Index方法返回的是一个int类型的变量
	strIndexRes := strings.Index(strIndex, "Web")
	fmt.Println(strIndexRes)
}
