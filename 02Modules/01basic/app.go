package main

import (
	"fmt"

	"github.com/TechMaster/mygomodule"
	"github.com/TechMaster/mygomodule/mathutil"
	"rsc.io/quote"
	"techmaster.vn/whatevername/utils"

	mygomoduleV2 "github.com/TechMaster/mygomodule/v2"
)

func main() {
	fmt.Println(quote.Go())
	var inputStr = "Hello"
	fmt.Println(mygomodule.Reverse(inputStr))
	utils.IsPrime(7)
	fmt.Println(utils.ProperCase("welcome to the techmaster!"))
	fmt.Println(mathutil.Add(10, 11))
	fmt.Println(mathutil.Minus(10, 11))
	fmt.Println(mygomoduleV2.Lower("Toi yeu Viet Nam"))
}
