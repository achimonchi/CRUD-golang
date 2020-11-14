package main

import (
	"fmt"

	"github.com/achimonchi/CRUD-golang/mongodb/config"
)

func main() {

	_, err := config.GetMongoDB()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hello World")
	}
}
