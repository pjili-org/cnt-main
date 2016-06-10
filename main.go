package main

import (
	_ "github.com/pjili-org/cnt-main/routers"
	"github.com/astaxie/beego"
  "github.com/pjili-org/cnt-main/controllers"
)

func main() {
  beego.ErrorController(&controllers.ErrorController{})
	beego.Run()
}

