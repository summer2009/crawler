          strCookedReturn := ctx.GetCookie()
					if strCookedReturn != "" {
						fmt.Println("!!!!!!!!!!!!!the old cookie!!!!!!!!!!!!!")
						fmt.Println(cookieStr)
						fmt.Println("!!!!!!!!!!!!!the cookie returned!!!!!!!!!!!!!")
						fmt.Println(strCookedReturn)
						ctx.GetSpider().Stop()

						//重新添加下载任务至队列,启动spider
						fmt.Println("You should add new tasks to the Queue and start the Spider")

						ctx.GetSpider().Start()

						//获取用户信息
						//https://weibo.com/6606483016/profile?topnav=1&wvr=6&is_all=1
						ctx.AddQueue(&request.Request{
							Url:          "https://weibo.com/" + "6606483016" + "/profile?topnav=1&wvr=6&is_all=1",
							Rule:         "用户资料",
							Header:       http.Header{"Cookie": []string{strCookedReturn}},
							DownloaderID: 0,
						})

					}
