package gUtils

// 文件锁

import (
	"os"
	"syscall"
)

type fileLock struct {
	f        *os.File
	filename string
}

func NewFileLock(file string) (*fileLock, error) {
	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	fl := &fileLock{
		f:        f,
		filename: file,
	}
	if err = syscall.Flock(int(fl.f.Fd()), syscall.LOCK_EX|syscall.LOCK_NB); err != nil {
		return nil, err
	}

	return fl, nil
}

func (fl *fileLock) Unlock() error {
	defer fl.f.Close()
	os.Remove(fl.filename)
	return syscall.Flock(int(fl.f.Fd()), syscall.LOCK_UN)
}
