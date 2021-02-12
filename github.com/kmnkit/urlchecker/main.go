package main

import (
	"fmt"
	"time"
)

func main() {
	// 메인함수의 역할이 종료될 때 그냥 다 끝나버림
	go sexyCount("nico")
	sexyCount("flynn")
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}
