package main

import (
	"MasteringGoTutorial/HYDRA/hlogger"
	"MasteringGoTutorial/HYDRA/hydraportal"
)

func main() {
	logger := hlogger.GetInstance()
	logger.Println("Starting Hydra web service")

	hydraportal.Run()
}