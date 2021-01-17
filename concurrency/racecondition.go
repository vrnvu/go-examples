package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

func RaceConditionDetector() {
	var counter int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			value := counter
			// Yield the thread and be placed back in queue.
			runtime.Gosched()
			value++
			counter = value
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 2; i++ {
			value := counter
			// Yield the thread and be placed back in queue.
			runtime.Gosched()
			value++
			counter = value
		}
	}()
	wg.Wait()
	fmt.Println("Final counter:", counter)
}
