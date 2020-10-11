package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

/*
默认显示的页面：用户注册页面

 */
func (c *MainController) Get() {

	c.TplName = "register.html"
}
