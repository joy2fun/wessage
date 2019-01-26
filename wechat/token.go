package wechat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/spf13/viper"
)

// RefreshToken 更新 access_token，如果 token 未过期，忽略
func RefreshToken() {
	expired := viper.GetInt("wx_token_expired_at")

	if int64(expired) > time.Now().Unix() {
		// fmt.Println("cache alive")
		return
	}

	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential", nil)
	q := req.URL.Query()
	q.Add("appid", viper.GetString("wx_appid"))
	q.Add("secret", viper.GetString("wx_secret"))
	req.URL.RawQuery = q.Encode()

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
	var dat map[string]interface{}
	if err := json.Unmarshal(body, &dat); err != nil {
		fmt.Println(string(body))
		os.Exit(1)
	}
	// fmt.Println(dat)
	if dat["errcode"] != nil {
		fmt.Println(string(body))
		os.Exit(1)
	}

	viper.Set("wx_token", dat["access_token"])
	viper.Set("wx_token_expired_at", time.Now().Unix()+3600) // ttl 3600s
	viper.WriteConfig()
}
