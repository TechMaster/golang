package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Product struct {
	Name     string
	Category string
	Price    int
}

func randomInt(min int, max int) int {
	return rand.Intn(max-min+1) + min
}
func main() {
	category := [4]string{"fashion", "electronics", "sport", "food"}
	products := [20]Product{}

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(products); i++ {
		products[i] = Product{
			fmt.Sprintf("%s %d", "Product", i),
			category[rand.Intn(len(category))],
			randomInt(100, 200),
		}
	}

	for _, product := range products {
		fmt.Println(product)
	}
	fmt.Println("------Sắp xếp theo giá cao nhất--------")
	sortedProducts := products[:]
	sort.Slice(sortedProducts, func(i, j int) bool {
		return sortedProducts[i].Price > sortedProducts[j].Price
	})
	for _, product := range sortedProducts {
		fmt.Println(product)
	}
	fmt.Println("------Top 5 giá cao nhất--------")
	for i := 0; i < 5; i++ {
		fmt.Println(sortedProducts[i])
	}
	fmt.Println("--Đếm sản phẩm theo category----")
	categoryCount := map[string]int{}
	for _, product := range products {
		if _, ok := categoryCount[product.Category]; ok {
			categoryCount[product.Category]++
		} else {
			categoryCount[product.Category] = 0
		}
	}

	for key, value := range categoryCount {
		fmt.Println(key, value)
	}
}
