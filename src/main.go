package main

import (
	"fmt"
	"github.com/callingsid/shopping_product/src/app"
)


var (
	appName = "product"
)

func main() {
	fmt.Printf("Starting %v\n", appName)
	app.StartApp()
	//select {} // block forever
	//done := make(chan bool)

	//<-done
	//<-make(chan int)
}





