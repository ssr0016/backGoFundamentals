package utils

import (
	"fmt"
	"time"
)

// What is select?
// The select statement is used to choose from multiple send/receive channel operations. The select statement blocks until one of the send/receive operations is ready. If multiple operations are ready, one of them is chosen at random. The syntax is similar to switch except that each of the case statements will be a channel operation. Let’s dive right into some code for better understanding.

func server1(ch chan string) {
	time.Sleep(6 * time.Second)
	ch <- "from server 1"
}

func server2(ch chan string) {
	time.Sleep(3 * time.Second)
	ch <- "from server 2"
}

func SelectExample() {
	output1 := make(chan string)
	output2 := make(chan string)
	go server1(output1)
	go server2(output2)
	select {
	case s1 := <-output1:
		fmt.Println(s1)
	case s2 := <-output2:
		fmt.Println(s2)
	}
}

// ===================
func processSelect(ch chan string) {
	time.Sleep(10500 * time.Millisecond)
	ch <- "process successful"
}

func SelectDefaultCase() {
	ch := make(chan string)
	go processSelect(ch)
	for {
		time.Sleep(1000 * time.Millisecond)
		select {
		case v := <-ch:
			fmt.Println("received value: ", v)
			return
		default:
			fmt.Println("no value received")
		}

	}
}

func SelectDeadlockAndDefaultAndDefaultCase() {
	ch := make(chan string)
	select {
	case <-ch:
	default:
		fmt.Println("default case executed")
	}
}

func SelectDeadlockAndDefaultAndDefaultCase1() {
	var ch chan string
	select {
	case v := <-ch:
		fmt.Println("received value: ", v)
	default:
		fmt.Println("default case executed")
	}
}

func server3(ch chan string) {
	ch <- "server 2"
}

func server4(ch chan string) {
	ch <- "server 2"
}

func RandomSelection() {
	output3 := make(chan string)
	output4 := make(chan string)
	go server3(output3)
	go server4(output4)
	time.Sleep(1 * time.Second)
	select {
	case s1 := <-output3:
		fmt.Println(s1)
	case s2 := <-output4:
		fmt.Println(s2)
	}
}
