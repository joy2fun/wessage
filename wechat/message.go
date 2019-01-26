package wechat

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Msg struct {
	Receiver   string
	TemplateId string
	Content    string
}

var receiver string
var templateId string

var CmdSend = &cobra.Command{
	Use:   "send [content]",
	Short: "Send template message",
	Long:  ``,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		const msgTpl = `{
			"touser":"{{.Receiver}}",
			"template_id":"{{.TemplateId}}",
			"data":{
				"msg": {
					"value": "{{.Content}}"
				}
			}
		}`

		if len(receiver) == 0 {
			receiver = viper.GetString("wx_receiver")
		}

		if len(templateId) == 0 {
			templateId = viper.GetString("wx_template_id")
		}

		t := template.Must(template.New("msg").Parse(msgTpl))
		buf := bytes.NewBufferString("")
		t.Execute(buf, Msg{receiver, templateId, args[0]})

		RefreshToken()
		SendTemplateMessage(buf)
	},
}

func InitSendCmd() {
	CmdSend.Flags().StringVarP(&templateId, "template", "t", "", "template id. use 'wx_template_id' in config if not present")
	CmdSend.Flags().StringVarP(&receiver, "receiver", "r", "", "receiver's openId. use 'wx_receiver' in config if not present")
}

func SendTemplateMessage(payload *bytes.Buffer) {
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
