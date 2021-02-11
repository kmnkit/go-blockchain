package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string // slice

}

func main() {
	favFood := []string{"Galbi", "Chicken"}
	marco := person{name: "Marco", age: 20, favFood: favFood}
	fmt.Println(marco.name)
}
