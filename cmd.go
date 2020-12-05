package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func Cmd(commandName string, params []string) (string, error) {
	return CmdAndChangeDir("", commandName, params)
	//cmd := exec.Command(commandName, params...)
	//fmt.Println("Cmd", cmd.Args)
	//var out bytes.Buffer
	//cmd.Stdout = &out
	//cmd.Stderr = cmd.Stdout
	//err := cmd.Run()
	//return out.String(), err
	//err := cmd.Start()
	//if err != nil {
	//	return out.String(), err
	//	//log.Fatal(err)
	//}
	//err = cmd.Wait()
	//return out.String(), err
}

func CmdAndChangeDir(dir string, commandName string, params []string) (string, error) {
	cmd := exec.Command(commandName, params...)
	//fmt.Println("CmdAndChangeDir", dir, cmd.Args)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = cmd.Stdout
	if dir != "" {
		cmd.Dir = dir
	}
	err := cmd.Run()
	return out.String(), err
	//err := cmd.Start()
	//if err != nil {
	//	return "", err
	//}
	//
	//err = cmd.Wait()
	//return out.String(), err
}

func CmdAndRealTimePrint(dir string, commandName string, out io.Writer, params []string) error {
	cmd := exec.Command(commandName, params...)
	fmt.Println("CmdAndRealTimeOut", cmd.Args)
	fmt.Fprintf(out, "$%s %s\n", dir, strings.Join(cmd.Args, " "))
	// 命令的错误输出和标准输出都连接到同一个管道
	stdout, err := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout
	if dir != "" {
		cmd.Dir = dir
	}
	if err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}
	// 从管道中实时获取输出并打印到终端
	//for {
	//	tmp := make([]byte, 1024)
	//	_, err := stdout.Read(tmp)
	//	fmt.Fprint(out, string(tmp))
	//	if err != nil {
	//		break
	//	}
	//}
	reader := bufio.NewReader(stdout)
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Fprint(out, line)
		//fmt.Println(line)
	}

	if err = cmd.Wait(); err != nil {
		return err
	}
	return nil
}

func ExecShell(s string) (string, error) {
	fmt.Println(s)
	cmd := exec.Command("/bin/bash", "-c", s)
	var out bytes.Buffer

	cmd.Stdout = &out
	cmd.Stderr = cmd.Stdout
	err := cmd.Run()
	return out.String(), err
	//err := cmd.Run()
	//if err != nil {
	//	return out.String(), err
	//}
	////fmt.Printf("%s", out.String())
	//return out.String(), nil
}

