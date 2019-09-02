package wechat

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/spf13/viper"
)

type msg struct {
	Receiver   string
	TemplateID string
	Content    string
	Link       string
}

// SendTemplateMessage 发送模板消息
func SendTemplateMessage(message, link, receiver, templateID string) {
	const msgTpl = `{
		"touser":"{{.Receiver}}",
		"template_id":"{{.TemplateID}}",
		"url": "{{.Link}}",
		"data":{
			"msg": {
				"value": "{{.Content}}"
			}
		}
	}`

	if len(receiver) == 0 {
		receiver = viper.GetString("wx_receiver")
	}

	if len(templateID) == 0 {
		templateID = viper.GetString("wx_template_id")
	}

	t := template.Must(template.New("msg").Parse(msgTpl))
	buf := bytes.NewBufferString("")
	t.Execute(buf, msg{receiver, templateID, message, link})

	callTemplateMessageAPI(buf)
}

func callTemplateMessageAPI(payload *bytes.Buffer) {
	token := viper.GetString("wx_token")
	req, err := http.NewRequest("POST", "https://api.weixin.qq.com/cgi-bin/message/template/send?access_token="+token, payload)

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
	fmt.Println(string(body))
}
