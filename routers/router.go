package routers

import (
	"ihome/controllers"
	"github.com/astaxie/beego"
)

func init() { //路由模块
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/v1.0/areas", &controllers.AreasController{},"get:GetAreas")
}
