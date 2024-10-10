package gUtils

import (
	"log"
	"runtime"
)

func TryRecoverCatch() {
	if err := recover(); err != nil {
		var buf [4096]byte
		n := runtime.Stack(buf[:], false)
		log.Printf("Catch Stack =>\n %s\nReason: %v\n", string(buf[:n]), err)

	}
}

func SafeGo(f func()) {
	go run(f)
}

func run(f func()) {
	defer TryRecoverCatch()
	f()
}
