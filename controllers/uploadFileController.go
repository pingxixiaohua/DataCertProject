package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"bufio"
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"time"
)

type UploadFileController struct {
	beego.Controller
}

//跳转新增认证文件的页面upload_file.html
func (u *UploadFileController) Get()  {
	phone := u.GetString("phone")
	u.Data["Phone"] = phone
	u.TplName = "home.html"
}

func (u *UploadFileController) Post() {
	//1.获取客户端上传的文件以及其他form表单的信息
	fileTitle := u.Ctx.Request.PostFormValue("upload_title")
	phone := u.Ctx.Request.PostFormValue("phone")

	file, header, err := u.GetFile("upload_file")
	if err != nil {
		u.Ctx.WriteString("文件解析失败，请重试")
		return
	}

	//3、关闭文件
	defer file.Close()

	fmt.Println("自定义的文件标题:",fileTitle)
	fmt.Println("文件名称:",header.Filename)
	fmt.Println("文件的大小:",header.Size)//字节大小

	fmt.Println(file)
	//u.Ctx.WriteString("解析到上传文件，文件名是："+header.Filename)

	//2、将文件保存在本地的一个目录中

	//文件全路径： 路径 + 文件名 + "." + 扩展名
	uploadDir := "./static/img/" + header.Filename
	/* 777 文件权限
												读	写	执行
		文件所有者拥有的权限						4	2	1
		文件所有者所在的组的用户对文件拥有的权限	4	2	1
		其他用户对文件拥有的权限					4	2	1
	 */
	savaFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE,777)
	//saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)

	//创建一个writer：用于向硬盘上写一个文件
	writer := bufio.NewWriter(savaFile)
	file_size, err := io.Copy(writer,file)
	if err != nil {
		u.Ctx.WriteString("保存电子数据失败，请重试")
		return
	}
	defer savaFile.Close()
	fmt.Println("拷贝的文件的大小是：",file_size)

	//2、计算文件的hash
	hashFile, err := os.Open(uploadDir)
	defer hashFile.Close()
	fmt.Println(savaFile)
	hash, err := util.MD5HashReader(hashFile)

	//3、将上传的记录保存到数据库中
	record := models.UploadRecord{}
	record.FileName = header.Filename
	record.FileSize = header.Size
	record.FileTitle = fileTitle
	record.CertTime = time.Now().Unix()
	record.FileCert = hash
	record.Phone = phone //手机
	_, err = record.SaveRecord()
	if err != nil {
		fmt.Println(err.Error())
		u.Ctx.WriteString("数据认证错误，请重试")
		return
	}

	//将要认证的文件hash值及个人信息保存到区块链上，上链
	_, err = blockchain.CHAIN.SaveData([]byte(hash))
	if err != nil {
		u.Ctx.WriteString("认证数据上链失败，请重试")
	}

	//4、从数据库中读取phone用户认证数据记录
	records, err := models.QueryRecordByPhone(phone)

	//5、根据文件保存结果，返回相应的提示信息或者页面跳转
	if err != nil {
		fmt.Println(err)
		u.Ctx.WriteString("获取认证数据失败，请重试")
		return
	}
	u.Data["Records"] = records
	u.Data["Phone"] = phone
	u.TplName = "list_record.html"


}