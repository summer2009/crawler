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

																
	}
