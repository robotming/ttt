package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"
	"weather/services/util"
)

type CalendarService struct {
}

// 获取一年中所有天数
func (this *CalendarService) GetAllDaysInYear(year int) []string {
	days := make([]string, 0)

	// year := time.Now().Year()
	// fmt.Printf("%T， %v\n", year, year)

	for month := 1; month <= 12; month++ {
		for day := 1; day <= 31; day++ {
			//如果是2月
			if month == 2 {
				if isLeapYear(year) && day == 30 { //闰年2月29天
					break
				} else if !isLeapYear(year) && day == 29 { //平年2月28天
					break
				} else {
					days = append(days, fmt.Sprintf("%d%02d%02d", year, month, day))
				}
			} else if month == 4 || month == 6 || month == 9 || month == 11 { //小月踢出来
				if day == 31 {
					break
				}
				days = append(days, fmt.Sprintf("%d%02d%02d", year, month, day))
			} else {
				days = append(days, fmt.Sprintf("%d%02d%02d", year, month, day))
			}
		}
	}
	return days
}

//判断是否为闰年
func isLeapYear(year int) bool { //y == 2000, 2004
	//判断是否为闰年
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

// 获取所有特殊日期
func (this *CalendarService) GetHolidays(date string) int {
	file := "./static/holidays/" + util.Substr(date, 0, 4) + "_data.json"

	if contents, err := ioutil.ReadFile(file); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)

		var holidaysJson map[string]int //所有特殊处理
		json.Unmarshal([]byte(result), &holidaysJson)

		info, ok := holidaysJson[util.Substr(date, 4, 4)]
		if ok == false {
			return -1
		}
		return info
	}
	return -1
}

// 获取周末
func (this *CalendarService) GetWeek(date string) int {

	p, _ := time.Parse("20060102", date) // 时间
	w := p.Weekday()
	t := int(w)

	if t == 0 || t == 6 {
		return 1
	}

	return 0

}
