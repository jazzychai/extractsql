package main

import (
	"flag"
	"fmt"
	"os"
)

var filePath, tableName, resultPath string

func init() {
	var help = `从数据库备份文件(.sql)中提取单表数据
https://github.com/jazzychai/extractsql
参数:
	-f 备份文件路径
	-t 需要提取的表名
	-r 结果保存文件(默认为result.sql)
	-h 使用帮助
示例:
	./extractsql -f db.sql -t user -r result_user.sql
	`
	flag.StringVar(&filePath, "f", "", "备份文件路径")
	flag.StringVar(&tableName, "t", "", "需要提取的表名")
	flag.StringVar(&resultPath, "r", "result.sql", "结果保存文件")
	flag.Usage = func() { fmt.Print(help) }
	flag.Parse()

	if filePath == "" && tableName == "" {
		fmt.Print(help)
		os.Exit(0)
	}
	if filePath == "" {
		fmt.Println("请指定备份文件，例如 -f db.sql")
		os.Exit(0)
	}
	if tableName == "" {
		fmt.Println("请指定需要搜索的表名，例如 -t name")
		os.Exit(0)
	}

}

func main() {
	fmt.Println("开始查找")
	SearchLine(filePath, "INSERT INTO `"+tableName+"`", resultPath)
	fmt.Println("查找结束")
}
