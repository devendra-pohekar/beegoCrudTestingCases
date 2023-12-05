package main

import (
	"crud/routers"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	// logs.SetLogger(logs.AdapterMultiFile, `{"filename":"testLog","separate":["emergency", "alert", "critical", "error", "warning", "notice", "info", "debug"]}`)

	routers.RoutersFunction()
	beego.Run()
}
