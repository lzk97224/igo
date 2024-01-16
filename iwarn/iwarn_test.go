package iwarn

import "testing"

func TestSendMsg(t *testing.T) {
	SendMsg([]string{
		"MJ报警",
		"code:1",
		"msg:账号不可用",
	}, "")
}
