package utils

import (
	"fmt"
	"sync"
)

// Mutex
// A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section of code at any point in time to prevent race conditions from happening.

// Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. Any code that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.

// mutex.Lock()
// x = x + 1
// mutex.Unlock()

var x = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func ProgramWithARaceCondition() {
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}

	w.Wait()
	fmt.Println("final value of x", x)
}

func increment1(wg *sync.WaitGroup, m *sync.Mutex) {
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func SolvingTheRaceConditionUsingAMutex() {
	var w sync.WaitGroup
	var m sync.Mutex
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment1(&w, &m)
	}
	w.Wait()
	fmt.Println("final value of x", x)

}

func increment2(wg *sync.WaitGroup, ch chan bool) {
	ch <- true
	x = x + 1
	<-ch
	wg.Done()
}

func SolvingTheRaceConditionUsingChannel() {
	var w sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment2(&w, ch)
	}
	w.Wait()
	fmt.Println("final value of x", x)
}
