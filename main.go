package main

import (
	"fmt"

	"github.com/viitorags/gowork/config"
	"github.com/viitorags/gowork/router"
)

func main() {

	err := config.Init()
	if err != nil {
		fmt.Println(err)
	}

	router.Initialize()
}
