package concurrency

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Goroutines() {
	f("direct")

	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("lambda going")

	time.Sleep(time.Second)

	fmt.Println("done")
}

func Channels() {
	// Create a new channel
	// Channels are typed by the value they convey
	messages := make(chan string)

	// Send into a channel with channel <- syntax
	// Here we send a ping to the messages channel from a goroutine
	// using a lambda fun
	go func() { messages <- "ping" }()

	// The <-channel syntax receives a value from the channel.
	// By default sends and receives block until both sender and receiver
	// are ready. This property allows us to wait at the end for the "ping" msg
	// without adding any other synchronization
	msg := <-messages
	fmt.Println(msg)

}

func ChannelBuffering() {

	// A channel of strings buffering UP TO 2 values
	messages := make(chan string, 2)

	// Since its a buffered channel we can send these values
	// into the channel without a corresponding concurrent
	// receive...
	messages <- "buffered"
	messages <- "channel"

	fmt.Println(<-messages)
	fmt.Println(<-messages)

	//This example won't work

	// messages := make(chan string, 2)
	// messages <- "buffered"

	// It gets blocked here
	// messages <- "channel"
}

func worker(done chan bool) {
	fmt.Println("Working..")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func ChannelSync() {
	done := make(chan bool, 1)
	go worker(done)
	// We block until we receive the done message
	<-done
}

// Pings is send only
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// pings is receive only, pongs is send
// we receive a message form pings and send it  to pongs
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func ChannelDirections() {

	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	// We can specify in the functions type args
	// If we only want to send or receive values from a channel
	ping(pings, "passed message")
	pong(pings, pongs)

	fmt.Println(<-pongs)
}

func Select() {
	// Select lets you wait on multiple channel operations
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// We use select to await both values simultaneously
		// We then print when each one arrives
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}

	}
}

func Timeouts() {
	// Note the channel is buffered
	// So the send in the goroutine is nonblocking
	// This is a common patter to prevent leaks
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "result 1"
	}()

	// Our select awaits the result of <-time.After
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(1 * time.Second):
		fmt.Println("timeout 1")
	}

	// We make another channel
	// if we attempted to re use c1 in the select
	// we would obtain the result 1!
	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "result 2"
	}()

	// now the timeout won't be called
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(3 * time.Second):
		fmt.Println("timeout 2")
	}

}

func NonBlockingChannelOperations() {
	// Note this are not the same as buffered with size 1
	// Both Channels need to be ready
	messages := make(chan string)
	signals := make(chan bool)

	// if we executed the following the following we won't need to block both channels
	// messages := make(chan string, 1)
	// signals := make(chan bool, 1)

	// nonblocking receive
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	default:
		fmt.Println("no message received")
	}

	// nonblocking send
	// here msg cannot be sent to the messages channel
	// the channel has no buffer and there is no receiver
	// therefore the default is selected
	msg := "hi"
	messages <- msg
	select {
	case messages <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// We can use multiple cases to implement a multi-way non-blocking select
	// Here we attempt non-blocking receives on both messages and signals
	select {
	case msg := <-messages:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no activity")
	}

}

func ClosingChannels() {
	// notice the buffer size is bigger than the number of jobs
	// what would happen if we change it to a smaller number?
	// the channel will block and don't send all the jobs
	jobs := make(chan int, 5)
	// jobs := make(chan int, 2)
	done := make(chan bool)

	go func() {
		for {
			// more will be false if jobs has been closed
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	// We send from the main() goroutines to a worker
	// when all jobs are send we close the channel
	close(jobs)
	fmt.Println("sent all jobs")
	// we await the worker by blocking with the done channel
	<-done
}

func RangeOverChannels() {
	// Notice the buffer size
	// what happens if we decrease it or send to the queue more values?
	// it will deadlock
	// A difference with the previous example was that a goroutines was executed in background
	// And after its execution we started sending values to our channel
	// In this case this is performed sequentially
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)

	for elem := range queue {
		fmt.Println(elem)
	}
}

func RangeOverChannelsWorker() {
	// Modification combining past examples to obtain similar behavior
	queue := make(chan string, 2)
	done := make(chan bool)
	go func() {
		for elem := range queue {
			fmt.Println(elem)
		}
		done <- true
	}()
	queue <- "one"
	queue <- "two"
	queue <- "three"
	queue <- "four"
	close(queue)
	<-done

}

func Timers() {
	// Timers represent a single event in the future.
	// You tell them how long you want to wait
	// It produces a channel that will notify you
	timer1 := time.NewTimer(2 * time.Second)

	// Blocks on the timer's channel C until it sends a value indicating
	// that the timer fired
	<-timer1.C
	fmt.Println("timer 1 fired")

	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer 2 fired")
	}()
	// A big difference with a sleep is that you can stop a timer
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 stopped")
	}
	time.Sleep(2 * time.Second)
}

func Tickers() {
	// A ticker is a channel that sends values
	ticker := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	// We await values with a ticker.C
	// Notice how we have a main goroutine and a worker one
	// This prevents any deadlock with the locking ticker
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("Tick at", t)
			}
		}
	}()

	time.Sleep(1600 * time.Millisecond)
	// Tickers can be stopped like Timers
	ticker.Stop()
	done <- true
	fmt.Println("ticker stopped")
}

func Worker(id int, jobs <-chan int, results chan<- int) {
	// We have seen this pattern before, range over jobs
	// When jobs get closed it stops
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}

}

func WorkerPools() {
	// We make two channels, our pool of workers gets executed in one
	// And collect results in another channel
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// We start 3 workers that get blocked because there is no jobs
	for w := 1; w <= 3; w++ {
		go Worker(w, jobs, results)
	}

	// Here we send 5 jobs to our workers then close the jobs channel
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	// Finally we collect the results of the work
	// This also ensures that the worker goroutines have finished
	// An alternative wait would be to use WaitGroup
	for a := 1; a <= numJobs; a++ {
		<-results
	}

}

func WorkerWait(id int, wg *sync.WaitGroup) {
	// We pass our WaitGroup pointer
	// On return we notify that we are done
	defer wg.Done()
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("worker %d done\n", id)
}

func WaitGroups() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		// each goroutine increments the waiting group
		// since we have 5 of them, our waiting group will wait 5 done "signals"
		wg.Add(1)
		go WorkerWait(i, &wg)
	}
	wg.Wait()
}

// Note I did not include the results for this example
// We simply pass the wg, defer the notification and consume the range of jobs
func WorkerWaitExtended(id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for j := range jobs {
		fmt.Println("worker", id, "started job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
	}

}

func WaitGroupsExtended() {
	const numJobs = 5
	jobs := make(chan int, numJobs)

	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go WorkerWaitExtended(w, jobs, &wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()

}

func RateLimiting() {

	requests := make(chan int, 5)
	// Suppose we want to limit this requests
	// We serve them all at the same time then close the channel
	for i := 1; i <= 5; i++ {
		requests <- i
	}
	close(requests)

	// Limiter channel will receive a value every 200 ms
	limiter := time.Tick(200 * time.Millisecond)

	// By blocking the limiter, we limit ourselves to 1 req every 200ms
	for req := range requests {
		<-limiter
		fmt.Println("request", req, time.Now())
	}

	// We may want to allow bursts of requests
	// This channel allows bursts of 3 messages
	burstyLimiter := make(chan time.Time, 3)
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// Every 200 ms we try to add a message to burstyLimiter up to its limit of 3
	go func() {
		for t := range time.Tick(200 * time.Millisecond) {
			burstyLimiter <- t
		}
	}()

	// Now we simulate 5 more incoming requests
	burstyRequests := make(chan int, 5)
	for i := 1; i <= 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)
	// The first 3 messages will benefit form burstyLimiter
	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("request", req, time.Now())
	}
}

func AtomicCounters() {
	// atomic counters accessed by multiple goroutines
	var ops uint64
	var wg sync.WaitGroup
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for c := 0; c < 1000; c++ {
				// atomic add with the pointer & syntax
				atomic.AddUint64(&ops, 1)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("ops:", ops)
}

func Mutexes() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}

	var readOps uint64
	var writeOps uint64

	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)

	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func StatefulGoroutines() {
	// In the previous example we used explicit locking with mutexes.
	// This channel-based approach aligns with Go’s ideas of sharing memory
	// by communicating and having each piece of data owned by exactly 1 goroutine.
	var readOps uint64
	var writeOps uint64

	// In this example our state will be owned by a single goroutine
	// This guarantees that data is never corrupted with concurrent access.
	// In order to read or write, other goroutines will send messages to the
	// owning goroutine and receive corresponding replies
	// These readOps and writeOps structs encapsulate those requests
	// and a way for the owning goroutine to respond

	// The reads and writes channels will be used by other goroutines to
	// issue read and write requests
	reads := make(chan readOp)
	writes := make(chan writeOp)

	// Here is the goroutine that owns the state, which is a map, its private
	// to this goroutine.
	// The goroutine repeatedly selects on reads and writes channels, responding
	// to requests as they arrive. A response is executed by first performing
	// the requested operation and then sending a value on the response channel
	// resp to indicate success
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddUint64(&readOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddUint64(&writeOps, 1)
				time.Sleep(time.Millisecond)
			}
		}()
	}

	time.Sleep(time.Second)

	readOpsFinal := atomic.LoadUint64(&readOps)
	fmt.Println("readOps:", readOpsFinal)
	writeOpsFinal := atomic.LoadUint64(&writeOps)
	fmt.Println("writeOps:", writeOpsFinal)
	// For this particular case the goroutine-based approach was a bit more
	// involved than the mutex-based one. It might be useful in certain cases though,
	// for example where you have other channels involved or when managing multiple
	// such mutexes would be error-prone. You should use whichever approach feels
	// most natural

}

// BadThreadBroadcastPattern for is blocking the cpu
// Active waiting
func BadThreadBroadcastPattern() {
	rand.Seed(time.Now().UnixNano())
	count := 0
	finished := 0
	var mu sync.Mutex

	for i := 0; i < 10; i++ {
		go func() {
			vote := rand.Float32() > 0.5
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
		}()
	}

	// BAD
	for {
		mu.Lock()
		if count >= 5 || finished == 10 {
			break
		}
		mu.Unlock()
	}
	if count >= 5 {
		fmt.Println("win")
	} else {
		fmt.Println("lose")
	}
	mu.Unlock()
}

// CondThreadBroadcastPattern uses a cond and broadcast
func CondThreadBroadcastPattern() {
	rand.Seed(time.Now().UnixNano())
	count := 0
	finished := 0
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func() {
			vote := rand.Float32() > 0.5
			mu.Lock()
			defer mu.Unlock()
			if vote {
				count++
			}
			finished++
			// Notice we first BROADCAST
			// LATER WE unlock!
			cond.Broadcast()
		}()
	}

	// First get the lock
	mu.Lock()
	for count < 5 && finished != 10 {
		// White with cond while condition unmet
		cond.Wait()
	}
	// Now we own the lock from this point onwards
	if count >= 5 {
		fmt.Println("win")
	} else {
		fmt.Println("lose")
	}
	mu.Unlock()
}
