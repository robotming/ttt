package routers

import (
	"weather/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})                                     // 首页
	beego.Router("/api/weather/:cityname", &controllers.WeatherController{})             // 天气数据
	beego.Router("/api/regions", &controllers.RegionsController{})                       // 全部地区
	beego.Router("/api/regionbyip", &controllers.RegionsController{}, "*:GetregionByIp") // 根据ip获取地区
	beego.Router("/api/holiday/:year", &controllers.CalendarController{}, "*:Holiday")   // 获取节假日
}
