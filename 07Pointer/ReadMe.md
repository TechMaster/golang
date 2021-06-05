> Bài này tôi sẽ trình bày mấy vấn đề căn bản trong Golang qua code mẫu:
> Pointer trong Golang
> Receiver function
> [Code mẫu ở đây](https://github.com/TechMaster/golang/tree/main/07Pointer)

# 1. Pointer trong Golang

Java, C#, JavaScript, Python không có con trỏ.
C, C++ có khái niệm con trỏ và rất nhiều kỹ thuật quản lý bộ nhớ với con trở.
Golang cũng có con trỏ, nhưng bộ nhớ trong Golang được thu hồi tự động, lập trình viên không phải bận tâm.
Con trỏ trong Golang dùng để chỉ đến một biến (vùng dữ liệu), thay vì copy biến đó. Nó hữu ích khi chúng ta cần sửa đổi biến, hoặc khi việc copy biến có kích thước lớn sẽ tốn kém bộ nhớ.

- ```*``` dùng để khai báo kiểu dữ liệu con trỏ. 
- ```&``` dùng để lấy địa chỉ của một biến
- ```*``` liên quan gì với ```&```? trả lời giá trị của con trỏ là địa chỉ của biến mà con trỏ đó chỉ tới

## 1.1 Tạo biến con trỏ bằng từ khoá ```new```
```go
age := new(int) //age là con trỏ đến biến kiểu int
*age = 26
fmt.Println(*age) //26: giá trị của biến mà con trỏ age chỉ đến
fmt.Println(age) //0xc00012a008: địa chỉ của biến mà con trỏ age chỉ đến
fmt.Println(&age) //0xc00012e018: địa chỉ của chính con trỏ age
```

## 1.2 Lấy địa chỉ biến bằng ```&```
```go
tuoi := 27 //khởi tạo một biến dạng value
age = &tuoi //con trỏ age chỉ đến biến tuổi
fmt.Println(*age)  //27
*age += 1 //tăng value mà age chỉ đến thêm 1
fmt.Println(tuoi) //28
```


## 1.3 Khác với C, C++, Golang không hỗ trợ phép tính với con trỏ
Tuy nhiên nếu bạn chấp nhận rủi ro có thể dùng gói thư viện unsafe
[Tham khảo thêm](https://stackoverflow.com/questions/32700999/pointer-arithmetic-in-go)

# 2. Receiver
Golang không có class mà chỉ có struct. Một đối tượng struct cũng có các method. Khi ta gọi đến các method này. Đối tượng gọi sẽ là caller. Còn đối tượng có method thực thi sẽ là receiver (nhận).
Receiver trong Golang chỉ nhận những kiểu do lập trình viên định nghĩa.

## 2.1 Pointer receiver vs Value receiver

Khác biệt giữa Pointer receiver vs Value receiver đã nói trong slide Golang Basic.
- Pointer receiver dùng để thay đổi giá trị của đối tượng. Nó không thread safe.
- Value receiver không thay đổi giá trị của đối tượng. Nó thread safe.

```go
type Person struct {
	Name string
	Age  int
}

func (p *Person) String() string {
	return fmt.Sprintf("%s : %d", p.Name, p.Age)
}

//Pointer receiver có thể thay đổi giá trị của đối tượng
func (p *Person) IncreaseAge() {
	p.Age += 1
}

//Value receiver không thay đổi giá trị của của đối tượng
func (p Person) NotReallyIncreaseAge() {
	p.Age += 1 //warning: ineffective assignment to field Person.Age
}
```

## 2.2 Pointer receiver có thể nhận value, Value receiver có thể nhận pointer
```go
bob := &Person{"Bob", 46} //bob là pointer
bob.NotReallyIncreaseAge()
fmt.Println(bob)

tom := new(Person) //tom là pointer
tom.Name = "Tom"
tom.Age = 27
tom.IncreaseAge()
tom.NotReallyIncreaseAge()  //Value receiver chấp nhận cả pointer
fmt.Println(tom)

alice := Person{Name: "Alice", Age: 18} //alice là value
alice.IncreaseAge() //Pointer receiver chấp nhận value
fmt.Println(&alice)
```

## 2.3 Receiver không nhận kiểu cơ bản, nhưng có cách xử lý (workaround)
Chúng ta chỉ có thể định nghĩa hàm Receiver cho kiểu struct, composite...
với những kiểu đơn giản như ```int, float32``` bạn sẽ nhận được báo lỗi ```cannot define new methods on non-local type int```

```go
func (p *int) Double() {

}
```

Phải sửa lại như sau
```go
type MyInt int //khai báo một kiểu mới thực chất vẫn là int

func (p *MyInt) Double() {
	*p = *p * 2
}

func main() {
  muoi := new(MyInt)
  *muoi = 10
  muoi.Double()
  fmt.Println(*muoi) //in ra 20
}
```

## 2.4 Receiver cho slice []int
Định nghĩa như sau sẽ báo lỗi ```invalid receiver type []int ([]int is not a defined type)```

```go
func (p []int) DoubleIt() {
	for i, v := range p {
		p[i] = v * 2
	}
}
```

Phải định nghĩa một kiểu có tên ```type SliceInt []int```
```go
type SliceInt []int
func (p SliceInt) DoubleIt() { //giờ không báo lỗi nữa
	for i, v := range p {
		p[i] = v * 2
	}
}
var sliceInt SliceInt = []int{1, 2, 3, 4} //khai báo biến cũng phải đúng kiểu sliceInt
//sliceInt := []int{1, 2, 3, 4} cách này không biên dịch được
sliceInt.DoubleIt()
fmt.Println(sliceInt) //[2 4 6 8]
```