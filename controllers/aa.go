package controllers

import (
	"DataCertProject/models"
	"fmt"
	"github.com/astaxie/beego"
)

type AaController struct {
	beego.Controller
}

func (a *AaController) Get() {
	phone := a.GetString("Phone")

	//4、从数据库中读取phone用户认证数据记录
	records, err := models.QueryRecordByPhone(phone)

	//5、根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil {
		fmt.Println(err)
		a.Ctx.WriteString("获取认证数据失败，请重试")
		return
	}
	a.Data["Phone"] = phone
	a.Data["Records"] = records
	a.TplName = "list_record.html"
}
