package main

import (
	"fmt"

	"github.com/sudhanshuraheja/golang-sample/pkg/config"
)

func main() {
	config.Init()
	fmt.Println("Sample CLI")
}
