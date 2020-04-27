package main

import (
	"fmt"
	"sync"
	"time"
)

func useUnbufferedChannels() {
	fmt.Println("starting useUnbufferedChannels()")
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(chArg chan int, wgArg *sync.WaitGroup) {
		// receiving message from channel, arrow coming out of channel
		fmt.Println("receiving message:", <-chArg) 
		wgArg.Done()
	} (ch, wg)

	go func(chArg chan int, wgArg *sync.WaitGroup) {
		fmt.Println("Sending 42")
		// sending message to channel, arrow going into channel
		chArg <- 42 
		wgArg.Done()
	} (ch, wg)

	wg.Wait()
}

func useBufferedChannels() {
	fmt.Println("starting useBufferedChannels()")
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2)

	wg.Add(2)
	go func(chArg chan int, wgArg *sync.WaitGroup) {
		// receiving message from channel, arrow coming out of channel
		fmt.Println("receiving message in first go routine:", <-chArg) 
		chArg <- 27
		wgArg.Done()
	} (ch, wg)

	go func(chArg chan int, wgArg *sync.WaitGroup) {
		fmt.Println("Sending 42")
		// sending message to channel, arrow going into channel
		chArg <- 42 
		time.Sleep(5 * time.Millisecond)
		fmt.Println("receiving message in second go routine:", <-chArg) 
		wgArg.Done()
	} (ch, wg)

	wg.Wait()
}

func useDirectionalChannels() {
	fmt.Println("starting useBufferedChannels()")
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2)

	wg.Add(2)
	go func(chArg <-chan int, wgArg *sync.WaitGroup) {
		// receiving message from channel, arrow coming out of channel
		fmt.Println("receiving message in first go routine:", <-chArg) 
		
		wgArg.Done()
	} (ch, wg)

	go func(chArg chan<- int, wgArg *sync.WaitGroup) {
		fmt.Println("Sending 42")
		// sending message to channel, arrow going into channel
		chArg <- 42 
		time.Sleep(5 * time.Millisecond)
		
		wgArg.Done()
	} (ch, wg)

	wg.Wait()
}

func closeChannels() {
	fmt.Println("starting closeChannels")
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2)

	wg.Add(2)
	go func(chArg chan int, wgArg *sync.WaitGroup) {
		// receiving message from channel, arrow coming out of channel
		fmt.Println("receiving message in first go routine:", <-chArg) 
		close(chArg)
		fmt.Println( "trying to receive a message in first go routine w/closed channel:", <-chArg) 
		wgArg.Done()
	} (ch, wg)

	go func(chArg chan int, wgArg *sync.WaitGroup) {
		fmt.Println("Sending 42")
		// sending message to channel, arrow going into channel
		chArg <- 42 
		// close(chArg) // no effect if by itself
		// chArg <- 27 // this will cause a panic
		// time.Sleep(5 * time.Millisecond)
		
		wgArg.Done()
	} (ch, wg)

	wg.Wait()
}

func useIfWithChannels() {
	fmt.Println("starting useIfWithChannels()")
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2)

	wg.Add(2)
	go func(chArg <-chan int, wgArg *sync.WaitGroup) {
		// receiving message from channel, arrow coming out of channel
		if msg, ok := <-chArg; ok {
			fmt.Println("receiving message in first go routine:", msg, ", is it okay?:", ok) 
		} else {
			fmt.Println("Message is not ok:", ok, ", so channel was closed")
		}
		
		// if ok is false, channel was closed
		wgArg.Done()
	} (ch, wg)

	go func(chArg chan<- int, wgArg *sync.WaitGroup) {
		// fmt.Println("Sending 42")
		// sending message to channel, arrow going into channel
		// chArg <- 0
		close(chArg)
		time.Sleep(5 * time.Millisecond)
		
		wgArg.Done()
	} (ch, wg)

	wg.Wait()
}

func useForWithChannels() {
	fmt.Println("starting useForWithChannels()")
	wg := &sync.WaitGroup{}
	ch := make(chan int, 2)

	wg.Add(2)
	go func(chArg <-chan int, wgArg *sync.WaitGroup) {
		// receiving message from channel, arrow coming out of channel
		fmt.Println("receiving message in first go routine:", <-chArg) 
		// for i := 0; i < 10; i++ {
		// issue with for loop: We may not know how many messages we will get
		// so do this
		for msg := range chArg {
			// msg will be populated by <-chArg) 
			fmt.Println("receiving message in first go routine:", msg)
			// fmt.Println( <-chArg) 
		}
		wgArg.Done()
	} (ch, wg)

	go func(chArg chan<- int, wgArg *sync.WaitGroup) {
		
		// sending message to channel, arrow going into channel
		for i := 0; i < 10; i++ {
			fmt.Println("Sending in second go routine ", i)
			chArg <- i
		}
		close(ch) // must close channel so receiving channel will know there are no more messages
		time.Sleep(5 * time.Millisecond)
		
		wgArg.Done()
	} (ch, wg)

	wg.Wait()
	
}

func main() {
	// useUnbufferedChannels()
	// useBufferedChannels()
	// useDirectionalChannels()
	// closeChannels()
	// useIfWithChannels()
	useForWithChannels()
	
}


