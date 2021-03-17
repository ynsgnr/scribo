package blocker

import (
	"errors"
	"sync"
	"time"
)

var ErrLongWaitPeriod = errors.New("ErrLongWaitPeriod")

type Blocker interface {
	//Checks the key and sleeps if required to thorttle
	//if sleep amount is bigger than maxPeriod it returns ErrLongWaitPeriod
	CheckBlock(key string) error
}

type blocker struct {
	tryCount           map[string]int64
	mutex              *sync.Mutex
	period             time.Duration
	cleanPeriod        time.Duration
	maxWaitPeriod      time.Duration
	throttleAfterTries int64
}

func NewBlocker(period, cleanPeriod, maxWaitPeriod time.Duration, throttleAfterTries int64) Blocker {
	b := &blocker{
		tryCount:           make(map[string]int64),
		mutex:              &sync.Mutex{},
		period:             period,
		cleanPeriod:        cleanPeriod,
		maxWaitPeriod:      maxWaitPeriod,
		throttleAfterTries: throttleAfterTries,
	}
	go b.clean()
	return b
}

func (b *blocker) CheckBlock(key string) error {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	tries := b.tryCount[key]
	b.tryCount[key]++
	if tries >= b.throttleAfterTries {
		timeout := b.period * time.Duration(tries)
		if timeout >= b.maxWaitPeriod {
			return ErrLongWaitPeriod
		}
		time.Sleep(timeout)
	}
	return nil
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
