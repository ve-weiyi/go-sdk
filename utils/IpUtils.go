package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type BaiduIpResp struct {
	Status       string `json:"status"`
	T            string `json:"t"`
	SetCacheTime string `json:"set_cache_time"`
	Data         []struct {
		ExtendedLocation string `json:"ExtendedLocation"`
		OriginQuery      string `json:"OriginQuery"`
		Appinfo          string `json:"appinfo"`
		DispType         int    `json:"disp_type"`
		Fetchkey         string `json:"fetchkey"`
		Location         string `json:"location"` //地址
		Origip           string `json:"origip"`
		Origipquery      string `json:"origipquery"`
		Resourceid       string `json:"resourceid"`
		RoleId           int    `json:"role_id"`
		ShareImage       int    `json:"shareImage"`
		ShowLikeShare    int    `json:"showLikeShare"`
		Showlamp         string `json:"showlamp"`
		Titlecont        string `json:"titlecont"`
		Tplt             string `json:"tplt"`
	} `json:"data"`
}

// GetIpSource 获取ip对应的城市地区
func GetIpSource(ipAddress string) string {
	resp, err := http.Get(fmt.Sprintf("http://opendata.baidu.com/api.php?query=" + ipAddress + "&co=&resource_id=6006&oe=utf8"))
	if err != nil {
		return ""
	}
	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}

	var result BaiduIpResp
	if err := json.Unmarshal(out, &result); err != nil {
		return ""
	}

	if len(result.Data) > 0 {
		return result.Data[0].Location
	} else {
		return ""
	}
}
