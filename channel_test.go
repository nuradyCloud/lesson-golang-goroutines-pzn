package lesson_golang_goroutines_pzn

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	//channel <- "Ady"
	//fmt.Println(<-channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Nur Ady Pamungkas"
		fmt.Println("Done send data to channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)

}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Nur Ady Pamungkas"
}

func TestChannelAsParam(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}
