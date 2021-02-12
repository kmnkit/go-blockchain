package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string) // 어느 형태의 데이터를 받을지 구체적으로 짖어해 줘야 함.

	people := [5]string{"nico", "flynn", "dal", "japanGuy", "larry"}
	for _, person := range people {
		go isSexy(person, c)
	}
	for i := 0; i < len(people); i++ {
		fmt.Println(<-c)
	}
	// resultOne := <-c
	// resultTwo := <-c
	// resultThree := <-c
	// resultFour := <-c
	// resultFive := <-c
	// fmt.Println("Received this Message:", resultOne)
	// fmt.Println("Received this Message:", resultTwo)
	// fmt.Println("Received this Message:", resultThree)
	// fmt.Println("Received this Message:", resultFour)
	// fmt.Println("Received this Message:", resultFive)
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 1)
	c <- person + " is Sexy"
}
