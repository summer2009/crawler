package main

import (
	"context"
	"strconv"

	//"fmt"
	//"io/ioutil"
	"log"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	//var res string
	err = c.Run(ctxt, chromedp.Tasks{
		// 访问教务处页面
		chromedp.Navigate(`https://weibo.com`),

		// 等待table渲染成功，成功则说明已经获取到了正确的页面
		//chromedp.WaitVisible(`鱼翔浅底36`, chromedp.ByQuery),
		//chromedp.Text(`#1465346477`, &res, chromedp.ByID, chromedp.NodeVisible),
		chromedp.Sleep(time.Second * 40),
		// 获取body标签的html字符
		//chromedp.OuterHTML("body", &res),
		// 将chromedp.OuterHTML("body", &res) 替换为下面的代码
		chromedp.ActionFunc(func(ctx context.Context, h cdp.Executor) error {
			// 获取cookie
			cookies, err := network.GetAllCookies().Do(ctx, h)

			// 将cookie拼接成header请求中cookie字段的模式
			var c string
			for _, v := range cookies {
				c = c + v.Name + "=" + v.Value + ";"

				log.Println(v.Name + "=" + v.Value + ";" + strconv.FormatFloat(v.Expires, 'E', -1, 64))
			}
			log.Println("***********************************************************************************************************")
			log.Println("***********************************************************************************************************")
			log.Println("***********************************************************************************************************")
			log.Println("***********************************************************************************************************")
			log.Println(c)

			if err != nil {
				return err
			}
			return nil
		}),
	})
	if err != nil {
		log.Fatal(err)
	}

	// 关闭chrome实例
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// 等待chrome实例关闭
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}

	// 输出html字符串
	//log.Printf(res)
}
