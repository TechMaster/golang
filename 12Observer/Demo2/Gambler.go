package main

import "fmt"

type Gambler struct {
	name string
}

func (g *Gambler) Update(value int) {
	if value > 90 {
		fmt.Println("Thắng cá độ to rồi")
	} else {
		fmt.Println("Thua cá độ rồi")
	}
}
