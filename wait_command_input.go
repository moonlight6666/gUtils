package gUtils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type WaitCommandInput struct {
	scanner *bufio.Scanner
}

func NewWaitCommandInput() *WaitCommandInput {
	return &WaitCommandInput{
		scanner: bufio.NewScanner(os.Stdin),
	}
}
func (w *WaitCommandInput) Wait(name string) string {
	return w.do(name, "")
}

func (w *WaitCommandInput) WaitWithDefault(name string, defaultValue string) string {
	return w.do(name, defaultValue)
}

func (w *WaitCommandInput) do(name, defaultValue string) string {
	var input string
	for {
		if defaultValue == "" {
			fmt.Printf("请输入%s：", name)
		} else {
			fmt.Printf("请输入%s(默认值:%s)：", defaultValue, name)
		}
		w.scanner.Scan()
		input = strings.TrimSpace(w.scanner.Text())
		if input == "" {
			if defaultValue != "" {
				return defaultValue
			}
			fmt.Println("请重新输入~")
			continue
		}
		break
	}
	return input
}
