package main

import (
	_ "github.com/udistrital/configuracion_api/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/lib/pq"
	notificacionlib "github.com/udistrital/notificaciones_lib"
	"github.com/udistrital/auditoria"
	"github.com/udistrital/utils_oas/apiStatusLib"
	"github.com/udistrital/utils_oas/responseformat"

)

func init() {
	orm.RegisterDataBase("default", "postgres", "postgres://"+beego.AppConfig.String("PGuser")+":"+beego.AppConfig.String("PGpass")+"@"+beego.AppConfig.String("PGurls")+"/"+beego.AppConfig.String("PGdb")+"?sslmode=disable&search_path="+beego.AppConfig.String("PGschemas")+"")
}

func main() {
	beego.ErrorHandler("400",nil)
	orm.Debug = true
	beego.BConfig.RecoverFunc = responseformat.GlobalResponseHandler
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	apistatus.Init()
	auditoria.InitMiddleware()
	notificacionlib.InitMiddleware()
	beego.Run()
	
}
