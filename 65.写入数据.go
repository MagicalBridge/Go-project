package main

import (
	"fmt"
	"os"
)

func main001() {
	// 这里使用的是相对路径
	fp, err := os.Create("./a.txt")

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	// 延迟关闭文件
	defer fp.Close()

	// 将字符串写入到文件中
	fp.WriteString("hello world\r\n")
	fp.WriteString("性感荷官在线发牌")
}

func main002() {
	fp, err := os.Create("./b.txt")

	if err != nil {
		fmt.Println("文件创建失败")
		return
	}

	defer fp.Close()
	var str = "性感法师在线讲课" // 一个汉字占用的字节数据是3

	// Write接收的是一个byte类型的数据
	count, err := fp.Write([]byte(str))

	if err != nil {
		fmt.Println("写入文件失败")
		return
	} else {
		fmt.Println(count)
	}
}

func main() {
	//fp, err := os.Create("D:/a.txt")
	//os.Open(文件名) 只读方式打开文件
	//os.OpenFile(文件名，打开方式，打开权限)    如果文件不存在会报错
	//打开方式 os.O_RDONLY（只读方式打开）os.O_WRONLY（只写方式打开）os.O_RDWR（可读可写方式打开）os.O_APPEND（追加方式打开）
	//打开权限 0-7 rwx  6 （rw-）读写权限  7 （rwx）读写执行权限
	fp, err := os.Create("./c.txt")

	if err != nil {
		fmt.Println("打开文件失败")
		return
	}

	defer fp.Close()

	fp.WriteString("hello world")

	// 获取光标流位置
	// 获取从文件起始到结尾有多少个字符
	// 指定位置写入
	// fp.WriteAt([]byte("hello world"),count)
	fp.WriteAt([]byte("itcast"), 0)
}
