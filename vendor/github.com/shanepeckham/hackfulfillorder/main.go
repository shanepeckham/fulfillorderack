package main

import (
	_ "github.com/shanepeckham/hackfulfillorder/routers"

	"github.com/astaxie/beego"
)

func main() {

	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	beego.Run()
}
