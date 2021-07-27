# Struct

Golang không có class mà chỉ có Struct

- Struct không có tính kế thừa.
- Composition over Inheritance
- Struct Golang có method nhưng chúng ta không viết method bên trong Struct như khai báo Class trong Java

Pointer receiver có thể thay đổi thuộc tính của struct khi kết thúc phương thức

Value receiver thay đổi thuộc tính bên trong phương thức thì được, thì kết thúc phương thức thì giá trị ban đầu của struct giữ nguyên.

Tại sao?
Pointer receiver làm việc trực tiếp trên đối tượng
Value receiver làm việc với bản copy của đối tượng

Khi dùng lệnh in địa chỉ sẽ thấy Pointer receiver dùng đối tượng trùng với đối tượng truyền vào, còn Value receiver dùng đối tượng khác địa chỉ
```go
fmt.Printf("pointer receiver %p\n", p)
```

Có thể khai báo Struct chưa bao nhiêu Struct khác tuỳ ý. Đây là Composition

```go
type Person struct {
	FirstName string //kiểu đứng sau tên biến
	LastName  string
	Age       int
	Address1  Address
	Address2  Address
}

type Address struct {
	Country string
	City    string
}
```

Khởi tạo 
```go
person := Person{"Trinh", "Cuong", 45,
		Address{"Vietnam", "Hanoi"},
		Address{"USA", "California"}}
```

Khả năng thứ 2:
```go
type Person struct {
	FirstName string //kiểu đứng sau tên biến
	LastName  string
	Age       int
	Address1  Address  //value
	Address2  *Address //pointer
}
```

```go
person := Person{"Trinh", "Cuong", 45,
		Address{"Vietnam", "Hanoi"},
		&Address{"USA", "California"}}  //Lấy địa chỉ của đối tượng Address{"USA", "California"}
```
