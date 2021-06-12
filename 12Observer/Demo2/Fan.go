package main

import "fmt"

// Fan sẽ là đối tượng quan sát / lắng nghe thay đổi từ simpleSubject
type Fan struct {
	name string
}

// Hàm tuân thủ từ interface Observer
func (u *Fan) Update(point int) {
	fmt.Printf("%s nhận được cập nhật : %d\n", u.name, point)
}
