package concurrency

import (
	"fmt"
	"runtime"
	"sync"
)

// OneProcessor runs with one logical processor
func OneProcessor() {
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
		fmt.Println()
	}()

	fmt.Println("waiting gorutines")
	wg.Wait()
	fmt.Println("end")
}

// TwoProcessor runs with two logical processor
func TwoProcessor() {
	runtime.GOMAXPROCS(2)
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("waiting gorutines")
	wg.Wait()
	fmt.Println("end")
}

// DefaultProcessor runs default behaviour
func DefaultProcessor() {
	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("start goroutines")

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'A'; char < 'A'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	go func() {
		defer wg.Done()
		for count := 0; count < 3; count++ {
			for char := 'a'; char < 'a'+26; char++ {
				fmt.Printf("%c ", char)
			}
		}
	}()

	fmt.Println("waiting gorutines")
	wg.Wait()
	fmt.Println("end")
}
