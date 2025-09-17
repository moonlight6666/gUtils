package gUtils

import (
	"fmt"
	"log"
	"sync"
)

type monitorCostTime struct {
	startTime    int64
	endTime      int64
	functionName string
}

var pool = &sync.Pool{
	New: func() interface{} {
		return &monitorCostTime{}
	},
}

func NewMonitorCostTime(name string) *monitorCostTime {
	p := pool.Get()
	pool.Put(p)

	p1 := p.(*monitorCostTime)
	p1.functionName = name
	return p1
}

func (s *monitorCostTime) Start() {
	log.Printf("执行 [%s] 开始\n", s.functionName)
	s.startTime = MilliTimestamp()
}

func (s *monitorCostTime) Stop() {
	log.Printf("执行 [%s] 完成, 耗时: %s\n", s.functionName, formatTime(s.UseTime()))
}

func (s *monitorCostTime) UseTime() int64 {
	s.endTime = MilliTimestamp()
	useTime := s.endTime - s.startTime
	return useTime
}

func (s *monitorCostTime) FormatUseTime() string {
	return formatTime(s.UseTime())
}

func formatTime(t int64) string {
	switch {
	case t < 1000:
		return fmt.Sprintf("%dms", t)
	default:
		return fmt.Sprintf("%.2fs", float64(t)/1000)
	}
}
