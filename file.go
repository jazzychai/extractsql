package main

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// fileName:文件名字(带全路径)
// keyword: 搜索关键词
// resultName: 结果保存文件
func SearchLine(fileName string, keyword string, resultName string) {
	f, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)
	for {
		linecontent, err := rd.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		if strings.Contains(linecontent, keyword) {
			appendToFile(resultName, linecontent)
		}
	}
}

// fileName:文件名字(带全路径)
// content: 写入的内容
func appendToFile(fileName string, content string) error {
	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	return err
}
