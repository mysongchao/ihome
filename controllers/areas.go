package controllers

import (    //业务文件
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"ihome/models"
)

type AreasController struct {
	beego.Controller
}


func (c *AreasController) Retdata(resp map[string]interface{}){
	c.Data["json"] = &resp
	c.ServeJSON()
}


func (c *AreasController) GetAreas() {

	// 打印调用的函数
	beego.Info("--------------GET api/v1.0/areas Getreas()------------------")

	// 创建返回空间
	resp := make(map[string]interface{})
	defer c.Retdata(resp) // 延迟调用
	/*从缓存数据库中获取缓存数据，如果存在就发给前端*/

	/*如果缓存没有就查询mysql 获取 数据*/
	o :=orm.NewOrm() // 创 orm句柄

	var area []models.Area  // 缓存空间
	qs:=o.QueryTable("area")
	num,err:= qs.All(&area)
	if err!= nil{
		beego.Info("o.QueryTable(area) err",err,num)
		resp["errno"]= 404
		resp["errmsg"]="数据库查询失败"
		//c.Data["json"] = &resp
		//c.ServeJSON()
		return
	}
	if num ==0 {
		beego.Info("o.QueryTable(area) 404",num)
		resp["errno"]= 404
		resp["errmsg"]="没有数据"
		resp["data"]=area
		//c.Data["json"] = &resp
		//c.ServeJSON()
 		return
	}
	beego.Info("o.QueryTable(area) ok")
	resp["errno"]= 0
	resp["errmsg"]="成功"
	resp["data"]=area
	//c.Data["json"] = &resp
	//c.ServeJSON()
	/*将获取 的数据打包JSON 数据存入数据库  缓存数据库*/

	/*将数据发送给前端	*/
}
