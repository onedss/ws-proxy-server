package main

import (
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
	_ "github.com/beego/samples/WebIM/routers"
)
const (
	APP_VER = "0.1.1.0227"
)

func main() {
	beego.Info(beego.BConfig.AppName, APP_VER)

	// Register template functions.
	beego.AddFuncMap("i18n", i18n.Tr)

	beego.Run(":8089")
}
