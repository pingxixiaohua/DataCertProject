package routers

import (
	"DataCertProject/controllers"
	"github.com/astaxie/beego"
)

/*
	路由功能。用于接收并分发接受到的浏览器请求
*/

func init() {
    beego.Router("/", &controllers.MainController{})
    //用户注册的接口请求
	beego.Router("/user_register", &controllers.RegisterController{})
    //访问直接登陆的页面访问接口
    beego.Router("/login.html",&controllers.LoginController{})
	//用户登录请求接口
	beego.Router("/user_login",&controllers.LoginController{})
	//上传文件接口
	beego.Router("/upload",&controllers.UploadFileController{})
	//认证数据列表页面，跳回新增取证
	beego.Router("/upload_file.html",&controllers.UploadFileController{})
	//查看认证证书页面cert_detail.html
	beego.Router("/cert_detail.html",&controllers.CertDetailController{})


    beego.Router("/list_record.html",&controllers.AaController{})

}
