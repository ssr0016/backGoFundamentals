package utils

import (
	"fmt"
	"time"
)

// Channels can be thought of as pipes using which Goroutines communicate. Similar to how water flows from one end to another in a pipe, data can be sent from one end and received from the other end using channels.

func DeclaringChannels() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T", a)
	}
}

func hello3(done chan bool) {
	fmt.Println("Hello go goroutine is going to sleep")
	time.Sleep(4 * time.Second)
	fmt.Println("Hello go goroutine awake and going to write to done")
	done <- true
}

func ChannelExampleProgram() {
	done := make(chan bool)
	fmt.Println("Main going to call hello3 go goroutine")
	go hello3(done)
	<-done
	fmt.Println("Main received data")
}

// =================================================

func calcSquares(number int, squareop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit
		number /= 10
	}
	squareop <- sum
}

func calcCubes(number int, cubeop chan int) {
	sum := 0
	for number != 0 {
		digit := number % 10
		sum += digit * digit * digit
		number /= 10
	}

	cubeop <- sum
}

func AnotherExampleForChannels() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares(number, sqrch)
	go calcCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}

func sendData(sendch chan<- int) {
	sendch <- 10
}

func UndirectionalChannels() {
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}

func producer(chnl chan int) {

	for i := 0; i < 10; i++ {
		chnl <- i
	}
	close(chnl)
}

func ClosingChannelsAndForRangeLoopsOnChannels() {
	ch := make(chan int)
	go producer(ch)
	for {
		v, ok := <-ch
		if ok == false {
			break
		}
		fmt.Println("Received ", v, ok)
	}

	// Alternative
	// ch := make(chan int)
	// go producer(ch)
	// for v := range ch {
	// 	fmt.Println("Received ",v)
	// }

	// ============================================

}

func digits(number int, dchnl chan int) {
	for number != 0 {
		digit := number % 10
		dchnl <- digit
		number /= 10
	}
	close(dchnl)
}

func calcSquares1(number int, squareop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit
	}
	squareop <- sum
}

func calcCubes1(number int, cubeop chan int) {
	sum := 0
	dch := make(chan int)
	go digits(number, dch)
	for digit := range dch {
		sum += digit * digit * digit
	}
	cubeop <- sum
}

func AnotherExampleForChannels1() {
	number := 589
	sqrch := make(chan int)
	cubech := make(chan int)
	go calcSquares1(number, sqrch)
	go calcCubes1(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)
}
