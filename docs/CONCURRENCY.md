# 1. Concurrency Patterns in Go


> We all know that go comes with great concurrency tools like go routines and channels. But is this really everything? This session starts with an overview of common concurrency patterns and ends with best practices on lockless programming that won't let your head explode.

[Arne Claus - Concurrency Patterns in Go](https://www.youtube.com/watch?v=YEKjSzIwAdA): Notes from the first half aprox.

- [1. Concurrency Patterns in Go](#1-concurrency-patterns-in-go)
- [2. Concurrency](#2-concurrency)
  - [2.1. Concurrency in detail](#21-concurrency-in-detail)
- [3. Communication Sequential Processes (CSP)](#3-communication-sequential-processes-csp)
- [4. Go's concurrency toolset](#4-gos-concurrency-toolset)
  - [4.1. Channels](#41-channels)
    - [4.1.1. Blocking channels](#411-blocking-channels)
    - [4.1.2. Blocking breaks concurrency](#412-blocking-breaks-concurrency)
    - [4.1.3. Closing channels](#413-closing-channels)
  - [4.2. Select](#42-select)
    - [4.2.1. Making channels non-blocking](#421-making-channels-non-blocking)
  - [4.3. Shape your data flow](#43-shape-your-data-flow)
    - [4.3.1. Fan-out](#431-fan-out)
    - [4.3.2. Turnout](#432-turnout)
    - [4.3.3. Quit channel](#433-quit-channel)

# 2. Concurrency

> Concurrency is about design

- **Design** your program as a collection of independent processes
- **Design** these processes to *eventually* run in parallel
- **Design** your code so that the outcome is always the same

## 2.1. Concurrency in detail

- Group code (and data) by identifing independent tasks
- No race conditions
- No deadlocks
- More workers - faster execution

# 3. Communication Sequential Processes (CSP)

- Each rpocess is built for sequential execution
- **Data is communicated between rpocesses via channels. No shared state!**
- Scale by adding more processes

# 4. Go's concurrency toolset

- go routines
- channels
- select
- sync package

## 4.1. Channels

- Bucket chain, queue
- A channel has 3 components:
  - **sender**
  - buffer
  - **receiver**
- The buffer is optional

### 4.1.1. Blocking channels

```go
unbuffered := make(chan int)

// Only sender blocks
a := <-unbuffered
// Only receiver blocks
unbuffered <- 1

// A coruoutine? It blocks until the sender
go func() {<-unbuffered}()
unbuffered <- 1
```

Another example with a buffered channel:

```go
unbuffered := make(chan int, 1)

// sender blocks
a := <-unbuffered
// receiver DOES NOT blocks
unbuffered <- 1
// receiver blocks because we have reach the limit of the buffer
unbuffered <- 1
```

### 4.1.2. Blocking breaks concurrency

- Remember?
  - No deadlocks
  - More workers = fsater execution
- Blocking can lead to deadlocks
- Blocking can prevent scaling

### 4.1.3. Closing channels

- Close sends a special "closed" message
- The receiver at some point see "closed"
- If you try to send more: panic!

```go
c := make(chan int)
close(c)
fmt.Println(<-c) // receive and print
// What is printed?
// 0, false
// 0 is the zero value, false means no more date
```

The receiver always knows that the channel is closed

The sender never does

That is why you always close a channel from the sender side

## 4.2. Select

- Like a switch statement on channel operations
- The order of cases doesn't matter at all
- There is a default case
- The first non-blocking case is chosen (send and/or receive)

### 4.2.1. Making channels non-blocking

```go
func TryReceive(c <-chan int) (data int, more, ok bool) {
    select {
        case data, more = <-c:
            return data, more, true
        default:
            return 0, true, false
    }
}
```

In some cases we need a timer, we want to wait for the next message

```go
func TryReceiveWithTimeout(c <-chan int, duration time.Duration) (data int, more, ok bool) {
    select {
        case data, more = <-c:
            return data, more, true
        case <- time.After(duration):
            return 0, true, false
    }
}
```

## 4.3. Shape your data flow

- Channels are streams of data
- Dealing with multiple streams is the true power of select

``` 
Fan-out
   / =>
==>- =>
   \ =>
```

``` 
Funnel
=>\
=>---=>
=>/
```

``` 
Turnout
=>\    / =>
=>--=>-- =>
=>/    \ =>
```

### 4.3.1. Fan-out

We have a set of workers and we want to send data when they have availability
```go
func Fanout(in <-chan int, OutA, OutB chan int) {
    for data := range In {
        select {
            case OutA <- data:
            case OutB <- data:
        }
    }
} 
```

### 4.3.2. Turnout

We have to get data from multiple channels.

```go
func Turnout(InA, InB <-chan int, OutA, OutB chan int) {
    for {
        select {
            case data, more = <-InA:
            case data, more = <-InB:
        }
        if !more {
            // if we close one channel, the select will always
            // chose that 
            // so when any of the channels is closed we quit
            return
        }
        select {
            case OutA <- data:
            case OutB <- data:
        }
    }
}

```

### 4.3.3. Quit channel

In order to fix the closing channel issue from the previous example, we add a Quit channel to delegate the closing
```go
func Turnout(Quit <-chan int, InA, InB, OutA, OutB chan int) {
    for {
        select { 
            case data = <-InA:
            case data = <-InB:
        }
        case <- Quit {
            // remember close generates a message
            close(InA)
            close(InB)

            // Flush the remaining data
            Fanout(InA, OutA, OutB) 
            Fanout(InB, OutA, OutB)
            return
        }
    }
}
```
