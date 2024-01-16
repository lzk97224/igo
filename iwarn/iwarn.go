package iwarn

import (
	"bytes"
	"github.com/lzk97224/igo/istr"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"strings"
)

func SendMsg(lines []string, url string) {
	content := strings.Join(lines, "\\n")
	body := "{\"msg_type\":\"text\",\"content\":{\"text\":\"" + content + "\"}}"
	resp, err := http.Post(url, "application/json", bytes.NewReader(istr.ToBytes(body)))
	if err != nil {
		log.Errorf("请求失败：%v", err)
	}

	all, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("读取结果失败：%v", err)
	}
	log.Debugf("发送结果:%v", all)
}
