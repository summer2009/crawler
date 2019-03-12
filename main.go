package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/NextronSystems/go-elasticsearch"
	"github.com/cuixin/csv4g"
)

type WeiboUser struct {
	//"好友名",
	//"好友ID",
	//"性别",
	//"认证",
	//"关注",
	//"粉丝",
	//"微博",
	//"地址",
	//"简介",

	Name     string
	ID       string
	Sex      string
	Idenfied string
	Followed int
	Fans     int
	Weibos   int
	Addr     string
	Intro    string
	//CurLink  string
	//UpLink   string
	//DLTime   string
}

//获取指定目录下的所有文件和目录
func GetFilesAndDirs(dirPth string) (files []string, dirs []string, err error) {
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetFilesAndDirs(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".csv")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	return files, dirs, nil
}

//获取指定目录下的所有文件,包含子目录下的文件
func GetAllFiles(dirPth string) (files []string, err error) {
	var dirs []string
	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	//suffix = strings.ToUpper(suffix) //忽略后缀匹配的大小写

	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth + PthSep + fi.Name())
		} else {
			// 过滤指定格式
			ok := strings.HasSuffix(fi.Name(), ".csv")
			if ok {
				files = append(files, dirPth+PthSep+fi.Name())
			}
		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

func main() {

	documentClient, err := elasticsearch.Open("http://localhost:9200")
	if err != nil {
		panic(err)
	}
	if err := documentClient.Ping(); err != nil {
		panic(err)
	}
	//E:\bwb\golang\src\github.com\henrylee2cn\pholcus\pholcus_pkg\text_out\2019-01-29 072713
	//fileNamePath := "E:\\bwb\\golang\\src\\github.com\\henrylee2cn\\pholcus\\pholcus_pkg\\text_out\\2019-01-31 061434\\微博粉丝列表__aabe406a__好友列表"
	//fileNamePath := "D:\\bwb\\weiboData\\2019-03-07 080603"
	fileNamePath := "D:\\bwb\\golang\\src\\github.com\\henrylee2cn\\pholcus\\pholcus_pkg\\text_out\\2019-03-10 192901"

	/*
		files, dirs, _ := GetFilesAndDirs(fileNamePath)

		for _, dir := range dirs {
			fmt.Printf("获取的文件夹为[%s]\n", dir)
		}

		for _, table := range dirs {
			temp, _, _ := GetFilesAndDirs(table)
			for _, temp1 := range temp {
				files = append(files, temp1)
			}
		}

		for _, table1 := range files {
			fmt.Printf("获取的文件为[%s]\n", table1)
		}

	*/
	fmt.Printf("=======================================\n")
	xfiles, _ := GetAllFiles(fileNamePath)
	for _, file := range xfiles {
		fmt.Printf("获取的文件为[%s]\n", file)
		numOfStart := 0
		csv, err := csv4g.NewWithOpts(file, WeiboUser{}, csv4g.Comma(','), csv4g.LazyQuotes(true), csv4g.SkipLine(numOfStart))
		if err != nil {
			fmt.Printf("Error %v\n", err)
			return
		}
		for i := 0; i < csv.LineLen-20; i = i + 20 {
			//for i := 0; i < 2; i++ {

			/*
				tt := &WeiboUser{}
				err = csv.Parse(tt)
				if err != nil {
					fmt.Printf("Error on parse %v\n", err)
					return
				}
				//fmt.Println("current user:", tt.Name, tt.ID, tt.Sex, tt.Addr, tt.Fans, tt.Followed, tt.Idenfied, tt.Intro)
				document := map[string]interface{}{
					"Name":     tt.Name,
					"ID":       tt.ID,
					"Sex":      tt.Sex,
					"Idenfied": tt.Idenfied,
					"Followed": tt.Followed,
					"Fans":     tt.Fans,
					"Weibos":   tt.Weibos,
					"Addr":     tt.Addr,
					"Intro":    tt.Intro,
				}
				if err := documentClient.InsertDocument("weibouser01", "doc", tt.ID, document, elasticsearch.RefreshTrue); err != nil {
					fmt.Println("could not insert document: %s", err)
				}

			*/

			tt0 := &WeiboUser{}
			err = csv.Parse(tt0)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt1 := &WeiboUser{}
			err = csv.Parse(tt1)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt2 := &WeiboUser{}
			err = csv.Parse(tt2)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt3 := &WeiboUser{}
			err = csv.Parse(tt3)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt4 := &WeiboUser{}
			err = csv.Parse(tt4)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt5 := &WeiboUser{}
			err = csv.Parse(tt5)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt6 := &WeiboUser{}
			err = csv.Parse(tt6)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt7 := &WeiboUser{}
			err = csv.Parse(tt7)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt8 := &WeiboUser{}
			err = csv.Parse(tt8)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt9 := &WeiboUser{}
			err = csv.Parse(tt9)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt10 := &WeiboUser{}
			err = csv.Parse(tt10)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt11 := &WeiboUser{}
			err = csv.Parse(tt11)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt12 := &WeiboUser{}
			err = csv.Parse(tt12)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt13 := &WeiboUser{}
			err = csv.Parse(tt13)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt14 := &WeiboUser{}
			err = csv.Parse(tt14)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt15 := &WeiboUser{}
			err = csv.Parse(tt15)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt16 := &WeiboUser{}
			err = csv.Parse(tt16)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt17 := &WeiboUser{}
			err = csv.Parse(tt17)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt18 := &WeiboUser{}
			err = csv.Parse(tt18)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			tt19 := &WeiboUser{}
			err = csv.Parse(tt19)
			if err != nil {
				fmt.Printf("Error on parse %v\n", err)
				return
			}
			docs := map[string]map[string]interface{}{
				tt0.ID: {
					"Name":     tt0.Name,
					"ID":       tt0.ID,
					"Sex":      tt0.Sex,
					"Idenfied": tt0.Idenfied,
					"Followed": tt0.Followed,
					"Fans":     tt0.Fans,
					"Weibos":   tt0.Weibos,
					"Addr":     tt0.Addr,
					"Intro":    tt0.Intro,
				},
				tt1.ID: {
					"Name":     tt1.Name,
					"ID":       tt1.ID,
					"Sex":      tt1.Sex,
					"Idenfied": tt1.Idenfied,
					"Followed": tt1.Followed,
					"Fans":     tt1.Fans,
					"Weibos":   tt1.Weibos,
					"Addr":     tt1.Addr,
					"Intro":    tt1.Intro,
				},
				tt2.ID: {
					"Name":     tt2.Name,
					"ID":       tt2.ID,
					"Sex":      tt2.Sex,
					"Idenfied": tt2.Idenfied,
					"Followed": tt2.Followed,
					"Fans":     tt2.Fans,
					"Weibos":   tt2.Weibos,
					"Addr":     tt2.Addr,
					"Intro":    tt2.Intro,
				},
				tt3.ID: {
					"Name":     tt3.Name,
					"ID":       tt3.ID,
					"Sex":      tt3.Sex,
					"Idenfied": tt3.Idenfied,
					"Followed": tt3.Followed,
					"Fans":     tt3.Fans,
					"Weibos":   tt3.Weibos,
					"Addr":     tt3.Addr,
					"Intro":    tt3.Intro,
				},
				tt4.ID: {
					"Name":     tt4.Name,
					"ID":       tt4.ID,
					"Sex":      tt4.Sex,
					"Idenfied": tt4.Idenfied,
					"Followed": tt4.Followed,
					"Fans":     tt4.Fans,
					"Weibos":   tt4.Weibos,
					"Addr":     tt4.Addr,
					"Intro":    tt4.Intro,
				},
				tt5.ID: {
					"Name":     tt5.Name,
					"ID":       tt5.ID,
					"Sex":      tt5.Sex,
					"Idenfied": tt5.Idenfied,
					"Followed": tt5.Followed,
					"Fans":     tt5.Fans,
					"Weibos":   tt5.Weibos,
					"Addr":     tt5.Addr,
					"Intro":    tt5.Intro,
				},
				tt6.ID: {
					"Name":     tt6.Name,
					"ID":       tt6.ID,
					"Sex":      tt6.Sex,
					"Idenfied": tt6.Idenfied,
					"Followed": tt6.Followed,
					"Fans":     tt6.Fans,
					"Weibos":   tt6.Weibos,
					"Addr":     tt6.Addr,
					"Intro":    tt6.Intro,
				},
				tt7.ID: {
					"Name":     tt7.Name,
					"ID":       tt7.ID,
					"Sex":      tt7.Sex,
					"Idenfied": tt7.Idenfied,
					"Followed": tt7.Followed,
					"Fans":     tt7.Fans,
					"Weibos":   tt7.Weibos,
					"Addr":     tt7.Addr,
					"Intro":    tt7.Intro,
				},
				tt8.ID: {
					"Name":     tt8.Name,
					"ID":       tt8.ID,
					"Sex":      tt8.Sex,
					"Idenfied": tt8.Idenfied,
					"Followed": tt8.Followed,
					"Fans":     tt8.Fans,
					"Weibos":   tt8.Weibos,
					"Addr":     tt8.Addr,
					"Intro":    tt8.Intro,
				},
				tt9.ID: {
					"Name":     tt9.Name,
					"ID":       tt9.ID,
					"Sex":      tt9.Sex,
					"Idenfied": tt9.Idenfied,
					"Followed": tt9.Followed,
					"Fans":     tt9.Fans,
					"Weibos":   tt9.Weibos,
					"Addr":     tt9.Addr,
					"Intro":    tt9.Intro,
				},
				tt10.ID: {
					"Name":     tt10.Name,
					"ID":       tt10.ID,
					"Sex":      tt10.Sex,
					"Idenfied": tt10.Idenfied,
					"Followed": tt10.Followed,
					"Fans":     tt10.Fans,
					"Weibos":   tt10.Weibos,
					"Addr":     tt10.Addr,
					"Intro":    tt10.Intro,
				},
				tt11.ID: {
					"Name":     tt11.Name,
					"ID":       tt11.ID,
					"Sex":      tt11.Sex,
					"Idenfied": tt11.Idenfied,
					"Followed": tt11.Followed,
					"Fans":     tt11.Fans,
					"Weibos":   tt11.Weibos,
					"Addr":     tt11.Addr,
					"Intro":    tt11.Intro,
				},
				tt12.ID: {
					"Name":     tt12.Name,
					"ID":       tt12.ID,
					"Sex":      tt12.Sex,
					"Idenfied": tt12.Idenfied,
					"Followed": tt12.Followed,
					"Fans":     tt12.Fans,
					"Weibos":   tt12.Weibos,
					"Addr":     tt12.Addr,
					"Intro":    tt12.Intro,
				},
				tt13.ID: {
					"Name":     tt13.Name,
					"ID":       tt13.ID,
					"Sex":      tt13.Sex,
					"Idenfied": tt13.Idenfied,
					"Followed": tt13.Followed,
					"Fans":     tt13.Fans,
					"Weibos":   tt13.Weibos,
					"Addr":     tt13.Addr,
					"Intro":    tt13.Intro,
				},
				tt14.ID: {
					"Name":     tt14.Name,
					"ID":       tt14.ID,
					"Sex":      tt14.Sex,
					"Idenfied": tt14.Idenfied,
					"Followed": tt14.Followed,
					"Fans":     tt14.Fans,
					"Weibos":   tt14.Weibos,
					"Addr":     tt14.Addr,
					"Intro":    tt14.Intro,
				},
				tt15.ID: {
					"Name":     tt15.Name,
					"ID":       tt15.ID,
					"Sex":      tt15.Sex,
					"Idenfied": tt15.Idenfied,
					"Followed": tt15.Followed,
					"Fans":     tt15.Fans,
					"Weibos":   tt15.Weibos,
					"Addr":     tt15.Addr,
					"Intro":    tt15.Intro,
				},
				tt16.ID: {
					"Name":     tt16.Name,
					"ID":       tt16.ID,
					"Sex":      tt16.Sex,
					"Idenfied": tt16.Idenfied,
					"Followed": tt16.Followed,
					"Fans":     tt16.Fans,
					"Weibos":   tt16.Weibos,
					"Addr":     tt16.Addr,
					"Intro":    tt16.Intro,
				},
				tt17.ID: {
					"Name":     tt17.Name,
					"ID":       tt17.ID,
					"Sex":      tt17.Sex,
					"Idenfied": tt17.Idenfied,
					"Followed": tt17.Followed,
					"Fans":     tt17.Fans,
					"Weibos":   tt17.Weibos,
					"Addr":     tt17.Addr,
					"Intro":    tt17.Intro,
				},
				tt18.ID: {
					"Name":     tt18.Name,
					"ID":       tt18.ID,
					"Sex":      tt18.Sex,
					"Idenfied": tt18.Idenfied,
					"Followed": tt18.Followed,
					"Fans":     tt18.Fans,
					"Weibos":   tt18.Weibos,
					"Addr":     tt18.Addr,
					"Intro":    tt18.Intro,
				},
				tt19.ID: {
					"Name":     tt19.Name,
					"ID":       tt19.ID,
					"Sex":      tt19.Sex,
					"Idenfied": tt19.Idenfied,
					"Followed": tt19.Followed,
					"Fans":     tt19.Fans,
					"Weibos":   tt19.Weibos,
					"Addr":     tt19.Addr,
					"Intro":    tt19.Intro,
				},
			}
			if _, err := documentClient.InsertDocuments("weibouser01", "doc", docs); err != nil {
				fmt.Println("could not insert documents: %s", err)
			}
		}
	}
}
