package main

import (
	"fmt"
	"os"
	"time"
)

// func main() {
// WatchFile("testfile.txt")

// }

func WatchFile(fileName string) {
	filestat1, err := os.Stat(fileName)
	PrintFatalError(err)
	for {
		time.Sleep(1 * time.Second)
		filestat2, err := os.Stat(fileName)
		PrintFatalError(err)

		//we just check the modified time against the prevous modified time.
		if filestat1.ModTime() != filestat2.ModTime() {

			fmt.Println("File was modified at", filestat2.ModTime())
			filestat1, err = os.Stat(fileName)
			PrintFatalError(err)

		}
	}
}
