package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"Jaro", "5", "ALA,IOI"},
		{"Hala", "4", "ABD,B0O"},
		{"Kay", "3", "HBJ,D3N"},
	}

	file, err := os.Create("cfilew.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	w := csv.NewWriter(file)

	w.Comma = ';'
	//just writes it all to the new file
	w.WriteAll(records)

	////writes it all line by line
	// for _, record := range records {
	// 	if err := w.Write(record); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }
	////flush commits the written data to the disk
	// w.Flush()

	err = w.Error()
	if err != nil {
		log.Fatal(err)
	}
}
