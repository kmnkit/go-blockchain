package ex4

import "fmt"

type Dog struct {
	name   string
	weight int
}

type Cat struct {
	name   string
	weight int
}

// 구조체 Dog 메소드 구현
func (d Dog) run() {
	fmt.Println(d.name, "is running!")
}

func (c Cat) run() {
	fmt.Println(c.name, "is running!")
}

func act(animal interface{ run() }) {
	animal.run()
}

// 덕타이핑은 메소드로만 타입을 판단한다.
func ex4() {
	// 익명예제 즉시 선언 후 사용 가능
	dog := Dog{"Poll", 10}
	act(dog)
}
