package main

import "github.com/gevinzone/lgthw/loop1/c06WebClient/client"

func main() {
	if err := client.DefaultGet("https://www.baidu.com/"); err != nil {
		panic(err)
	}
	cli := client.Setup(true, false)
	if err := client.DoGetOps(cli, "https://www.baidu.com/"); err != nil {
		panic(err)
	}
	if err := client.DoGetOps(cli, "http://www.baidu.com/"); err != nil {
		panic(err)
	}
	c := client.MyClient{Client: cli}
	if err := c.DoGetOps("https://www.baidu.com/"); err != nil {
		panic(err)
	}
	client.Setup(false, false)
	if err := client.DefaultGet("http://www.baidu.com/"); err != nil {
		panic(err)
	}
	client.Setup(true, true)
	if err := client.DefaultGet("https://www.baidu.com/"); err != nil {
		panic(err)
	}

}
