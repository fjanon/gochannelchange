package main

import (
	"fmt"
	"time"
)

// this code allows to change the channel passed in to a go routine during execution
// that can be useful if we need to use a channel with a different size
// since a channel cannot be resized, we create and use another one
// the go routine using the channel updates the channel when the previous one is closed
// the go routine receives the channel through a handle - a pointer to the channel pointer

func createchannel(size int) chan int {
	var datachannel chan int = make(chan int, size)
	return datachannel
}

func receivedata(channelhandle *(*chan int)) {
	channel := **channelhandle
	for {
		data, more := <-channel
		if more {
			fmt.Printf("receivedata - channelhandle: %v channel: %v data: %v\n", channelhandle, channel, data)
		} else {
			// update the channel if the previous one was closed
			channel = **channelhandle
			fmt.Printf("receivedata - updating channel - channelhandle: %v channel: %v\n", channelhandle, channel)
		}
	}
}

func senddata(datachannel chan int, start int, max int) {
	fmt.Printf("\nsenddata - datachannel: %v type: %T\n", datachannel, datachannel)

	for i := start; i < max; i += 1 {
		fmt.Println("senddata - sending data i: ", i)
		datachannel <- i
	}
	fmt.Println("")
}

func main() {
	fmt.Println("Channel change in a go routine")

	// create and send data through a first channel
	channel1 := createchannel(10)
	fmt.Printf("main - channel1: %v type: %T\n", channel1, channel1)

	channelhandle := &channel1
	fmt.Printf("main - channelhandle: %v type: %T\n", channelhandle, channelhandle)

	// start the go routine with the pointer to the channel

	go receivedata(&channelhandle)

	start := 0
	max := 10
	senddata(channel1, start, max)

	time.Sleep(2 * time.Second)

	// create and send data through a second channel
	// the pointer to the channel has been changed
	channel2 := createchannel(5)

	// update the handle so the go routine picks up the change
	channelhandle = &channel2

	fmt.Printf("\nmain - channelhandle: %v type: %T\n", channelhandle, channelhandle)
	fmt.Printf("main - *channelhandle: %v type: %T\n", *channelhandle, *channelhandle)

	// closing first channel to trigger the channel change in the receiving goroutine
	close(channel1)

	// send data through the new channel
	start = 20
	max = 25
	senddata(channel2, start, max)

	// use a channel or a waitgroup to properly wait for the go routine completion
	time.Sleep(10 * time.Second)
}

/*
Channel change in a go routine
main - channel1: 0xc00006e000 type: chan int
main - channelhandle: 0xc00000e030 type: *chan int

senddata - datachannel: 0xc00006e000 type: chan int
senddata - sending data i:  0
senddata - sending data i:  1
senddata - sending data i:  2
senddata - sending data i:  3
senddata - sending data i:  4
senddata - sending data i:  5
senddata - sending data i:  6
senddata - sending data i:  7
senddata - sending data i:  8
senddata - sending data i:  9

receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 0
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 1
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 2
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 3
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 4
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 5
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 6
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 7
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 8
receivedata - channelhandle: 0xc00000e038 channel: 0xc00006e000 data: 9

main - channelhandle: 0xc00000e048 type: *chan int
main - *channelhandle: 0xc000076000 type: chan int

senddata - datachannel: 0xc000076000 type: chan int
senddata - sending data i:  20
senddata - sending data i:  21
senddata - sending data i:  22
senddata - sending data i:  23
senddata - sending data i:  24

receivedata - updating channel - channelhandle: 0xc00000e038 channel: 0xc000076000
receivedata - channelhandle: 0xc00000e038 channel: 0xc000076000 data: 20
receivedata - channelhandle: 0xc00000e038 channel: 0xc000076000 data: 21
receivedata - channelhandle: 0xc00000e038 channel: 0xc000076000 data: 22
receivedata - channelhandle: 0xc00000e038 channel: 0xc000076000 data: 23
receivedata - channelhandle: 0xc00000e038 channel: 0xc000076000 data: 24

*/
