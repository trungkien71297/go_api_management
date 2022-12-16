package main

import (
	"fmt"
	"go_api_management/app"
	"time"
)

func main() {
	fmt.Println(time.Now().Unix())
	app.StartApplication()
}
