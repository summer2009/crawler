package pholcus_lib

// 基础包
import (
	"github.com/henrylee2cn/pholcus/app/downloader/request" //必需
	. "github.com/henrylee2cn/pholcus/app/spider"           //必需

	//. "github.com/henrylee2cn/pholcus/app/spider/common"    //选用
	"github.com/henrylee2cn/pholcus/common/goquery" //DOM解析
	"github.com/henrylee2cn/pholcus/logs"           //信息输出

	// net包
	"net/http" //设置http.Header
	// "net/url"

	// 编码包
	// "encoding/xml"
	// "encoding/json"

	// 字符串处理包
	// "regexp"
	"strconv"
	"strings"

	// 其他包
	"fmt"
	// "time"
	// "io/ioutil"
	"github.com/cuixin/csv4g"
)

var cookieStr = ""
var glevel = 0
var curLine = 0

func init() {
	WeiboFans.Register()
}

var WeiboFans = &Spider{
	Name:         "微博粉丝列表",
	Description:  `新浪微博粉丝 [自定义输入格式 "Level"::"filePathName"::"startLineNum"::"ID"::"Cookie"][最多支持250页，内设定时1~2s]`,
	Pausetime:    2000,
	Keyin:        KEYIN,
	Limit:        LIMIT,
	EnableCookie: true,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			param := strings.Split(ctx.GetKeyin(), "::")
			if len(param) != 5 {
				logs.Log.Error("自定义输入的参数不正确！")
				return
			}

			level, err := strconv.Atoi(strings.Trim(param[0], " "))
			glevel = level
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("level=", level)
			filePathName := strings.Trim(param[1], " ")
			startLineNum := strings.Trim(param[2], " ")

			fmt.Println(filePathName, startLineNum)

			id := strings.Trim(param[3], " ")
			cookie := strings.Trim(param[4], " ")

			cookieStr = cookie

			var count1 = 0
			var count2 = 0
			if ctx.GetLimit() < count1 {
				count1 = ctx.GetLimit()
			}
			if ctx.GetLimit() < count2 {
				count2 = ctx.GetLimit()
			}

			//如果是Level = 0
			if level == 0 {
				//获取用户信息
				//https://weibo.com/6606483016/profile?topnav=1&wvr=6&is_all=1
				ctx.AddQueue(&request.Request{
					Url:          "https://weibo.com/" + id + "/profile?topnav=1&wvr=6&is_all=1",
					Rule:         "用户资料",
					Header:       http.Header{"Cookie": []string{cookie}},
					DownloaderID: 0,
				})
			} else {
				//level != 0 则逐行读取文件中的uid
				if filePathName != "" {
					// /*
					numOfStart := 0
					if startLineNum == " " {
						numOfStart = 0
					} else {
						numOfStart, _ = strconv.Atoi(startLineNum)
					}

					fmt.Println("startLineNum=", numOfStart)
					//打开文件,逐行读取文件中的uid
					//filePathName = strings.Replace(filePathName,"/\","\\",-1)

					csv, err := csv4g.NewWithOpts("d:\\3764.csv", WeiboUser{}, csv4g.Comma(','), csv4g.LazyQuotes(true), csv4g.SkipLine(numOfStart))
					if err != nil {
						fmt.Printf("Error %v\n", err)
						return
					}
					for i := 0; i < csv.LineLen; i++ {
						tt := &WeiboUser{}
						err = csv.Parse(tt)
						if err != nil {
							fmt.Printf("Error on parse %v\n", err)
							return
						}
						fmt.Println("current line:", i)
						curLine = i
						if glevel == 0 {
							//获取用户信息
							//https://weibo.com/6606483016/profile?topnav=1&wvr=6&is_all=1

							ctx.AddQueue(&request.Request{
								Url:          "https://weibo.com/" + tt.ID + "/profile?topnav=1&wvr=6&is_all=1",
								Rule:         "用户资料",
								Header:       http.Header{"Cookie": []string{cookie}},
								DownloaderID: 0,
							})

						} else {
							//无需获取用户信息，直接从文件中读取用户id，fans，follows数据
							countOfFollows := tt.Followed/30 + 1
							countOfFans := tt.Fans/20 + 1
							if countOfFollows > 5 {
								countOfFollows = 5
							}
							if countOfFans > 5 {
								countOfFans = 5
							}

							//获取粉丝列表
							//for i := count1; i > 0; i-- {
							for i := 1; i <= countOfFans; i++ {
								ctx.AddQueue(&request.Request{
									//https://weibo.com/p/1005056850846383/follow?relate=fans&page=2#Pl_Official_HisRelation__59
									Url:          "https://weibo.com/p/100505" + tt.ID + "/follow?relate=fans&page=" + strconv.Itoa(i) + "#Pl_Official_HisRelation__59",
									Rule:         "好友列表",
									Header:       http.Header{"Cookie": []string{cookieStr}},
									DownloaderID: 0,
								})
							}

							/*
								//获取关注列表
								for i := 1; i <= countOfFollows; i++ {
									ctx.AddQueue(&request.Request{
										//https://weibo.com/p/1005056850846383/follow?page=5#Pl_Official_HisRelation__59
										Url:          "https://weibo.com/p/100505" + tt.ID + "/follow?page=" + strconv.Itoa(i) + "#Pl_Official_HisRelation__59",
										Rule:         "关注列表",
										Header:       http.Header{"Cookie": []string{cookieStr}},
										DownloaderID: 0,
									})
								}
							*/
							/*
								//https://weibo.com/6606483016/profile?topnav=1&wvr=6&is_all=1
								//https://weibo.com/u/6850846383?refer_flag=1005050005_&is_hot=1
								ctx.AddQueue(&request.Request{
									Url:          "https://weibo.com/" + tt.ID + "/profile?topnav=1&wvr=6&is_all=1",
									Rule:         "用户资料",
									Header:       http.Header{"Cookie": []string{cookie}},
									DownloaderID: 0,
								})
							*/

						}

						//https://weibo.com/u/6850846383?refer_flag=1005050005_&is_hot=1
					}

					/*
						//获取用户信息
						//https://weibo.com/6606483016/profile?topnav=1&wvr=6&is_all=1
						ctx.AddQueue(&request.Request{
							Url:          "https://weibo.com/" + strconv.Itoa(tt.Id) + "/profile?topnav=1&wvr=6&is_all=1",
							Rule:         "用户资料",
							Header:       http.Header{"Cookie": []string{cookie}},
							DownloaderID: 0,
						})
					*/

					//*/

				}

			}

		},

		Trunk: map[string]*Rule{
			"好友列表": {
				ItemFields: []string{
					"Name",
					"ID",
					"Sex",
					"Idenfied",
					"Followed",
					"Fans",
					"Weibos",
					"Addr",
					"Intro",

					/*
						"好友名",
						"好友ID",
						"性别",
						"认证",
						"关注",
						"粉丝",
						"微博",
						"地址",
						"简介",
					*/
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					var 属性 map[string]string
					fmt.Println(query.Find("title").Text())
					//fmt.Println(query.Text())
					//fmt.Println(query.Find("script").Text())
					query.Find("script").Each(func(i int, s *goquery.Selection) {
						//fmt.Println("search in script for 好友列表")
						//fmt.Println(s.Text())

						strTmp := s.Text()
						if strings.Contains(strTmp, "follow_list") {
							fmt.Println("work in 好友列表")
							slice0 := strings.Split(strTmp, "\":\"")
							for i, v := range slice0 {
								if strings.Contains(v, "follow_list") {
									strTmp02 := strings.Replace(v, "\\n", "", -1)
									strTmp03 := strings.Replace(strTmp02, "\\r", "", -1)
									strTmp04 := strings.Replace(strTmp03, "\\", "", -1)
									strTmp05 := strings.Replace(strTmp04, "\"})", "", -1)
									fmt.Printf("下标: %d \n", i)
									doc, err := goquery.NewDocumentFromReader(strings.NewReader(strTmp05))
									if err != nil {
										fmt.Println(err)
									}
									doc.Find(".follow_item").Each(func(j int, sx *goquery.Selection) {
										//	fmt.Println("%d fans founded!", j+1)
										/*
											name, _ := sx.Find(".info_name a").Attr("title")
											fmt.Println(name)
											url, _ := sx.Find(".info_name a").Attr("href")
											uid := strings.Replace(url, "/u", "", -1)
											uid = strings.Replace(uid, "/", "", -1)
											url = "https://weibo.com/p/100505" + uid + "/info?mod=pedit_more"
											fmt.Println(uid)
										*/

										actionData, _ := sx.Attr("action-data")

										dataTmp := strings.Split(actionData, "&")
										name := strings.Replace(dataTmp[1], "fnick=", "", -1)
										uid := strings.Replace(dataTmp[0], "uid=", "", -1)

										var 认证 string = ""
										if _, isExist := sx.Find(".info_name i").Attr("title"); isExist {
											认证 = "认证"

										}
										var 性别 string = ""
										sex, _ := sx.Find("a i").Attr("class")
										if sex == "W_icon icon_female" {
											性别 = "female"
										} else {
											性别 = "male"
										}

										关注 := sx.Find(".info_connect em a").Eq(0).Text()
										粉丝 := sx.Find(".info_connect em a").Eq(1).Text()
										微博 := sx.Find(".info_connect em a").Eq(2).Text()
										adr := sx.Find(".info_add").Eq(0).Text()
										地址 := strings.Replace(adr, "地址", "", -1)
										地址 = strings.Replace(地址, "t", "", -1)
										简介 := sx.Find(".info_intro span").Eq(0).Text()
										uidTmp := strings.Split(uid, "?")
										uidNew := uidTmp[0]
										fmt.Println(关注, 粉丝, 微博, 地址)
										fmt.Println("process line:", curLine)
										if 属性 == nil {
											属性 = map[string]string{}
										}
										//20190112
										属性["Name"] = name
										属性["ID"] = uidNew
										属性["Sex"] = 性别
										属性["Followed"] = 关注
										属性["Fans"] = 粉丝
										属性["Weibos"] = 微博
										属性["Idenfied"] = 认证
										属性["Addr"] = 地址
										属性["Intro"] = 简介

										结果 := map[int]interface{}{
											0: name,
											1: uidNew,
											2: 性别,
											3: 认证,
											4: 关注,
											5: 粉丝,
											6: 微博,
											7: 地址,
											8: 简介,
										}
										for k, v := range 属性 {
											idx := ctx.UpsertItemField(k)
											结果[idx] = v
										}
										fmt.Println(结果)
										// 结果输出
										ctx.Output(结果)

										/*
											x := &request.Request{
												Url:          url,
												Rule:         "好友资料",
												DownloaderID: 0,
												Temp: map[string]interface{}{
													"好友名":  name,
													"好友ID": uidNew,
													"性别":   性别,
													"认证":   认证,
													"关注":   关注,
													"粉丝":   粉丝,
													"微博":   微博,
													"地址":   地址,
													"简介":   简介,
												},
											}
											ctx.AddQueue(x)

											//根据粉丝uid，粉丝数量，产生获取粉丝的粉丝的队列信息
											url_1 := "https://weibo.com/p/100505" + uidNew + "/follow?relate=fans&from=100505&wvr=6&mod=headfans&current=fans#place"
											strTmp := "SINAGLOBAL=2628254495597.748.1520806133490; wb_cmtLike_3087607255=1; UOR=,,login.sina.com.cn; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9W533e8L5PBZ.qFocT7ZKVYG5JpX5KMhUgL.Foe71hMcehMESK-2dJLoI0qLxKqLB-qL12qLxKqL1KMLBK-LxKqLBK5LB.eLxKnLBKnL1h5LxKnL12zLBo2LxKnLB-qLB-Bt; wvr=6; wb_view_log_3087607255=1600*9001; Ugrow-G0=ea90f703b7694b74b62d38420b5273df; ALF=1578794249; SSOLoginState=1547258249; SCF=Aumi17g3by3yDYXtGGP-nY6seE2pBVJAZkeeOOdrRIMXxRKt6-950aVCW_Q-wZLjb5sPawilzCunU7tYLHh8f5g.; SUB=_2A25xPTnZDeRhGeVO41UX8CnOzjmIHXVSSywRrDV8PUNbmtAKLVrckW9NTX0uk5gkCZc3139iPY1IK6qDXbmDhgGg; SUHB=0G0erSN-3zayCF; TC-V5-G0=40eeee30be4a1418bde327baf365fcc0; TC-Page-G0=cdcf495cbaea129529aa606e7629fea7; _s_tentry=login.sina.com.cn; Apache=37805710225.4492.1547258253538; ULV=1547258253606:10:1:1:37805710225.4492.1547258253538:1546002408107"
											x1 := &request.Request{
												Url:          url_1,
												Rule:         "粉丝列表",
												Header:       http.Header{"Cookie": []string{strTmp}},
												DownloaderID: 0,
												Temp: map[string]interface{}{
													"好友名":  name,
													"好友ID": uidNew,
													"认证":   认证,
													"关注":   关注,
													"粉丝":   粉丝,
													"微博":   微博,
													"地址":   地址,
												},
											}
											ctx.AddQueue(x1)

											//https://weibo.com/p/1005056880196455/follow?page=3#Pl_Official_HisRelation__59
											//url_1 := "https://weibo.com/p/100505" + uidNew + "/follow?relate=fans&from=100505&wvr=6&mod=headfans&current=fans#place"
											for i := 2; i < 11; i++ {
												ctx.AddQueue(&request.Request{
													Url:          "https://weibo.com/" + uidNew + "/follow?page=" + strconv.Itoa(i) + "#Pl_Official_HisRelation__59",
													Rule:         "粉丝列表",
													Header:       http.Header{"Cookie": []string{strTmp}},
													DownloaderID: 0,
													Temp: map[string]interface{}{
														"好友名":  name,
														"好友ID": uidNew,
														"认证":   认证,
														"关注":   关注,
														"粉丝":   粉丝,
														"微博":   微博,
														"地址":   地址,
													},
												})
											}

										*/
									})

								}
							}
							//fmt.Println(slice0)
						}
					})
				},
			},
			"关注列表": {
				ItemFields: []string{
					"Name",
					"ID",
					"Sex",
					"Idenfied",
					"Followed",
					"Fans",
					"Weibos",
					"Addr",
					"Intro",
					/*
						"好友名",
						"好友ID",
						"性别",
						"认证",
						"关注",
						"粉丝",
						"微博",
						"地址",
						"简介",
					*/
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					var 属性 map[string]string
					fmt.Println(query.Find("title").Text())
					//fmt.Println(query.Text())
					//fmt.Println(query.Find("script").Text())
					query.Find("script").Each(func(i int, s *goquery.Selection) {
						fmt.Println("search in script for 关注列表")
						//fmt.Println(s.Text())

						strTmp := s.Text()
						if strings.Contains(strTmp, "pl.relation.myFollow.index") {
							fmt.Println("work in 关注列表")
							slice0 := strings.Split(strTmp, "\":\"")
							for i, v := range slice0 {
								if strings.Contains(v, "WB_search_s") {
									strTmp02 := strings.Replace(v, "\\n", "", -1)
									strTmp03 := strings.Replace(strTmp02, "\\r", "", -1)
									strTmp04 := strings.Replace(strTmp03, "\\", "", -1)
									strTmp05 := strings.Replace(strTmp04, "\"})", "", -1)
									fmt.Printf("下标: %d \n", i)
									//fmt.Println(strTmp05)
									doc, err := goquery.NewDocumentFromReader(strings.NewReader(strTmp05))
									if err != nil {
										fmt.Println(err)
									}
									doc.Find(".member_li").Each(func(j int, sx *goquery.Selection) {
										fmt.Println("%i fans founded!", j+1)
										//fmt.Println(sx.Text())
										actionData, _ := sx.Find(".W_btn_b").Attr("action-data")

										dataTmp := strings.Split(actionData, "&")
										name := strings.Replace(dataTmp[1], "nick=", "", -1)
										uid := strings.Replace(dataTmp[2], "uid=", "", -1)
										fmt.Println(uid)
										var 认证 string = ""
										if _, isExist := sx.Find(".info_name i").Attr("title"); isExist {
											认证 = "认证"
										}
										var 性别 string = ""
										性别 = strings.Replace(dataTmp[3], "sex=", "", -1)
										关注 := ""
										粉丝 := ""
										微博 := ""
										地址 := ""
										简介 := sx.Find("div .text").Eq(0).Text()
										fmt.Println(关注, 粉丝, 微博, 地址)
										if 属性 == nil {
											属性 = map[string]string{}
										}
										//20190112

										属性["Name"] = name
										属性["ID"] = uid
										属性["Sex"] = 性别
										属性["Followed"] = 关注
										属性["Fans"] = 粉丝
										属性["Weibos"] = 微博
										属性["Idenfied"] = 认证
										属性["Addr"] = 地址
										属性["Intro"] = 简介

										结果 := map[int]interface{}{
											0: name,
											1: uid,
											2: 性别,
											3: 认证,
											4: 关注,
											5: 粉丝,
											6: 微博,
											7: 地址,
											8: 简介,
										}
										for k, v := range 属性 {
											idx := ctx.UpsertItemField(k)
											结果[idx] = v
										}
										fmt.Println(结果)
										// 结果输出
										ctx.Output(结果)

										/*
											x := &request.Request{
												Url:          url,
												Rule:         "好友资料",
												DownloaderID: 0,
												Temp: map[string]interface{}{
													"好友名":  name,
													"好友ID": uidNew,
													"性别":   性别,
													"认证":   认证,
													"关注":   关注,
													"粉丝":   粉丝,
													"微博":   微博,
													"地址":   地址,
													"简介":   简介,
												},
											}
											ctx.AddQueue(x)

											//根据粉丝uid，粉丝数量，产生获取粉丝的粉丝的队列信息
											url_1 := "https://weibo.com/p/100505" + uidNew + "/follow?relate=fans&from=100505&wvr=6&mod=headfans&current=fans#place"
											strTmp := "SINAGLOBAL=2628254495597.748.1520806133490; wb_cmtLike_3087607255=1; UOR=,,login.sina.com.cn; SUBP=0033WrSXqPxfM725Ws9jqgMF55529P9D9W533e8L5PBZ.qFocT7ZKVYG5JpX5KMhUgL.Foe71hMcehMESK-2dJLoI0qLxKqLB-qL12qLxKqL1KMLBK-LxKqLBK5LB.eLxKnLBKnL1h5LxKnL12zLBo2LxKnLB-qLB-Bt; wvr=6; wb_view_log_3087607255=1600*9001; Ugrow-G0=ea90f703b7694b74b62d38420b5273df; ALF=1578794249; SSOLoginState=1547258249; SCF=Aumi17g3by3yDYXtGGP-nY6seE2pBVJAZkeeOOdrRIMXxRKt6-950aVCW_Q-wZLjb5sPawilzCunU7tYLHh8f5g.; SUB=_2A25xPTnZDeRhGeVO41UX8CnOzjmIHXVSSywRrDV8PUNbmtAKLVrckW9NTX0uk5gkCZc3139iPY1IK6qDXbmDhgGg; SUHB=0G0erSN-3zayCF; TC-V5-G0=40eeee30be4a1418bde327baf365fcc0; TC-Page-G0=cdcf495cbaea129529aa606e7629fea7; _s_tentry=login.sina.com.cn; Apache=37805710225.4492.1547258253538; ULV=1547258253606:10:1:1:37805710225.4492.1547258253538:1546002408107"
											x1 := &request.Request{
												Url:          url_1,
												Rule:         "粉丝列表",
												Header:       http.Header{"Cookie": []string{strTmp}},
												DownloaderID: 0,
												Temp: map[string]interface{}{
													"好友名":  name,
													"好友ID": uidNew,
													"认证":   认证,
													"关注":   关注,
													"粉丝":   粉丝,
													"微博":   微博,
													"地址":   地址,
												},
											}
											ctx.AddQueue(x1)

											//https://weibo.com/p/1005056880196455/follow?page=3#Pl_Official_HisRelation__59
											//url_1 := "https://weibo.com/p/100505" + uidNew + "/follow?relate=fans&from=100505&wvr=6&mod=headfans&current=fans#place"
											for i := 2; i < 11; i++ {
												ctx.AddQueue(&request.Request{
													Url:          "https://weibo.com/" + uidNew + "/follow?page=" + strconv.Itoa(i) + "#Pl_Official_HisRelation__59",
													Rule:         "粉丝列表",
													Header:       http.Header{"Cookie": []string{strTmp}},
													DownloaderID: 0,
													Temp: map[string]interface{}{
														"好友名":  name,
														"好友ID": uidNew,
														"认证":   认证,
														"关注":   关注,
														"粉丝":   粉丝,
														"微博":   微博,
														"地址":   地址,
													},
												})
											}

										*/
									})

								}
							}
							//fmt.Println(slice0)
						}
					})
				},
			},
			"粉丝列表": {
				ItemFields: []string{
					"Name",
					"ID",
					"Sex",
					"Idenfied",
					"Followed",
					"Fans",
					"Weibos",
					"Addr",
					"Intro",
					/*
						"好友名",
						"好友ID",
						"性别",
						"认证",
						"关注",
						"粉丝",
						"微博",
						"地址",
						"简介",
					*/
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					var 属性 map[string]string
					//var title string
					//var detail string
					fmt.Println(query.Find("title").Text())
					fmt.Println("work in 粉丝列表")
					//fmt.Println(query.Find("script").Text())
					query.Find("script").Each(func(i int, s *goquery.Selection) {
						//fmt.Println("search in script for 好友列表")
						//fmt.Println(s.Text())

						strTmp := s.Text()
						//if strings.Contains(strTmp, "pl.content.followTab.index") {
						if strings.Contains(strTmp, "follow_list") {
							fmt.Println("follow_list")
							slice0 := strings.Split(strTmp, "\":\"")
							for i, v := range slice0 {
								if strings.Contains(v, "li class") {
									strTmp02 := strings.Replace(v, "\\n", "", -1)
									strTmp03 := strings.Replace(strTmp02, "\\r", "", -1)
									strTmp04 := strings.Replace(strTmp03, "\\", "", -1)
									strTmp05 := strings.Replace(strTmp04, "\"})", "", -1)
									fmt.Printf("下标: %d \n", i)
									//fmt.Println(strTmp05)
									doc, err := goquery.NewDocumentFromReader(strings.NewReader(strTmp05))
									if err != nil {
										fmt.Println(err)
									}

									doc.Find(".follow_item").Each(func(j int, sx *goquery.Selection) {
										fmt.Println("work in 粉丝列表 %d fans founded!", j+1)

										userInfo, _ := sx.Attr("action-data")
										fmt.Println(userInfo)
										//uid=3194212763&fnick=凤翅紫金冠f&sex=m
										//从userInfo中提取uid、fnick、sex
										slice1 := strings.Split(userInfo, "&")
										uidX := strings.Split(slice1[0], "=")
										fnick := strings.Split(slice1[1], "=")
										sex := strings.Split(slice1[2], "=")
										fmt.Println(uidX[1], fnick[1], sex[1])

										name := sx.Find(".info_name a").Text()
										fmt.Println(name)
										url, _ := sx.Find(".info_name a").Attr("href")
										fmt.Println(url)
										uid := strings.Replace(url, "/u", "", -1)
										uid = strings.Replace(uid, "/", "", -1)
										url = "https://weibo.com/p/100505" + uid + "/info?mod=pedit_more"
										var 认证 string = ""
										if _, isExist := sx.Find(".info_name i").Attr("title"); isExist {
											认证 = "认证"

										}
										关注 := sx.Find(".info_connect em a").Eq(0).Text()
										粉丝 := sx.Find(".info_connect em a").Eq(1).Text()
										微博 := sx.Find(".info_connect em a").Eq(2).Text()
										地址 := sx.Find(".info_add span").Eq(0).Text()
										简介 := sx.Find(".info_intro span").Eq(0).Text()
										fmt.Println(关注, 粉丝, 微博, uid, 简介)
										//uidNew := strings.Replace(uid, "?refer_flag=1005050006_", "", -1)
										uidNew := strings.TrimLeft(uid, "?")
										fmt.Println(uidNew)
										if 属性 == nil {
											属性 = map[string]string{}
										}
										//title = s.Find(".pt_title").Text()
										//title = Deprive2(title)
										//detail = s.Find(".pt_detail").Text()
										//detail = Deprive2(detail)
										属性["好友名"] = name
										属性["好友id"] = uidX[1]
										属性["性别"] = sex[1]
										属性["关注"] = 关注
										属性["粉丝"] = 粉丝
										属性["微博"] = 微博
										属性["认证"] = 认证
										属性["地址"] = 地址
										属性["简介"] = 简介

										结果 := map[int]interface{}{
											0: name,
											1: uidX[1],
											2: sex[1],
											3: 认证,
											4: 关注,
											5: 粉丝,
											6: 微博,
											7: 地址,
											8: 简介,
										}
										for k, v := range 属性 {
											idx := ctx.UpsertItemField(k)
											结果[idx] = v
										}
										fmt.Println(结果)
										// 结果输出
										ctx.Output(结果)

										x := &request.Request{
											Url:          url,
											Rule:         "好友资料",
											DownloaderID: 0,
											Temp: map[string]interface{}{
												"好友名":  name,
												"好友ID": uidX[1],
												"性别":   sex[1],
												"认证":   认证,
												"关注":   关注,
												"粉丝":   粉丝,
												"微博":   微博,
												"地址":   地址,
												"简介":   简介,
											},
										}
										ctx.AddQueue(x)

									})

								}
							}
							//fmt.Println(slice0)
						}
					})
					//结果

				},
			},
			"用户资料": {
				ItemFields: []string{
					"Name",
					"ID",
					"Sex",
					"Idenfied",
					"Followed",
					"Fans",
					"Weibos",
					"Addr",
					"Intro",
					/*
						"好友名",
						"好友ID",
						"性别",
						"认证",
						"关注",
						"粉丝",
						"微博",
						"地址",
						"简介",
					*/
				},
				ParseFunc: func(ctx *Context) {
					query := ctx.GetDom()
					//var 属性 map[string]string
					//var title string
					//var detail string
					//好友资料的返回页中含有用户的信息，可在此处理函数中添加获取该用户fans的请求到队列中
					query.Find("script").Each(func(i int, s *goquery.Selection) {
						fmt.Println("search in script for 用户资料")
						//fmt.Println(s.Text())

						strXTmp := s.Text()
						if strings.Contains(strXTmp, "Pl_Core_T8CustomTriColumn__3") {
							fmt.Println("find Pl_Core_T8CustomTriColumn__3")
							slice1 := strings.Split(strXTmp, "\":\"")
							for i, v := range slice1 {
								if strings.Contains(v, "t_link S_txt1") {
									strXTmp02 := strings.Replace(v, "\\n", "", -1)
									strXTmp03 := strings.Replace(strXTmp02, "\\r", "", -1)
									strXTmp04 := strings.Replace(strXTmp03, "\\", "", -1)
									strXTmp05 := strings.Replace(strXTmp04, "\"})", "", -1)
									fmt.Printf("下标: %d \n", i)
									docX, err := goquery.NewDocumentFromReader(strings.NewReader(strXTmp05))
									if err != nil {
										fmt.Println(err)
									}

									url, _ := docX.Find(".S_line1 a").Attr("href")
									fmt.Println(url)

									url = strings.Replace(url, "//weibo.com/", "", -1)
									url = strings.Replace(url, "/follow?from=page_100505&wvr=6&mod=headfollow#place", "", -1)
									uid := strings.Replace(url, "p/100505", "", -1)
									strCountOfFollows := docX.Find("strong").Eq(0).Text()
									strCountOfFans := docX.Find("strong").Eq(1).Text()
									strCountOfWeibos := docX.Find("strong").Eq(2).Text()
									fmt.Println(uid, strCountOfFollows, strCountOfFans, strCountOfWeibos)

									//添加获取用户粉丝队列
									countOfFollows, _ := strconv.Atoi(strCountOfFollows)
									countOfFans, _ := strconv.Atoi(strCountOfFans)
									if glevel == 0 {
										countOfFollows = countOfFollows/30 + 1
										countOfFans = countOfFans/20 + 2
										//获取粉丝列表
										//for i := count1; i > 0; i-- {
										for i := 1; i <= countOfFans; i++ {
											ctx.AddQueue(&request.Request{
												Url:          "https://weibo.com/" + uid + "/fans?cfs=600&relate=fans&t=1&f=1&type=&Pl_Official_RelationFans__90_page=" + strconv.Itoa(i) + "#Pl_Official_RelationFans__90",
												Rule:         "好友列表",
												Header:       http.Header{"Cookie": []string{cookieStr}},
												DownloaderID: 0,
											})
										}

										//获取关注列表
										for i := 1; i <= countOfFollows; i++ {
											ctx.AddQueue(&request.Request{
												Url:          "https://weibo.com/p/100505" + uid + "/myfollow?t=1&cfs=&Pl_Official_RelationMyfollow__95_page=" + strconv.Itoa(i) + "#Pl_Official_RelationMyfollow__95",
												Rule:         "关注列表",
												Header:       http.Header{"Cookie": []string{cookieStr}},
												DownloaderID: 0,
											})
										}
									} else {
										countOfFollows = countOfFollows/30 + 1
										countOfFans = countOfFans/20 + 2
										if countOfFollows > 5 {
											countOfFollows = 5
										}
										if countOfFans > 5 {
											countOfFans = 5
										}

										//获取粉丝列表
										//for i := count1; i > 0; i-- {
										for i := 1; i <= countOfFans; i++ {
											ctx.AddQueue(&request.Request{
												//https://weibo.com/p/1005056850846383/follow?relate=fans&page=2#Pl_Official_HisRelation__59
												Url:          "https://weibo.com/p/100505" + uid + "/follow?relate=fans&page=" + strconv.Itoa(i) + "#Pl_Official_HisRelation__59",
												Rule:         "好友列表",
												Header:       http.Header{"Cookie": []string{cookieStr}},
												DownloaderID: 0,
											})
										}

										//获取关注列表
										for i := 1; i <= countOfFollows; i++ {
											ctx.AddQueue(&request.Request{
												//https://weibo.com/p/1005056850846383/follow?page=5#Pl_Official_HisRelation__59
												Url:          "https://weibo.com/p/100505" + uid + "/follow?page=" + strconv.Itoa(i) + "#Pl_Official_HisRelation__59",
												Rule:         "关注列表",
												Header:       http.Header{"Cookie": []string{cookieStr}},
												DownloaderID: 0,
											})
										}
									}

								}
							}
							//fmt.Println(slice0)
						}
					})

				},
			},
		},
	},
}
