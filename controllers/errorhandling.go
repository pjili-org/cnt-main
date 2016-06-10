package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
    beego.Controller
}


func (c *ErrorController) Error401() {
    c.Data["Content"] = "Unauthorized"
    c.TplName = "ErrorPage.tpl"
}


func (c *ErrorController) Error403() {
    c.Data["Content"] = "Forbidden"
    c.TplName = "ErrorPage.tpl"
}


func (c *ErrorController) Error404() {
    c.Data["Content"] = "Not Found"
    c.TplName = "ErrorPage.tpl"
}


func (c *ErrorController) Error500() {
    c.Data["Content"] = "Oops something whent wrong"
    c.TplName = "ErrorPage.tpl"
}


func (c *ErrorController) Error503() {
    c.Data["Content"] = "Service Unavailable"
    c.TplName = "ErrorPage.tpl"
}
