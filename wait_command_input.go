package gUtils

import (
	"bufio"
	"fmt"
	"strings"
)

type WaitCommandInput struct {
	scanner bufio.Scanner
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
