package main

import (
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)

func main() {
	//连接数据库
	db_mysql.ConnectDb()

	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

