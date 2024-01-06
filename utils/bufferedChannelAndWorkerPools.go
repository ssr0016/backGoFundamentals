package utils

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// What are buffered channels?

// All the channels we discussed in the previous tutorial were basically unbuffered. As we discussed in the channels tutorial in detail, sends and receives to an unbuffered channel are blocking.

// It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. Similarly receives from a buffered channel are blocked only when the buffer is empty.

// Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.

// Here is an example of a buffered channel:

// ch := make(chan type, capacity)

func BufferedChannelAndWorkerPoolsExample() {
	ch := make(chan string, 2)
	ch <- "hello"
	ch <- "world"
	fmt.Println(<-ch)
	time.Sleep(4 * time.Second)
	fmt.Println(<-ch)
}

func write(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to channel")
	}
	close(ch)
}

func WriteImplementation() {
	ch := make(chan int, 2)
	go write(ch)
	time.Sleep(2 * time.Second)
	for v := range ch {
		fmt.Println("read value", v, "from channel")
		time.Sleep(2 * time.Second)
	}
}

func Deadlock() {
	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	ch <- "steve"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func ClosingBufferedChannels() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	n, open := <-ch
	fmt.Printf("Received: %d open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d open: %t\n", n, open)
	n, open = <-ch
	fmt.Printf("Received: %d open: %t\n", n, open)
}

func RangeCloringBufferedChannels() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)
	for n := range ch {
		fmt.Println("Received ", n)
	}
}

func LengthVsCapacity() {
	ch := make(chan string, 3)
	ch <- "hello"
	ch <- "world"
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))
	fmt.Println("read value", <-ch)
	fmt.Println("new length is", len(ch))
}

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("started Goroutine ", i)
	time.Sleep(2 * time.Second)
	fmt.Printf("Goroutine %d ended\n", i)
	wg.Done()
}

func WaitGroup() {
	no := 3
	var wg sync.WaitGroup
	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All go routines finished executing")
}

type Job struct {
	id       int
	randomno int
}

type Result struct {
	job         Job
	sumofdigits int
}

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

func digitsNum(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}

	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digitsNum(job.randomno)}
		results <- output
	}
	wg.Done()

}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomno := rand.Intn(999)
		job := Job{i, randomno}
		jobs <- job
	}

	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Job id %d, input random no %d, sum of digits %d\n", result.job.id, result.job.randomno, result.sumofdigits)
	}
	done <- true
}

func WorkerPoolExample() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<-done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken", diff.Seconds(), "seconds")
}
