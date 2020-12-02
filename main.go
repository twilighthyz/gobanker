package main

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	_ "gobanker/routers"
	"gobanker/spider"

	"github.com/astaxie/beego"
)

func init() {
	//orm.RegisterDriver("postgres",orm.DRPostgres)
	//var dataSource string = beego.AppConfig.String("pgsql_datasource")
	//orm.RegisterDataBase("default","postgres", dataSource)
	//orm.RunSyncdb("default",false,true)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	spider.CsIndexIndustryHandler()
	orm.Debug = true
	beego.Run()
}
