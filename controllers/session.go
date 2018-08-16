package controllers

import (    //业务文件
"github.com/astaxie/beego"
_ "github.com/go-sql-driver/mysql"

	"ihome/utils"
)

type SessionController struct {
	beego.Controller
}


func (c *SessionController) Retdata(resp interface{}){
	c.Data["json"] = resp
	c.ServeJSON()
}


func (c *SessionController) GetSession() {

	// 打印调用的函数
	beego.Info("--------------GET api/v1.0/session GetSession()------------------")

	// 创建返回空间
	resp := make(map[string]interface{})
	//初始化是没有用户登陆的状态

	resp["errno"] = utils.RECODE_SESSIONERR
	resp["errmsg"] = utils.RecodeText(resp["errno"].(string))
	//延迟调用发送给前端json数据
	defer c.Retdata(resp)

	/*从session中获取 name 字段 */


	/*如果有返回成功 并且返回 name字段*/

	/*如果没有 初始化的时候默认返回错误*/

	return
}
