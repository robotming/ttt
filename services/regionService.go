package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type RegionService struct {
}

// 根据ip获取地区
func (this *RegionService) GetRegionByIp(ip string) string {
	client := &http.Client{}

	request, _ := http.NewRequest("GET", "http://iplocation.7654.com/v1?ip="+ip, nil)

	if response, err := client.Do(request); err != nil {
		return ""
	} else {
		body, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()

		var region map[string]string
		if err := json.Unmarshal(body, &region); err == nil {
			return region["city"]
		}
	}
	return ""
}
