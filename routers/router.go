package routers

import (
	"github.com/pjili-org/cnt-main/controllers"
	"github.com/astaxie/beego"
)


func init() {
    beego.Router("/", &controllers.MainController{})
}
