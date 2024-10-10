package gUtils

import (
	"fmt"
	"sync"
)

type (
	noticeLimitData struct {
		Times    int
		LastTime int
	}
	NoticeLimitConfig struct {
		Times       int
		LimitSecond int
	}
	noticeLimit struct {
		sync.Map
		config []NoticeLimitConfig
	}
)

var DefaultNoticeLimit = NewDefaultNoticeLimit()

func NewNoticeLimit(config []NoticeLimitConfig) *noticeLimit {
	return &noticeLimit{
		config: config,
	}
}

func NewDefaultNoticeLimit() *noticeLimit {
	return NewNoticeLimit([]NoticeLimitConfig{
		{2, MinuteSecond * 10},
		{3, MinuteSecond * 30},
		{4, HourSecond * 1},
		{5, HourSecond * 3},
		{6, HourSecond * 6},
		{7, HourSecond * 12},
		{8, HourSecond * 24},
	})
}

func (nl *noticeLimit) IsNotice(notice string) bool {
	times := 1
	now := GetTimestamp()
	if v, ok := nl.Load(notice); ok {
		e := v.(noticeLimitData)
		times = e.Times + 1
		lastTime := e.LastTime

		limitSecond := 0
		for _, v := range nl.config {
			if times >= v.Times {
				limitSecond = Max(limitSecond, v.LimitSecond)
			}
		}
		fmt.Println("limitTime:", times, limitSecond, lastTime)
		if now <= lastTime+limitSecond {
			return false
		}
	}
	nl.updateData(notice, times, now)
	return true
}

func (nl *noticeLimit) Reset(notice string) {
	nl.Delete(notice)
}

func (nl *noticeLimit) TryReset(notice string) bool {
	if _, ok := nl.Load(notice); ok {
		nl.Reset(notice)
		return true
	}
	return false
}

func (nl *noticeLimit) updateData(notice string, times int, time int) {
	nl.Store(notice, noticeLimitData{
		Times:    times,
		LastTime: time,
	})
}
