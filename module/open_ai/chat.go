package open_ai

import (
	"WebFrame/core/config"
	debugM "WebFrame/core/debug"
	"WebFrame/core/time_tool"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
)

type ChatAns struct {
	Id      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	Choices []struct {
		Index        int        `json:"index"`
		Message      OneMessage `json:"message"`
		FinishReason string     `json:"finish_reason"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

type OneMessage struct {
	Role    string `json:"role"` //"system"、"user" 、 "assistant"
	Content string `json:"content"`
}

var (
	lastTime = int64(0)
)
var mutex sync.Mutex

func ChatBase(model string, message string) string {
	//创建客户端 后续可以移至http工具类中
	transport := http.DefaultTransport
	config.InitConfig("")
	if config.Config.Proxy != nil {
		proxyUrl, has := config.Config.Proxy[config.Config.Env]
		if has {
			transport = &http.Transport{
				Proxy: func(req *http.Request) (*url.URL, error) {
					return url.Parse(proxyUrl)
				},
			}
		}
	}
	client := &http.Client{
		Transport: transport,
	}

	reqData := map[string]interface{}{}
	reqData["model"] = model
	reqData["messages"] = []OneMessage{
		{
			Role:    "user",
			Content: message,
		},
	}
	reqBody, _ := json.Marshal(reqData)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		debugM.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Config.OpenAI.OPENAI_API_KEY)

	//加锁判断 后续需写一个锁的工具类
	mutex.Lock()
	now := time_tool.Now()
	if now-lastTime < 5 {
		return "5秒内只能请求一次"
	}
	lastTime = now
	mutex.Unlock()

	resp, err := client.Do(req)
	if err != nil {
		debugM.Error(err.Error())
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	debugM.Log(string(body))
	data := ChatAns{}
	json.Unmarshal(body, &data)

	ans, _ := json.Marshal(data.Choices[0].Message)
	return string(ans)
}

var (
	Cache = map[string][]OneMessage{}
)

func ChatForever(model string, session string, message string) string {
	//创建客户端 后续可以移至http工具类中
	transport := http.DefaultTransport
	config.InitConfig("")
	if config.Config.Proxy != nil {
		proxyUrl, has := config.Config.Proxy[config.Config.Env]
		if has {
			transport = &http.Transport{
				Proxy: func(req *http.Request) (*url.URL, error) {
					return url.Parse(proxyUrl)
				},
			}
		}
	}
	client := &http.Client{
		Transport: transport,
	}

	reqData := map[string]interface{}{}
	reqData["model"] = model
	reqData["messages"] = append(Cache[session], OneMessage{
		Role:    "user",
		Content: message,
	})
	reqBody, _ := json.Marshal(reqData)
	req, err := http.NewRequest("POST", "https://api.openai.com/v1/chat/completions", bytes.NewReader(reqBody))
	if err != nil {
		debugM.Error(err.Error())
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.Config.OpenAI.OPENAI_API_KEY)

	//加锁判断 后续需写一个锁的工具类
	mutex.Lock()
	hasUnlock := false
	defer func() {
		if !hasUnlock {
			mutex.Unlock()
		}
	}()
	now := time_tool.Now()
	if now-lastTime < 5 {
		return "5秒内只能请求一次"
	}
	lastTime = now
	mutex.Unlock()
	hasUnlock = true

	resp, err := client.Do(req)
	if err != nil {
		debugM.Error(err.Error())
	}

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	debugM.Log(string(body))
	data := ChatAns{}
	json.Unmarshal(body, &data)

	Cache[session] = append(Cache[session], OneMessage{
		Role:    "user",
		Content: message,
	})
	Cache[session] = append(Cache[session], data.Choices[0].Message)

	ans, _ := json.Marshal(data.Choices[0].Message)
	return string(ans)
}

func getHistory(session string) string {
	ans := ""
	for _, item := range Cache[session] {
		if ans != "" {
			ans += "\n"
		}
		if item.Role == "user" {
			ans += "你："
		} else {
			ans += "ChatGpt："
		}
		ans += item.Content
	}
	return ans
}
