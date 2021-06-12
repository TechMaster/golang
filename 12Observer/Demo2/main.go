package main

import "fmt"

func main() {
	FootballClub := NewFootballClub("Manchester United")
	tom := Fan{name: "Tom"}
	bob := Fan{name: "Bob"}
	alice := Fan{name: "Alice"}
	mike := Fan{name: "Mike"}

	trumCado := Gambler{name: "Trùm cá độ"}

	FootballClub.RegisterObserver(&tom)
	FootballClub.RegisterObserver(&bob)
	FootballClub.RegisterObserver(&alice)
	FootballClub.RegisterObserver(&mike)
	FootballClub.RegisterObserver(&trumCado)

	// Mọi Fan đều nhận được thông báo khi FootballClub thay đổi giá trị
	fmt.Println("------First update")
	FootballClub.SetPoint(89)

	//Loại bỏ mike ra khỏi danh sách nhận thông báo
	FootballClub.RemoveObserver(&mike)
	fmt.Println("------Second update")
	//Mọi người ngoại trừ mike đều nhận được thông báo
	FootballClub.SetPoint(92)
}
