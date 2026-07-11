package gUtils

import (
	"os"
	"os/signal"
	"syscall"
)

// 等待信号退出
func WaitShutdownSignal() os.Signal {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGUSR1, syscall.SIGUSR2)
	return <-c
}
