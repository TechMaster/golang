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
func LoadConfig() (config Configuration, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName("dev")
	viper.SetConfigType("yaml")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
```

Trong hàm ```main``` của file [app.go](app.go) gọi phương thức `LoadConfig()` thôi
```go
func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		panic(err.Error())
	}
}
```

