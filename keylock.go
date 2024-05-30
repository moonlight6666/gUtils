package gUtils

import "sync"

type KeyLock struct {
	mu    sync.Mutex
	locks map[string]*sync.Mutex
}

var GlobalKeyLock = NewKeyLock()

func NewKeyLock() *KeyLock {
	return &KeyLock{
		locks: make(map[string]*sync.Mutex),
	}
}

func (kl *KeyLock) Lock(key string) {
	kl.mu.Lock()
	if _, exists := kl.locks[key]; !exists {
		kl.locks[key] = &sync.Mutex{}
	}
	kl.mu.Unlock()

	kl.locks[key].Lock()
}

func (kl *KeyLock) TryLock(key string) bool {
	kl.mu.Lock()
	if _, exists := kl.locks[key]; !exists {
		kl.locks[key] = &sync.Mutex{}
	}
	kl.mu.Unlock()

	return kl.locks[key].TryLock()
}

func (kl *KeyLock) Unlock(key string) {
	kl.mu.Lock()
	defer kl.mu.Unlock()

	if lock, exists := kl.locks[key]; exists {
		lock.Unlock()
	}
}
