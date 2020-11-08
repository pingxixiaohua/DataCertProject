package controllers

import (
	"DataCertProject/blockchain"
	"DataCertProject/models"
	"DataCertProject/util"
	"fmt"
	"github.com/astaxie/beego"
	"strings"
)

/*
 *证书页面控制器
 */
type CertDetailController struct {
	beego.Controller
}

func (c *CertDetailController) Get()  {
	//获取前端get请求时携带的cer_id数据
	certId := c.GetString("cert_id")//获取get请求携带的数据
	fmt.Println("认证id：",certId)
	//根据cert_id到区块链上查询具体的信息
	block, err := blockchain.CHAIN.QueryBlockByCertId([]byte(certId))
	if err != nil {
		c.Ctx.WriteString("链上数据查询错误")
		return
	}
	//查询未遇到错误，有两种，查到了和未查到
	if block == nil {
		c.Ctx.WriteString("未找到相关数据，请检查重试")
	}
	//certId = hex.EncodeToString()
	certRecord, err := models.DeSerializeRecord(block.Data)
	certRecord.CertHashStr = string(certRecord.CertHash)
	certRecord.CertIdStr = strings.ToUpper(string(certRecord.CertId))
	certRecord.CertTimeFormat = util.TimeFormat(certRecord.CertTime,0,util.TIME_FORMAT_ONE)
	c.Data["CertRecord"] =certRecord

	//跳转页面
	c.TplName = "cert_detail.html"

}