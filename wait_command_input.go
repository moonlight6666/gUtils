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
	var input string
	for {
		fmt.Printf("请输入%s：", name)
		w.scanner.Scan()
		input = strings.TrimSpace(w.scanner.Text())
		if input == "" {
			fmt.Println("请重新输入~")
			continue
		}
		break
	}
	return input
}
