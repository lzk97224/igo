package iwarn

import (
	"bytes"
	"github.com/lzk97224/igo/ijson"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

type MsgBody struct {
	MsgType string  `json:"msg_type"`
	Content Content `json:"content"`
}

type Content struct {
	Text string `json:"text"`
}

func SendMsg(lines []string, url string) {
	content := strings.Join(lines, "\n")
	bodyContent := MsgBody{
		MsgType: "text",
		Content: Content{Text: content},
	}
	resp, err := http.Post(url, "application/json", bytes.NewReader(ijson.ObjToBytes(bodyContent)))
	if err != nil {
		log.Errorf("请求失败：%v", err)
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("读取结果失败：%v", err)
	}
	log.Debugf("发送结果:%v", all)
}
