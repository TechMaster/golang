package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initManufacturer() {
	sony := model.Manufacturer{Name: "Sony", CountryCode: "JP"}
	xiaomi := model.Manufacturer{Name: "Xiaomi", CountryCode: "CN"}
	leatherman := model.Manufacturer{Name: "Leatherman", CountryCode: "US"}
	mitsubishi := model.Manufacturer{Name: "Mitsubishi", CountryCode: "JP"}
	samsung := model.Manufacturer{Name: "Samsung", CountryCode: "KR"}
	lg := model.Manufacturer{Name: "LG", CountryCode: "KR"}
	delonghi := model.Manufacturer{Name: "Delonghi", CountryCode: "IT"}
	vinfast := model.Manufacturer{Name: "Vinfast", CountryCode: "VN"}

	Db.Create(&sony)
	Db.Create(&xiaomi)
	Db.Create(&leatherman)
	Db.Create(&mitsubishi)
	Db.Create(&samsung)
	Db.Create(&lg)
	Db.Create(&delonghi)
	Db.Create(&vinfast)
}

 
