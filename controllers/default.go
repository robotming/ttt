package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {
	this.Data["json"] = map[string]interface{}{
		`status`: 0,
		`msg`:    "这里是接口首页",
	}
	this.ServeJSON()
}
