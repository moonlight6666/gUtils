package gUtils

import (
	"log"
	"runtime"
	"sync"
)

type waitGroup struct {
	mux       sync.Mutex
	pending   int
	isClose   bool
	waitGroup *sync.WaitGroup
}

var DefaultWaitGroup = NewWaitGroup()

func NewWaitGroup() *waitGroup {
	return &waitGroup{
		waitGroup: new(sync.WaitGroup),
	}
}

func (*waitGroup) Go(f func()) {
	go run(f)
}

func (wg *waitGroup) Add() {
	wg.waitGroup.Add(1)
	wg.mux.Lock()
	wg.pending++
	wg.mux.Unlock()
}

func (wg *waitGroup) Done() {
	wg.waitGroup.Done()
	wg.mux.Lock()
	wg.pending--
	wg.mux.Unlock()
}

func (wg *waitGroup) WaitClose() {
	wg.mux.Lock()
	wg.isClose = true
	log.Printf("Waiting %d task close...\n", wg.pending)
	wg.mux.Unlock()

	wg.waitGroup.Wait()
	log.Printf("All close.\n")
}

func (wg *waitGroup) run(f func()) {
	wg.Add()

	defer func() {
		if err := recover(); err != nil {
			var buf [4096]byte
			n := runtime.Stack(buf[:], false)
			log.Printf("waitGroupError Catch Stack =>\n %s\nReason: %v\n", string(buf[:n]), err)
		}
		wg.Done()
	}()

	f()
}
