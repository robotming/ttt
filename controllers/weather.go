package controllers

import (
	"fmt"
	"weather/services"

	"github.com/astaxie/beego"
)

type WeatherController struct {
	beego.Controller
}

var (
	//天气服务
	weatherSer services.WeatherService
)

func (this *WeatherController) Get() {
	cityName := this.Ctx.Input.Param(":cityname")
	if cityName == "" {
		this.Ctx.WriteString("参数不正确")
		return
	}

	result, err := weatherSer.GetWeatherByCityname(cityName)

	if err != nil {
		this.Data["json"] = map[string]interface{}{
			`code`: 1,
			`data`: result,
			`msg`:  fmt.Sprintf("%v", err),
		}
	} else {
		this.Data["json"] = map[string]interface{}{
			`code`: 0,
			`data`: result,
			`msg`:  "查询成功",
		}
	}

	this.ServeJSON()
}
