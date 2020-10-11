package controllers

import (
	"DataCertProject/models"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

//直接访问login.html页面的请求
func (l *LoginController) Get() {
	//设置模板文件Login.html
	l.TplName = "login.html"

}

//用户登录接口
func (l *LoginController) Post() {
	var user models.User
	err := l.ParseForm(&user)
	if err != nil {
		l.Ctx.WriteString("抱歉，用户信息解析失败，请重试")
		return
	}

	//查询数据库的用户信息
	u, err := user.QueryUser()
	if err != nil {
		 l.Ctx.WriteString("用户登陆失败，请重试")
		return
	}
	//登陆成功，跳转home.html
	l.Data["Phone"] = u.Phone
	l.TplName = "home.html"//{{。Phone}}

}
