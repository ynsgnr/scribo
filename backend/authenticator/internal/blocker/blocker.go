package blocker

import (
	"sync"
	"time"
)

type Blocker interface {
	CheckBlock(key string)
}

type blocker struct {
	tryCount           map[string]int64
	mutex              *sync.Mutex
	period             time.Duration
	cleanPeriod        time.Duration
	throttleAfterTries int64
}

func NewBlocker(period time.Duration, cleanPeriod time.Duration, throttleAfterTries int64) Blocker {
	b := &blocker{
		tryCount:           make(map[string]int64),
		mutex:              &sync.Mutex{},
		period:             period,
		cleanPeriod:        cleanPeriod,
		throttleAfterTries: throttleAfterTries,
	}
	go b.clean()
	return b
}

func (b *blocker) CheckBlock(key string) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	tries := b.tryCount[key]
	b.tryCount[key]++
	if tries >= b.throttleAfterTries {
		timeout := b.period * time.Duration(tries)
		time.Sleep(timeout)
	}
}

func (b *blocker) clean() {
	t := time.NewTicker(b.cleanPeriod)
	for {
		<-t.C
		for k, v := range b.tryCount {
			go func(key string, v int64) {
				b.mutex.Lock()
				defer b.mutex.Unlock()
				b.tryCount[key] = v - 1
				if v == 0 {
					delete(b.tryCount, key)
				}
			}(k, v)
		}
	}
}
