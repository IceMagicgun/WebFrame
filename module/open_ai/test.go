package open_ai

import (
	"WebFrame/model/open_ai_model"
	"net/http"
)

func Test(res http.ResponseWriter, req *http.Request) {
	res.Write([]byte(ChatBase(open_ai_model.Chat_GPT35_Turbo, "这是一条对话测试，如果你收到消息的话能给我打个招呼吗？")))
}
