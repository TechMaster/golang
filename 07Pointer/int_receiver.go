package main

import "fmt"

type MyInt int

/* báo lỗi
func (p *int) Double() {

}*/

func (p *MyInt) Double() {
	*p = *p * 2
}

func DemoIntReceiver() {
	muoi := new(MyInt)
	*muoi = 10
	muoi.Double()
	fmt.Println(*muoi)
}
