package routers

import (
	"ihome/controllers"
	"github.com/astaxie/beego"
)

func init() { //路由模块
    beego.Router("/", &controllers.MainController{})
}
