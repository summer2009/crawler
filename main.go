package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
)

type StatusWeibo struct {
	Id             string
	Content        string
	PicPath        string
	DomainStr      string
	TimeOfLastPost string
	NumOfPost      int
}

func SubString(str string, begin, length int) string {
	fmt.Println("string to Sub =", str)
	rs := []rune(str)
	lth := len(rs)
	//fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	if begin < 0 {
		begin = 0
	}
	if begin >= lth {
		begin = lth
	}
	end := begin + length

	if end > lth {
		end = lth
	}
	//fmt.Printf("begin=%d, end=%d, lth=%d\n", begin, length, lth)
	return string(rs[begin:end])
}

func main() {
	fmt.Println("本程序用于将文本文件导入ES！")
	fmt.Println("*****************************************")
	fmt.Println("入库格式：")
	status := StatusWeibo{Id: "1", Content: "Good good study!", PicPath: " ", DomainStr: "weibo.com/ntgX", TimeOfLastPost: " ", NumOfPost: 1}
	t := reflect.TypeOf(status)
	v := reflect.ValueOf(status)
	for k := 0; k < t.NumField(); k++ {
		fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
	}
	/*
		fmt.Println("Id:", status.Id)
		fmt.Println("Content:", status.Content)
		fmt.Println("PicPath:", status.PicPath)
		fmt.Println("DomainStr:", status.DomainStr)
		fmt.Println("TimeOfLastPost:", status.TimeOfLastPost)
		fmt.Println("NumOfPost:", status.NumOfPost)
	*/
	fmt.Println(status)

	f, err := os.Open("F:\\demo-project-env\\mygo\\src\\github.com\\bwb\\sentence2132.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	i := 0

	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			if line == "" {
				break
			}
		}

		i++

		if line != "" {
			if i < 100 {
				fmt.Println(line)
				index := strings.Index(line, "、")
				var strContent string
				if index != -1 {
					strContent = SubString(line, index+1, len(line))
				}

				fmt.Println(strContent)
			}

		} else {
			fmt.Println("此处为空行")
		}
	}
}
