package main

import (
	"fmt"
	"time"

	"github.com/trungkien71297/go_api_management/app"
)

func main() {
	fmt.Println(time.Now().Unix())
	app.StartApplication()
}
