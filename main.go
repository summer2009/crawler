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

const domainStr string = "weibo.com/ntg7"

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

	fmt.Println(status)
	fmt.Println("*****************************************")

	f, err := os.Open("F:\\demo-project-env\\mygo\\src\\github.com\\bwb\\sentence2132.txt")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	rd := bufio.NewReader(f)

	for {
		line, err := rd.ReadString('\n') //以'\n'为结束符读入一行

		if err != nil || io.EOF == err {
			if line == "" {
				break
			}
		}

		fmt.Println(line)
		index := strings.Index(line, "、")
		var strContent string
		if index != -1 {
			strContent = strings.Replace(SubString(line, index+1, len(line)), "\r\n", "", -1)
		}

		fmt.Println(strContent)
		statusToEs := StatusWeibo{Id: "1", Content: strContent, PicPath: " ", DomainStr: domainStr, TimeOfLastPost: " ", NumOfPost: 0}
		t := reflect.TypeOf(statusToEs)
		v := reflect.ValueOf(statusToEs)
		for k := 0; k < t.NumField(); k++ {
			fmt.Printf("%s -- %v \n", t.Field(k).Name, v.Field(k).Interface())
		}

	}
}
