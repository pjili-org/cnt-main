package controllers

import (
	"github.com/astaxie/beego"
)

var(
  BaseUrl string = beego.AppConfig.String("BaseUrl")
  StaticAssetsUrl string = beego.AppConfig.String("StaticAssetsUrl")
)


func init() {
  beego.Info("controllers.init() - BaseUrl:", BaseUrl)
  beego.Info("controllers.init() - StaticAssetsUrl:", StaticAssetsUrl)
}


type MainController struct {
	beego.Controller
}


func (this *MainController) Prepare() {
  beego.Debug("In MainController.Prepare() - Start")

  this.Data["Title"] = "pjili.org"
  this.Data["BaseUrl"] = BaseUrl
  this.Data["StaticAssetsUrl"] = StaticAssetsUrl
}


func (c *MainController) Get() {
	c.TplName = "index.tpl"
}
