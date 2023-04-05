package debugM

import (
	"WebFrame/core/time_tool"
	"fmt"
)

func Log(message string) {
	println(time_tool.NowText() + " Log:" + message)
}

func Error(message string) {
	fmt.Printf("\033[31m%s Error:%s\033[0m\n", time_tool.NowText(), message)
	panic(message)
}

func Warring(message string) {
	println(time_tool.NowText() + " Warring:" + message)
}
