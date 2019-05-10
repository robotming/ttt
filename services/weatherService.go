package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	// "os"
	// "regexp"
	// "fmt"
	// "unicode"
	"errors"
	"weather/services/util"
)

type WeatherService struct {
}

// 获取天气
func (this *WeatherService) GetWeatherByCityname(cityName string) (interface{}, error) {
	// 根据code  获取天气
	cityCode, err := this.GetCityCode(cityName)
	if err != nil {
		return nil, err
	}
	// 根据code  获取天气
	weather, _ := this.GetWeather(cityCode)

	b, _ := json.Marshal(weather)

	var dat map[string]interface{}
	if err := json.Unmarshal(b, &dat); err == nil {

		b, _ := json.Marshal(dat["data"])
		var result []map[string]interface{}
		json.Unmarshal(b, &result)

		return result[1], nil
	}

	return weather, nil
}

// 获取城市代码
func (this *WeatherService) GetCityCode(cityName string) (string, error) {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://beta.renrentianqi.com/api/weather/city/"+cityName, nil)

	// 根据城市名 获取 code
	if resp, err := client.Do(reqest); err != nil {
		return "", err
	} else {
		body, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()
		if err == nil {
			data := string(body)
			var result map[string]interface{}
			err := json.Unmarshal(body, &result)

			if len(result["data"].([]interface{})) == 0 {
				return "", errors.New("没有查到城市")
			}

			cityCode := util.Substr(data, 10, 9)
			return cityCode, err
		}
		return "", err
	}
}

// 获取天气接口
func (this *WeatherService) GetWeather(cityCode string) (interface{}, error) {
	client := &http.Client{}
	reqest, _ := http.NewRequest("GET", "http://beta.renrentianqi.com/api/weather/15day/"+cityCode, nil)

	// 根据城市名 获取 code
	if resp, err := client.Do(reqest); err != nil {
		return nil, err
	} else {
		body, err := ioutil.ReadAll(resp.Body)

		resp.Body.Close()

		var dat map[string]interface{}
		if err := json.Unmarshal(body, &dat); err == nil {
			return dat["data"], err
		}
		return nil, err
	}
}
