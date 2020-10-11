package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	beego.Controller
}

func (r *RegisterController) Post()  {
	//1、解析请求数据
	var user models.User
	err := r.ParseForm(&user)
	if err != nil {
		r.Ctx.WriteString("解析数据错误，请重试")
		return
	}
	//2、保存用户信息到数据库
	_,err =user.SaveUser()
	if err != nil {
		r.Ctx.WriteString("用户注册失败，请重试")
		return
	}

	//注册成功后，跳转到注册页面
	r.TplName = "login.html"
	//3、返回前端结果
}