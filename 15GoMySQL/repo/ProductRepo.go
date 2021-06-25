package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initProduct() {
	var sony, xiaomi model.Manufacturer
	Db.Where("name = ?", "Sony").First(&sony)
	Db.Where("name = ?", "Xiaomi").First(&xiaomi)

	var taiNge_Cat, locKhongKhi_Cat model.Category
	Db.Where("name LIKE ?", "tai nghe%").First(&taiNge_Cat)
	Db.Where("name LIKE ?", "lọc không khí%").First(&locKhongKhi_Cat)

	sonyWH_1000XM4 := model.Product{
		Name:         "Sony WH-1000XM4",
		Description:  "Tai nghe chống ổn chủ động thế hệ 3 của Sony, hiệu suất giảm tạp âm lên đến 95%",
		Price:        5500000,
		Madein:       "ML",
		Manufacturer: &sony,
		Category:     &taiNge_Cat,
	}
	Db.Create(&sonyWH_1000XM4)

	xiaomiAirPurifer3C := model.Product{
		Name:         "Xiaomi Air Purifier 3C",
		Description:  "Máy lọc không khí cho phòng 28-40m², sản phẩm có khả năng loại bỏ bụi mịn PM2.5, lọc khuẩn, lọc mùi với tiêu chuẩn quốc tế, khả năng quản lý bằng ứng dụng Mi Home có tiếng việt...",
		Price:        2299000,
		Madein:       "CN",
		Manufacturer: &xiaomi,
		Category:     &locKhongKhi_Cat,
	}
	Db.Create(&xiaomiAirPurifer3C)
}

/*
type Product struct {
	ID             uint
	Name           string
	Description    string
	Price          uint
	Madein         string
	Country        *Country `gorm:"foreignKey:Madein"`
	ManufacturerID uint
	Manufacturer   *Manufacturer
	CategoryID     uint
	Category       *Category
}
*/
