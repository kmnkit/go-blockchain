package ex5

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 인터페이스는 모든 형을 다 받을 수 있다.
func printValue(i interface{}) {
	fmt.Println("ex1 :", i)
}

func ex5() {
	// 인터페이스 활용(빈 인터페이스)
	// 함수 내에서 어떠한 타입이라도 유연하게 매개변수로 받을 수 있다.(만능) -> 모든 타입 지정 가능하다.
	dog := Dog{"Poll", 10}
	cat := Cat{"Bob", 5}
	printValue(dog)
	printValue(cat)
	printValue(10)
	printValue(true)
	printValue("Animal")
	printValue("25.5")
	printValue([]Dog{})
	printValue([5]Dog{})
}
