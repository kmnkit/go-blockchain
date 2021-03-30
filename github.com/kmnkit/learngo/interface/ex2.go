package ex2

import "fmt"

// Dog 구조체
type Dog struct {
	name   string
	weight int
}

// bite 메소드 구현
func (d Dog) bite() {
	fmt.Println(d.name, "이 물었다!")
}

// Behavior 동물 행동 인터페이스 구현
type Behavior interface {
	bite()
}

func ex2() {
	// 인터페이스 구현 예제
	// 예제1
	dog1 := Dog{"poll", 10}
	// dog1.bite()

	// var inter1 Behavior
	// inter1 = dog1
	// fmt.Println(inter1)
	// inter1.bite()

	dog2 := Dog{"marry", 12}
	inter2 := Behavior(dog2)
	inter2.bite()

	inters := []Behavior{dog1, dog2} // 인터페이스 자체도 슬라이스가 될 수 있다.

	for idx, _ := range inters {
		inters[idx].bite()
	}

	// 값 형태의 Interface 정의
	for _, val := range inters {
		val.bite()
	}
}
