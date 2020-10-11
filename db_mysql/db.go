package db_mysql

import (
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func ConnectDb()  {
	//配置
	config := beego.AppConfig
	dbDriver := config.String("db_driverName")
	dbUser := config.String("db_user")
	dbPassword := config.String("db_password")
	dbIp := config.String("db_ip")
	dbName := config.String("db_name")

	//连接数据库
	connUrl := dbUser +":" + dbPassword + "@tcp("+dbIp+")/"+dbName+"?charset=utf8"

	db, err := sql.Open(dbDriver,connUrl)
	if err != nil {
		fmt.Println(err.Error())
		panic("数据库连接失败，请检查配置")
	}

	Db = db
}
