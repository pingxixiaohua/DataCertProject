package main

import (
	"DataCertProject/blockchain"
	"DataCertProject/db_mysql"
	_ "DataCertProject/routers"
	"github.com/astaxie/beego"
)


func main() {
	/*
		内存中的数据无法持久化存储，得先序列化
		序列化，可用json.Marshal将区块信息(结构体)转为[]byte格式的代码，储存到文件夹里，

		json.Marshal  结构体-->[]byte
		json.Unmarshal  []byte-->结构体

		除了json，还可用xml

	*/
	//准备一条区块链
	blockchain.NewBlockChain()



	//连接数据库
	db_mysql.ConnectDb()

	//静态资源路径设置
	beego.SetStaticPath("/js","./static/js")
	beego.SetStaticPath("/css","./static/css")
	beego.SetStaticPath("/img","./static/img")

	beego.Run()
}

