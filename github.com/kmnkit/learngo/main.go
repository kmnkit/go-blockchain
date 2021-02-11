package main

import "fmt"

func main() {
	// names := [5]string{"nico", "lynn", "dal"}
	// names[3] = "alalala"
	// names[4] = "ahahaha"
	// index를 넘어가므로 아래 코드는 에러가 남.
	// names[5] = "jajaja"
	names := []string{"nico", "lynn", "dal"}
	names = append(names, "flynn")
	fmt.Println(names)
}
