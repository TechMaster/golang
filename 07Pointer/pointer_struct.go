package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

//Pointer receiver có thể
func (p *Person) IncreaseAge() {
	p.Age += 1
}

//Value receiver không thay đổi giá trị của tham số truyền vào
func (p Person) NotReallyIncreaseAge() {
	p.Age += 1 //warning: ineffective assignment to field Person.Age
}

func EncryptPersonNamePointerArg(p *Person) {
	var result string
	for _, v := range p.Name {
		result = string(v) + result
	}
	p.Name = result
}

func EncryptPersonNameValueArg(p Person) {
	var result string
	for _, v := range p.Name {
		result = string(v) + result
	}
	p.Name = result
}

func DemoPointerStruct() {
	bob := &Person{"Bob", 46} //bob là con trỏ chỉ đến địa chỉ của đối tượng Person{"Cường", 46}
	bob.NotReallyIncreaseAge()
	fmt.Println(bob)

	tom := new(Person) //new(Person) trả về
	tom.Name = "Tom"
	tom.Age = 27
	tom.IncreaseAge()
	tom.NotReallyIncreaseAge()
	fmt.Println(tom)

	alice := Person{Name: "Alice", Age: 18} //Khởi tạo struct với tên thuộc tính sẽ tường minh hơn
	alice.IncreaseAge()
	fmt.Println(&alice)

	EncryptPersonNamePointerArg(tom) //thay đổi giá trị. Tom --> moT
	fmt.Println(tom)

	// EncryptPersonNamePointerArg(alice) không chấp nhận value truyền vào

	EncryptPersonNameValueArg(alice) //không thay đổi được giá trị
	fmt.Println(&alice)
	// EncryptPersonNameValueArg(tom) không chấp nhận pointer truyền vào
}
