package main

import (
	"github.com/sudhanshuraheja/golang-sample/pkg/config"
	"github.com/sudhanshuraheja/golang-sample/pkg/logger"
)

func main() {
	config.Init()
	logger.Init()

	logger.Infoln("Sample CLI")
}
