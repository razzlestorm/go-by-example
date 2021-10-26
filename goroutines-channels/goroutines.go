package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for ii := 0; ii < 3; ii++ {
		fmt.Println(from, ":", ii)
	}
}

func main() {
	// Suppose we have a function call f(s). Here’s how we’d call that in the usual way, running it synchronously.
	f("direct")

	// To invoke this function in a goroutine, use go f(s). This new goroutine will execute concurrently with the calling one.
	go f("goroutine")

	// You can also start a goroutine for an anonymous function call.
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// Our two function calls are running asynchronously in separate goroutines now. Wait for them to finish (for a more robust approach, use a WaitGroup).
	time.Sleep(time.Second)
	fmt.Println("done")

	//When we run this program, we see the output of the blocking call first, then the output of the two goroutines. The goroutines’ output may be interleaved, because goroutines are being run concurrently by the Go runtime.

	// Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.

	//Create a new channel with make(chan val-type). Channels are typed by the values they convey.
	messages := make(chan string)

	//Send a value into a channel using the channel <- syntax. Here we send "ping" to the messages channel we made above,
	// from a new goroutine.
	go func() { messages <- "ping" }()

	//The <-channel syntax receives a value from the channel. Here we’ll receive the "ping" message we sent above and print it out.
	msg := <-messages
	fmt.Println(msg)
	//When we run the program the "ping" message is successfully passed from one goroutine to another via our channel.
	// By default sends and receives block until both the sender and receiver are ready. This property allowed us to wait at the end of our program for the "ping" message without having to use any other synchronization.

}
