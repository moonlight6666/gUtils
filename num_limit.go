package gUtils

import (
	"fmt"
	"sync"
)

type numLimit struct {
	maxNum int
	num    int
	mutex  sync.Mutex
}

func NewNumLimit(maxNum int) *numLimit {
	return &numLimit{
		maxNum: maxNum,
	}
}

func (n *numLimit) TryAdd() bool {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	if n.num < n.maxNum {
		n.num++
		fmt.Println("numLimit TryAdd:", n.num)
		return true
	}
	return false
}

func (n *numLimit) Done() {
	n.mutex.Lock()
	defer n.mutex.Unlock()
	n.num--
	fmt.Println("numLimit Done:", n.num)
	if n.num < 0 {
		panic(fmt.Sprintf("numLimit fetal error:%d", n.num))
	}
}
