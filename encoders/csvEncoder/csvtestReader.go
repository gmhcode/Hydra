package main

// import (
// 	"encoding/csv"
// 	"fmt"
// 	"log"
// )

// func main() {
// 	file, err := os.Open("cfile.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()

// 	r := csv.NewReader(file)
// 	r.Comment = '#'
// 	//seperates by ";" instead of comma
// 	r.Comma = ';'

// 	//if the file is small, use read all instead`
// 	// readAllInstead(r)

// 	for {
// 		record, err := r.Read()
// 		//EOF == end of file... if we reached the end of file, we can break the for loop
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			if parseError, ok := err.(*csv.ParseError); ok {
// 				fmt.Println("bad column: ", parseError.Column)
// 				fmt.Println("bad line: ", parseError.Line)
// 				fmt.Println("error reported: ", parseError.Err)
// 				if parseError.Err == csv.ErrFieldCount {
// 					continue
// 				}
// 			}
// 			log.Fatal(err)
// 		}
// 		fmt.Println("CSV Row: ", record)
// 		// 		prints
// 		// 		CSV Row:  [Jaro 5 ALA IOI]
// 		// 		CSV Row:  [Hala 4 ABD B0O]
// 		//		CSV Row:  [Kay 3 HBJ D3n]

// 		//converts string to int
// 		i, err := strconv.Atoi(record[1])
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(i * 4)
// 	}
// }

//if the file is small, use read all instead
// func readAllInstead(r *csv.Reader) {
// 	records, err := r.ReadAll()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("using read all", records)
// }
