package main

import (
	"github.com/Hydra/hlogger"
	"github.com/Hydra/hydraweb/hydraportal"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra Web Service")

	hydraportal.Run()
}
