package math_tool

import (
	debugM "WebFrame/core/debug"
	"crypto/rand"
	"encoding/hex"
)

func UUID() string {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		debugM.Error(err.Error())
	}
	// 设置版本号和变体标识
	uuid[6] = (uuid[6] & 0x0f) | 0x40
	uuid[8] = (uuid[8] & 0x3f) | 0x80
	return hex.EncodeToString(uuid)
}
