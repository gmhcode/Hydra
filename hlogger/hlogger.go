package hlogger

import (
	"log"
	"os"
	"sync"
)

type hydraLogger struct {
	*log.Logger
	filename string
}

var hlogger *hydraLogger
var once sync.Once

//Gets the singleton instance of the hydraLogger
func GetInstance() *hydraLogger {
	once.Do(func() {
		hlogger = createLogger("hydralogger.log")
	})
	return hlogger
}

//Creates the hydraLogger singleton
func createLogger(fname string) *hydraLogger {
	file, _ := os.OpenFile(fname, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0777)

	return &hydraLogger{
		filename: fname,
		Logger:   log.New(file, "Hydra ", log.Lshortfile),
	}
}
