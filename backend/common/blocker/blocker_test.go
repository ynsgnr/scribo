package blocker

import (
	"sync"
	"testing"
	"time"
)

type BlockerTest struct {
	t                   *testing.T
	b                   Blocker
	key                 string
	minimumExpectedTime time.Duration
	err                 error
}

func TestBlocker(t *testing.T) {
	period := time.Millisecond
	cleanPeriod := 20 * time.Millisecond
	maxWaitPeriod := 100 * time.Millisecond
	tryCount := int64(2)
	key1 := "testKey1"
	key2 := "testKey2"
	tryCount1 := 5
	tryCount2 := 4

	wg := &sync.WaitGroup{}

	channel1 := make(chan BlockerTest, tryCount1)
	wg.Add(1)
	go RunTest(channel1, wg)
	channel2 := make(chan BlockerTest, tryCount2)
	wg.Add(1)
	go RunTest(channel2, wg)

	b := NewBlocker(period, cleanPeriod, maxWaitPeriod, tryCount)

	for i := 0; i < tryCount1; i++ {
		expectedTime := period * time.Duration(i)
		if int64(i) < tryCount {
			expectedTime = time.Duration(0)
		}
		channel1 <- BlockerTest{
			t:                   t,
			b:                   b,
			key:                 key1,
			minimumExpectedTime: expectedTime,
		}
	}
	for i := 0; i < tryCount2; i++ {
		expectedTime := period * time.Duration(i)
		if int64(i) < tryCount {
			expectedTime = time.Duration(0)
		}
		channel2 <- BlockerTest{
			t:                   t,
			b:                   b,
			key:                 key2,
			minimumExpectedTime: expectedTime,
		}
	}
	time.Sleep(cleanPeriod * time.Duration(tryCount1))
	for i := 0; i < tryCount1; i++ {
		expectedTime := period * time.Duration(i)
		if int64(i) < tryCount {
			expectedTime = time.Duration(0)
		}
		channel1 <- BlockerTest{
			t:                   t,
			b:                   b,
			key:                 key1,
			minimumExpectedTime: expectedTime,
		}
	}
	close(channel1)
	close(channel2)
	wg.Wait()
}

func TestBlockerMaxWait(t *testing.T) {
	period := time.Millisecond
	cleanPeriod := 100 * time.Millisecond
	maxWaitPeriod := 5 * time.Millisecond
	tryCount := int64(2)
	key := "testKey"
	tryCount1 := 5

	wg := &sync.WaitGroup{}

	channel := make(chan BlockerTest, tryCount)
	wg.Add(1)
	go RunTest(channel, wg)

	b := NewBlocker(period, cleanPeriod, maxWaitPeriod, tryCount)

	for i := 0; i < tryCount1; i++ {
		expectedTime := period * time.Duration(i)
		if int64(i) < tryCount {
			expectedTime = time.Duration(0)
		}
		channel <- BlockerTest{
			t:                   t,
			b:                   b,
			key:                 key,
			minimumExpectedTime: expectedTime,
		}
	}
	channel <- BlockerTest{
		t:   t,
		b:   b,
		key: key,
		err: ErrLongWaitPeriod,
	}
	close(channel)
	wg.Wait()
}

func RunTest(channel chan BlockerTest, wg *sync.WaitGroup) {
	for test := range channel {
		start := time.Now()
		err := test.b.CheckBlock(test.key)
		if err != test.err {
			test.t.Errorf("For %s: Expected error %v got %v", test.key, test.err, err)
		}
		timePassed := time.Since(start) + time.Microsecond
		if timePassed-test.minimumExpectedTime < 0 {
			test.t.Errorf("For %s should have taken more than %s, took %s", test.key, test.minimumExpectedTime.String(), timePassed.String())
		}
	}
	wg.Done()
}
