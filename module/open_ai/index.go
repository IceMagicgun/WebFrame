package open_ai

import (
	debugM "WebFrame/core/debug"
	"WebFrame/model/open_ai_model"
	"encoding/json"
	"net/http"
)

type Req struct {
	Text string
}

func AskOnce(res http.ResponseWriter, req *http.Request) {

	var reqData Req
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&reqData)
	req.Body.Close()
	if err != nil {
		debugM.Error(err.Error())
	}
	debugM.Log(reqData.Text)
	res.Write([]byte(ChatBase(open_ai_model.Chat_GPT35_Turbo, reqData.Text)))
}
