package wechat

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/olekukonko/tablewriter"
	"github.com/spf13/viper"
)

type tpl struct {
	ID      string `json:"template_id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// ListTemplates 列出当前账号所有的模板
func ListTemplates() {
	token := viper.GetString("wx_token")

	req, err := http.NewRequest("GET", "https://api.weixin.qq.com/cgi-bin/template/get_all_private_template?access_token="+token, nil)

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

	var dat map[string][]tpl
	if err := json.Unmarshal(body, &dat); err != nil {
		panic(string(body))
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"template_id", "title", "content"})

	for _, v := range dat["template_list"] {
		table.Append([]string{v.ID, v.Title, v.Content})
	}

	table.Render()
}
