package open_ai_model

const (
	Chat_GPT35_Turbo      = "gpt-3.5-turbo"      //普通gpt3.5模型(4,096 tokens) ***这是最低消耗的模型
	Chat_GPT35_Turbo_0301 = "gpt-3.5-turbo-0301" //23年3月1号最新的gpt3.5模型(4,096 tokens)
	Chat_GPT35_Text       = "text-davinci-003"   //专门为语言设置的模型(4,097 tokens)
	Chat_GPT35_code       = "code-davinci-002"   //专门为编程设置的模型(8,001 tokens)
	Chat_GPT4             = "gpt-4"              //普通gpt4模型(8,192 tokens)
	Chat_GPT4_0314        = "gpt-4-0314"         //23年3月14号最新的gpt4模型(8,192 tokens)
	Chat_GPT4_32k         = "gpt-4-32k"          //4倍token的gpt4模型(32,768 tokens)
	Chat_GPT4_32k_0314    = "gpt-4-32k-0314"     //23年3月14号最新的gpt4模型(32,768 tokens)
)
