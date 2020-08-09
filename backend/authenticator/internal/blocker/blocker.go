package blocker

import (
	"sync"
	"time"
)

type Blocker interface {
	CheckBlock(key string)
}

type blocker struct {
	tryCount     map[string]int64
	tryMutex     map[string]*sync.Mutex
	generalMutex *sync.RWMutex
	period       time.Duration
	cleanPeriod  time.Duration
}

func NewBlocker(period time.Duration, cleanPeriod time.Duration) Blocker {
	b := &blocker{
		tryCount:     make(map[string]int64),
		tryMutex:     make(map[string]*sync.Mutex),
		generalMutex: &sync.RWMutex{},
		period:       period,
		cleanPeriod:  cleanPeriod,
	}
	go b.clean()
	return b
}

func (b *blocker) CheckBlock(key string) {
	b.generalMutex.RLock()
	tryMutex, ok := b.tryMutex[key]
	if !ok {
		tryMutex = &sync.Mutex{}
		b.generalMutex.RUnlock()
		b.generalMutex.Lock()
		b.tryMutex[key] = tryMutex
		b.generalMutex.Unlock()
		b.generalMutex.RLock()
	}
	defer b.generalMutex.RUnlock()
	tryMutex.Lock()
	defer tryMutex.Unlock()
	count := b.tryCount[key]
	b.tryCount[key] = count + 1
	time.Sleep(b.period * time.Duration(count))
}

func (b *blocker) clean() {
	t := time.NewTicker(b.cleanPeriod)
	for {
		<-t.C
		deleteMutexes := make([]string, len(b.tryMutex))
		for k, m := range b.tryMutex {
			go func(key string, mutex *sync.Mutex) {
				mutex.Lock()
				defer mutex.Unlock()
				if count, ok := b.tryCount[key]; ok {
					if count == 0 {
						delete(b.tryCount, key)
						deleteMutexes = append(deleteMutexes, key)
					} else {
						b.tryCount[key] = count - 1
					}
				}
			}(k, m)
		}
		b.generalMutex.Lock()
		for _, key := range deleteMutexes {
			delete(b.tryMutex, key)
		}
		b.generalMutex.Unlock()
	}
}
