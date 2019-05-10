package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"weather/services"

	"github.com/astaxie/beego"
)

type CalendarController struct {
	beego.Controller
}

var (
	//天气服务
	CalendarService services.CalendarService
)

// 节假日  0-班 1-休 2-假
func (this *CalendarController) Holiday() {
	year := this.Ctx.Input.Param(":year")
	var info int
	data := make(map[string]int)

	// 检查日期格式
	// 根据年份获取所有天数
	yearint, _ := strconv.Atoi(year)
	days := CalendarService.GetAllDaysInYear(yearint)

	for _, v := range days {
		// 读取文件，获取节假日数据
		info = CalendarService.GetHolidays(v)
		// 获取周末数据
		if info == -1 {
			info = CalendarService.GetWeek(v)
		}
		data[v] = info
	}

	this.Data["json"] = map[string]interface{}{
		`code`: 200,
		`data`: data,
		`msg`:  "查询成功",
	}
	this.ServeJSON()
}

// 节假日  0-班 1-休 2-假
func (this *CalendarController) HolidayBak() {
	date := this.Ctx.Input.Param(":date")
	url := "http://holidays.zhangrui.com/?d=" + date
	client := &http.Client{}
	request, _ := http.NewRequest("GET", url, nil)
	response, _ := client.Do(request)

	body, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()
	// 处理数据
	var data map[string]interface{}
	json.Unmarshal(body, &data)
	fmt.Println(data)

	this.Data["json"] = map[string]interface{}{
		`code`: 200,
		`data`: data["info"],
		`msg`:  "查询成功",
	}
	this.ServeJSON()
}
