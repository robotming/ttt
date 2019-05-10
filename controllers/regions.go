package controllers

import (
	// "fmt"
	"strconv"
	. "weather/models"
	"weather/models/regions"
	"weather/services"

	"weather/services/util"

	"github.com/astaxie/beego"
)

type RegionsController struct {
	beego.Controller
}

var (
	result        []Regions
	RegionService services.RegionService
)

// 获取所有地区
func (this *RegionsController) Get() {
	result := regions.AllInDB() // 数据库获取所有地区

	// 计算第一级
	data := loop(result, "0")

	this.Data["json"] = map[string]interface{}{
		`code`: 0,
		`data`: data,
		`msg`:  "查询成功",
	}

	this.ServeJSON()
}

// 递归 组合地区数据
func loop(result []Regions, pid string) []interface{} {
	t := make([]interface{}, 0)
	for _, v := range result {
		if v.ParentId == pid {
			result := loop(result, strconv.FormatInt(v.Id, 10)) // Regions
			var temp = ChildrenRegions{
				Id:       v.Id,
				Name:     v.Name,
				ParentId: v.ParentId,
				Children: result,
			}
			t = append(t, temp)
		}
	}
	return t
}

// 根据ip获取地区
func (this *RegionsController) GetregionByIp() {
	// ip := this.Ctx.Input.Param(":ip")
	ip := this.Ctx.Input.IP()
	// 获取地址
	str := RegionService.GetRegionByIp(string(ip))
	regionStr := ""

	// 拆分地区
	result := util.UnicodeIndex(str, "省")
	regionStr = util.SubString(str, 0, result+1)
	str = util.SubString(str, result+1, 10)

	result = util.UnicodeIndex(str, "市")
	regionStr += " " + util.SubString(str, 0, result+1)

	str = util.SubString(str, result+1, 10)
	result = util.UnicodeIndex(str, "区")
	regionStr += " " + util.SubString(str, 0, result+1)

	this.Data["json"] = map[string]interface{}{
		`code`: 0,
		`data`: regionStr,
		`msg`:  "查询成功",
	}

	this.ServeJSON()
}
