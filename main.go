package main

import (
	"fmt"
	"github.com/astaxie/beego"
	"wechatdll/comm"
	_ "wechatdll/routers"
)

func main() {
	comm.RedisInitialize()
	_, err := comm.RedisClient.Ping().Result()
	if err != nil {
		panic(fmt.Sprintf("【Redis】连接失败，ERROR：%v", err.Error()))
	}
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/"] = "swagger"
	beego.SetLogFuncCall(false)
	//自定义错误页面
	beego.Run()

}
