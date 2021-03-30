package ex3

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
func (d Dog) bite() {
	fmt.Println(d.name, "Dog bites!")
}
func (d Dog) sound() {
	fmt.Println(d.name, "Dog barks!")
}
func (d Dog) run() {
	fmt.Println(d.name, "Dog is running!")
}

func (c Cat) bite() {
	fmt.Println(c.name, "Cat 할퀴다!")
}
func (c Cat) sound() {
	fmt.Println(c.name, "Cat cries!")
}
func (c Cat) run() {
	fmt.Println(c.name, "Cat is running!")
}

// 동물의 행동특성 인터페이스 선언
type Behavior interface {
	bite()
	sound()
	run()
}

// interface를 param으로 받는다.
func act(animal Behavior) {
	animal.bite()
	animal.sound()
	animal.run()
}

// 덕타이핑은 메소드로만 타입을 판단한다.
func ex3() {
	dog := Dog{"Poll", 10}
	cat := Cat{"Bob", 5}

	act(dog)
	act(cat)
}
