# Cấu trúc ứng dụng REST API, Viper, GORM

## 1. Giới thiệu
Dự án này nâng cấp những điểm sau đây từ buổi trước
1. Cấu trúc thư mục có thêm những sub package như config, controller, model, repo
2. Bổ xung cơ chế đọc cấu hình sử dụng package [github.com/spf13/viper](https://github.com/spf13/viper)
3. Sử dụng [GORM](https://gorm.io/index.html) để thao tác dữ liệu với MySQL


## 2. Cấu trúc lại thư mục

```
.
├── config  <-- Đọc cấu hình sử dụng Viper
│   └── config.go
├── controller <-- Các controller xử lý request đến
│   ├── ManufactureController.go
│   └── ProductController.go
├── model <-- Định nghĩa các model. Các bạn cần bổ xung thêm
│   ├── Category.go
│   ├── Country.go
│   ├── Manufacturer.go
│   └── Product.go
├── repo <-- Lưu các phương thức xử lý dữ liệu chuyên cho từng Model
│   ├── CategoryRepo.go
│   ├── CountryRepo.go
│   ├── ManufactureRepo.go
│   ├── ProductRepo.go
│   └── Repo.go  <-- Lưu biến toàn cục Database Connection `var Db *gorm.DB`
├── routes <-- Cấu hình định tuyến các request đến ứng với phương thức của Controller`
│   └── ConfigRouter.go
├── sql <-- Lưu các file SQL script tạo bảng và xoá bảng
│   ├── DropTable.sql
│   └── OnlineShop.sql
├── app.go <-- File chạy chính
├── dev.yml <-- File cấu hình ở môi trường development
├── go.mod
├── go.sum
├── GORM.md
└── ReadMe.md <-- File này quan trọng nhất, làm gì thì làm, luôn phải viết document!
```

## 3. Đọc cấu hình sử dụng Viper

Tạo file cấu hình định dạng YAML [dev.yml](dev.yml)
```yaml
db:
  host: localhost
  port: 3306
  user: demo
  pass: toiyeuhanoi123-
  database: demo
```

Hãy xem file đọc cấu hình [config/config.go](config/config.go)
Khai báo cấu trúc struct trong Golang tương ứng với cấu trúc của [dev.yml](dev.yml)
```go
var Config Configuration //Khai báo public và biến toàn cục để các module khác dùng nhé

type Configuration struct {
	Db DBConf //tương đương với thẻ db:
}

type DBConf struct {  // tương đương với các thẻ con bên dưới thẻ db:
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}
```

Hàm đọc thông tin từ file yaml rồi đổ vào struct Configuration
```go
func LoadConfig() (err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&Config)
	return
}
```

Trong hàm ```main``` của file [app.go](app.go) gọi phương thức `LoadConfig()` thôi
```go
func main() {
	err := config.LoadConfig() //Đọc cấu hình từ file dev.yml đổ vào biến toàn cục config.Config
	if err != nil {
		panic(err.Error())
	}
	repo.Connect(config.Config) //Truyền cấu hình vào phương thức Connect của package repo
}
```

## 4. Sử dụng GORM để thao tác dữ liệu

Lập trình bằng GORM có gì hay so với lập trình MySQL thuần?

* Code viết ngắn và gọn hơn cách cũ. Truy vấn MySQL thuần xong phải gán từng trường vào biến, rồi từ biến tạo thành struct...
* Lập trình hướng đối tượng thay vì hướng câu lệnh SQL
* Có thể định nghĩa cấu trúc bảng qua Golang Struct

Trong dự án này, tôi vẫn sử dụng https://dbdiagram.io/ để sinh ra DDL script (DDL = Data Definition Language). Tuy nhiên tôi tuân thủ quy ước của GORM:

1. Tên bảng là danh từ số nhiều, các các từ bằng dấu `_` (snake_case) ví dụ: `product_prices`, `relate_products`
2. Khoá chính, primary, mặc định có tên là `ID` kiểu unsigned integer. Có trường hợp đặc biệt vẫn customize được.
3. Tên các cột chữ thường, sử dụng snake_case

Code trong dbdiagram.io
```
Table products {
  ID int [pk, increment]
  name varchar [not null]
  description varchar [not null]
  madein varchar(2) [ref: > countries.code]
  price int
  manufacturer_id int [ref: > manufacturers.ID]
  category_id int [ref: > categories.ID]
}
```
sẽ sinh ra DDL script như dưới

```sql
CREATE TABLE `products` (
  `ID` int PRIMARY KEY AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `madein` varchar(2),
  `price` int,
  `manufacturer_id` int,
  `category_id` int
);
```


