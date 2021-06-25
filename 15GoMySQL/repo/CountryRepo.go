package repo

import (
	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initCountry() {
	countries := map[string]string{
		"VN": "Viet nam",
		"CN": "China",
		"US": "USA",
		"DE": "Germanny",
		"SG": "Singapore",
		"ES": "Spain",
		"KR": "South Korea",
		"JP": "Japan",
		"IT": "Italy",
		"RU": "Russia",
		"ML": "Malaysia",
	}

	for country_code, country_name := range countries {
		country := model.Country{Code: country_code, Name: country_name}
		result := Db.Create(&country)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
