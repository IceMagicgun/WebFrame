package open_ai

import (
	debugM "WebFrame/core/debug"
	"WebFrame/core/math_tool"
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

func AskForever(res http.ResponseWriter, req *http.Request) {
	session, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		session = &http.Cookie{
			Name:     "session",
			Value:    math_tool.UUID(),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   72 * 60 * 60,
		}
		http.SetCookie(res, session)
	}
	var reqData Req
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&reqData)
	req.Body.Close()
	if err != nil {
		debugM.Error(err.Error())
	}
	debugM.Log(reqData.Text)
	res.Write([]byte(ChatForever(open_ai_model.Chat_GPT35_Turbo, session.Value, reqData.Text)))
}

func GetHistory(res http.ResponseWriter, req *http.Request) {
	session, err := req.Cookie("session")
	if err == http.ErrNoCookie {
		session = &http.Cookie{
			Name:     "session",
			Value:    math_tool.UUID(),
			Path:     "/",
			HttpOnly: true,
			MaxAge:   72 * 60 * 60,
		}
		http.SetCookie(res, session)
	}
	res.Write([]byte(getHistory(session.Value)))
}
