package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/onedss/ws-proxy-server/routers"
)

const (
	APP_VER = "v1.0.0"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run(":8089")
}
