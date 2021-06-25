package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "demo:toiyeuhanoi123-@tcp(127.0.0.1:3306)/demo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	categories := map[string]([]string){
		"điện thoại, máy tính bảng": []string{"smart phone", "điện thoại phổ thông", "máy tính bảng", "máy đọc sách"},
		"điện tử, điện lạnh": []string{"tivi", "dàn âm thanh", "tủ lạnh - tủ mát", "máy điều hoà", "phụ kiện điện lạnh", "máy giặt", "hút bụi", "lọc không khí", "rủa bát"},
		"máy tính, laptop": []string{"desktop", "server", "laptop", "phụ kiện máy tính"},
		"camera, camcoder": []string{"máy ảnh", "máy ảnh kỹ thuật số", "máy quay phim", "camera giám sá", "camera hành trình", "balo", "ống kính"},
		"đồ bếp":           []string{"lò vi sóng", "máy say sinh tố", "máy đánh trứng", "bếp từ", "bếp hồng ngoại", "bếp ga"},
	}

	statement, err := db.Prepare("INSERT INTO category (name, parent_id) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer statement.Close()

	for key, subcategories := range categories {
		result, err := statement.Exec(key, nil)
		if err != nil {
			panic(err.Error())
		}

		id, err := result.LastInsertId()

		for _, subcat := range subcategories {
			_, err = statement.Exec(subcat, id)
		}
		if err != nil {
			panic(err.Error())
		}
	}
}
