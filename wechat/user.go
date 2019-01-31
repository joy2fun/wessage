package wechat

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

type openID struct {
	Openid string `json:"openid"`
	Lang   string `json:"lang"`
}

type openIDList struct {
	Items []string `json:"openid"`
}

type openIDRequest struct {
	Items []openID `json:"user_list"`
}

type openIDResponse struct {
	Total int        `json:"total"`
	Count int        `json:"count"`
	Data  openIDList `json:"data"`
	Next  string     `json:"next_openid"`
}

type userInfoResponse struct {
	OpenID string `json:"openid"`
	Name   string `json:"nickname"`
}

// ListUsers 展示获取用户列表
func ListUsers() {
	url := "https://api.weixin.qq.com/cgi-bin/user/get?access_token=" + viper.GetString("wx_token")
	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var dat openIDResponse
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	var items []openID

	for _, v := range dat.Data.Items {
		items = append(items, openID{v, "zh_CN"})
	}

	jsonBody := &openIDRequest{
		Items: items,
	}

	jsonStr, _ := json.Marshal(&jsonBody)
	buf := bytes.NewBuffer(jsonStr)

	callUserInfoAPI(buf)
}

// 批量获取用户信息
func callUserInfoAPI(payload *bytes.Buffer) {
	token := viper.GetString("wx_token")
	url := "https://api.weixin.qq.com/cgi-bin/user/info/batchget?access_token=" + token
	req, err := http.NewRequest("POST", url, payload)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if err != nil {
		panic(err)
	}

	body, _ := ioutil.ReadAll(resp.Body)

	var dat map[string][]userInfoResponse
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"open_id", "name"})

	for _, v := range dat["user_info_list"] {
		table.Append([]string{v.OpenID, v.Name})
	}

	table.Render()
}
