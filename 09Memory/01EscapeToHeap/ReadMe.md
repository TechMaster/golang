# Kiểm tra xem biến nào sẽ escape to memory

Escape to memory là thuật ngữ chỉ một biến sẽ được cấp phát ở vùng nhớ heap thay vì vùng nhớ stack. Những trường hợp nào biến sẽ Escape to memory? Ví dụ này giúp bạn hiểu rõ hơn cơ chế quản lý bộ nhớ stack và heap trong Golang. Nó không giống với C/C++ mà có sự chọn lựa tối ưu khi tạo một biến ở heap.

>How do I know whether a variable is allocated on the heap or the stack?
From a correctness standpoint, you don't need to know. Each variable in Go exists as long as there are references to it. The storage location chosen by the implementation is irrelevant to the semantics of the language.

>The storage location does have an effect on writing efficient programs. When possible, the Go compilers will allocate variables that are local to a function in that function's stack frame. However, if the compiler cannot prove that the variable is not referenced after the function returns, then the compiler must allocate the variable on the garbage-collected heap to avoid dangling pointer errors. Also, if a local variable is very large, it might make more sense to store it on the heap rather than the stack.

>In the current compilers, if a variable has its address taken, that variable is a candidate for allocation on the heap. However, a basic escape analysis recognizes some cases when such variables will not live past the return from the function and can reside on the stack.

Sử dụng lệnh ```go build -gcflags -m``` để phân tích xem biến nào ở trường hợp nào sẽ escape to heap.


## Kiểm tra Escape to Heap qua các trường hợp cụ thể

Phát biểu của Golang team tương đối chung chung, do đó tôi tạo các ví dụ tình huống khác nhau để kiểm tra khi nào, một biến sẽ escape to heap (chuyển lên vùng nhớ heap) thay vì được lưu ở stack

1. Các tham số value hay pointer khi truyền vào dạng ```interface{}``` sẽ được escape to heap
2. Cấp phát slice kích thước đủ lớn.
  Ví dụ cấp phát 10,000 phần tử thì chuyển lên heap
   ```go
   make([]Person, 10000) escapes to heap
   ```
   nhưng nếu cấp phát 1,000 thì không cần chuyển lên heap
   ```go
   make([]Person, 1000) does not escape
   ```
3. Khác với C/C++ keyword ```new``` cấp phát vùng nhớ trên heap, nhưng ở Golang lại thường cấp phát luôn trong stack ```new(Person) does not escape```
  ```go
  alice := new(Person)
	alice.Name = "Alice"
	alice.Age = 21
  ```

4. Nếu dùng ```new``` cấp phát một biến trong hàm để trả về thì biến đó sẽ escape to heap
  ```go
  func makeAPerson(name string, age int) (person *Person) {
    person = new(Person) //Escape to heap
    person.Name = name
    person.Age = age
    return
  }
  ```

5. Trả về slice từ hàm, thì slice đó escape to heap
  ```go
  func ReturnASlice(size int) (result []string) {
    result = make([]string, size) //make([]string, size) escapes to heap
    return
  }
  ```

6. Trả về mảng từ hàm, thì array đó không escape to heap
  ```go
  func ReturnArrayFromFunc() (seasons [4]string) {
	  seasons = [4]string{"Spring", "Summer", "Fall", "Winter"}
	  return
  }
  ```
   
7. Trả về một struct từ hàm, thì struct đó không escape to heap
  ```go
  func ReturnAStructFromFunc() Person {
    person := Person{Name: "Cường", Age: 46}
    return person //does not escape to heap
  }
  ```
8. Trả về một pointer struct từ hàm, vùng nhớ struct đó được tạo ra ở heap
  ```go
  func ReturnPointerStructFromFunc() *Person {
    person := Person{Name: "Cường", Age: 46} //moved to heap: person
    return &person
  }
  ```
9. Khi truyền một biến vào tham số kiểu ```interface{}```
  Trong ví dụ này biến nào truyền vào tham số ```a interface{}``` sẽ escape to heap
  ```go
  func GetType(a interface{}) string {
    return reflect.TypeOf(a).String()
  }

  func InterfaceArgumentEscapeToHeap() {
    tom := Person{Name: "Tom", Age: 18}
    if GetType(tom) == "*main.Person" { //tom escapes to heap
      fmt.Println("tom is Person") //"tom is Person" escapes to heap
    }
  }
  ```


  