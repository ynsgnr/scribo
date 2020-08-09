package blocker

import (
	"sync"
	"testing"
	"time"
)

func TestBlocker(t *testing.T) {
	wg := &sync.WaitGroup{}

	period := time.Millisecond
	cleanPeriod := time.Millisecond
	b := &blocker{
		tryCount:     make(map[string]int64),
		tryMutex:     make(map[string]*sync.Mutex),
		generalMutex: &sync.RWMutex{},
		period:       period,
		cleanPeriod:  cleanPeriod,
	}
	go b.clean()
	key := "testKey1"
	key2 := "testKey2"
	tryCount1 := 3
	tryCount2 := 2
	for i := 0; i < tryCount1; i++ {
		wg.Add(1)
		expectedTime := period * time.Duration(i)
		go CheckBlocker(t, b, key, wg, expectedTime)
	}
	for i := 0; i < tryCount2; i++ {
		wg.Add(1)
		expectedTime := period * time.Duration(i)
		go CheckBlocker(t, b, key2, wg, expectedTime)
	}
	wg.Wait()
	time.Sleep(cleanPeriod * time.Duration(tryCount1))
	for i := 0; i < tryCount1; i++ {
		wg.Add(1)
		expectedTime := period * time.Duration(i)
		go CheckBlocker(t, b, key, wg, expectedTime)
	}
	wg.Wait()
}

func CheckBlocker(t *testing.T, b *blocker, key string, wg *sync.WaitGroup, minimumExpectedTime time.Duration) {
	start := time.Now()
	b.CheckBlock(key)
	timePassed := time.Since(start) + time.Microsecond
	if timePassed-minimumExpectedTime < 0 {
		t.Errorf("For %s should have taken more than %s, took %s", key, minimumExpectedTime.String(), timePassed.String())
	}
	wg.Done()
}
