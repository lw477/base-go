package main

import (
	"flag"
	"fmt"
	"os"
)

//func main() {
//	bookworms, err := LoadBookworms("testdata/bookworms.json")
//	if err != nil {
//		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
//		os.Exit(1)
//	}
//
//	common := findCommonBooks(bookworms)
//	fmt.Println("Here are the books in common:")
//	displayBooks(common)
//}

/*
### 作业 1：改造程序，JSON 路径用 flag 命令行参数传入
需求：
1. 复用第 2 章学的 `flag` 包，新增 `-file` 参数，用来指定 json 文件路径。
2. 默认值：`testdata/bookworms.json`，不传参数就读取默认文件。
*/

func main() {
	defaultFilePath := "./testdata/bookworms.json"
	var filePath string
	// 定义命令行 flag：变量指针、参数名、默认值、帮助说明
	flag.StringVar(&filePath, "file", defaultFilePath, "")
	flag.Parse()

	bookworms, err := loadBookworms(filePath)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to load bookworms: %s\n", err)
		os.Exit(1)
	}

	common := findCommonBooks(bookworms)
	fmt.Println("Here are the books in common:")
	displayBooks(common)
}

/*
### 作业 2：新增测试用例
在 `TestLoadBookworms` 测试表里增加 2 个 case：
1. **JSON 文件为空文件**：`testdata/empty.json`（里面什么都不写），预期返回 error。
2. **JSON 格式非法**：`testdata/bad.json`，随便写一段残缺 json，比如 `[{ "name":"Fadi"`，预期返回 error。
>
> 要点：在 testdata 新建两个文件，测试表增加两条，`wantErr:true`。
*/
