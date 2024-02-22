package common

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

type Text struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content string `json:"content"`
	} `json:"text"`
}

var message = Text{
	Msgtype: "text",
}

var wechatNoteMap = make(map[string]string, len(MyConfig.Wechat))

func GetWechatUrl(note string) string {
	return WechatUrl + wechatNoteMap[note]
}

func Send(msg, note string) {
	url := GetWechatUrl(note)
	message.Text.Content = msg
	muta, _ := json.Marshal(message)
	req, err := http.NewRequest("POST", url, bytes.NewReader(muta))
	if err != nil {
		log.Println("err:" + err.Error())
		return
	}
	client := &http.Client{}
	_, err = client.Do(req)
	if err != nil {
		log.Println("err:" + err.Error())
	}
}
