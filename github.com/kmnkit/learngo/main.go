package main

import "fmt"

func main() {
	a := 2  // a의 값이 2
	b := &a // a의 메모리 주소

	fmt.Println(&a, b) // a의 메모리 주소와 b의 값(a의 메모리 주소)를 출력
	fmt.Println(*b)    // b에 있는 실제 값을 출력

	fmt.Println(&a) // a의 메모리주소
	fmt.Println(a)  // a에 있는 값
	fmt.Println(&b) // b의 메모리주소
	fmt.Println(b)  // b의 값
	*b = 20
	fmt.Println(a) // a의 값이 20으로 변경되었음
	fmt.Println(b) // b의 값(a의 주소값)은 그대로
}
