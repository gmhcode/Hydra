package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// func main() {

// 	//.Open is used for READ ONLY
// 	f1, err := os.Open("test1.txt")
// 	PrintFatalError(err)
// 	//Need to close file when you're done with it "cleans up resources"
// 	defer f1.Close()

// 	//Create a file for writing
// 	f2, err := os.Create("test2.txt")
// 	PrintFatalError(err)
// 	defer f2.Close()

// 	f3, err := os.OpenFile("test3.txt", os.O_APPEND|os.O_RDWR, 0666)
// 	//os.O_RDONLY //Read only
// 	//os.O_WRONLY //Write only
// 	//os.O_RDWR // Read and Write
// 	//os.O_APPEND //Append to end of file
// 	//os.O_CREATE // Create is non exist
// 	//os.O_TRUNC //Truncate file when opening
// 	//os.O_CREATE|os.0_RDWR |os.0_WRONLY

// 	//access control
// 	//0666 => Owner: (read & write),Group: other (read & write), and other (read & write)
// 	PrintFatalError(err)
// 	defer f3.Close()

// 	/*
// 	.########..########.##....##....###....##.....##.########....########.####.##.......########
// 	.##.....##.##.......###...##...##.##...###...###.##..........##........##..##.......##......
// 	.##.....##.##.......####..##..##...##..####.####.##..........##........##..##.......##......
// 	.########..######...##.##.##.##.....##.##.###.##.######......######....##..##.......######..
// 	.##...##...##.......##..####.#########.##.....##.##..........##........##..##.......##......
// 	.##....##..##.......##...###.##.....##.##.....##.##..........##........##..##.......##......
// 	.##.....##.########.##....##.##.....##.##.....##.########....##.......####.########.########
// 	*/
// 	err = os.Rename("test1.txt", "test1New.txt")
// 	PrintFatalError(err)

// 	/*
// 	.##.....##..#######..##.....##.########....########.####.##.......########
// 	.###...###.##.....##.##.....##.##..........##........##..##.......##......
// 	.####.####.##.....##.##.....##.##..........##........##..##.......##......
// 	.##.###.##.##.....##.##.....##.######......######....##..##.......######..
// 	.##.....##.##.....##..##...##..##..........##........##..##.......##......
// 	.##.....##.##.....##...##.##...##..........##........##..##.......##......
// 	.##.....##..#######.....###....########....##.......####.########.########
// 	*/
// 	err = os.Rename("./test1.txt", "./testfolder/test1.txt")
// 	PrintFatalError(err)
// 	/*
// 	..######...#######..########..##....##....########.####.##.......########
// 	.##....##.##.....##.##.....##..##..##.....##........##..##.......##......
// 	.##.......##.....##.##.....##...####......##........##..##.......##......
// 	.##.......##.....##.########.....##.......######....##..##.......######..
// 	.##.......##.....##.##...........##.......##........##..##.......##......
// 	.##....##.##.....##.##...........##.......##........##..##.......##......
// 	..######...#######..##...........##.......##.......####.########.########
// 	*/
// 	CopyFile("./test3.txt", "./testfolder/test3.txt")

// /*
// .########..########.##.......########.########.########.......###.......########.####.##.......########
// .##.....##.##.......##.......##..........##....##............##.##......##........##..##.......##......
// .##.....##.##.......##.......##..........##....##...........##...##.....##........##..##.......##......
// .##.....##.######...##.......######......##....######......##.....##....######....##..##.......######..
// .##.....##.##.......##.......##..........##....##..........#########....##........##..##.......##......
// .##.....##.##.......##.......##..........##....##..........##.....##....##........##..##.......##......
// .########..########.########.########....##....########....##.....##....##.......####.########.########
// */
// 	err = os.Remove("test2.txt")

// 	/*
// 	.########..########....###....########.....########.####.##.......########
// 	.##.....##.##.........##.##...##.....##....##........##..##.......##......
// 	.##.....##.##........##...##..##.....##....##........##..##.......##......
// 	.########..######...##.....##.##.....##....######....##..##.......######..
// 	.##...##...##.......#########.##.....##....##........##..##.......##......
// 	.##....##..##.......##.....##.##.....##....##........##..##.......##......
// 	.##.....##.########.##.....##.########.....##.......####.########.########
// 	*/
// 	bytes, err := ioutil.ReadFile("test3.txt")
// 	fmt.Println(string(bytes))

// 	// bufio scanner reads the file line by line.
// 	scanner := bufio.NewScanner(f3)
// 	count := 0
// 	for scanner.Scan() {
// 		count++
// 		fmt.Println("Found line:", count, scanner.Text())
// 	}

// 	//buffered write, efficient store in memorym saves disk I/O
// 	writebuffer := bufio.NewWriter(f3)
// 	for i := 1; i <= 5; i++ {
// 		writebuffer.WriteString(fmt.Sprintln("Added line", i))
// 	}
// 	writebuffer.Flush()

// 	GenerateFileStatusReport("test3.txt")

// 	filestat1, err := os.Stat("test3.text")
// 	PrintFatalError(err)

// }

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}

//CopyFile - Copy file fname1 to fname2
func CopyFile(fname1, fname2 string) {
	oldFile, err := os.Open(fname1)
	PrintFatalError(err)
	defer oldFile.Close()

	newFile, err := os.Create(fname2)
	PrintFatalError(err)
	defer newFile.Close()

	//copy bytes from source to destination
	_, err = io.Copy(newFile, oldFile)
	PrintFatalError(err)

	//flush file contents to disc
	err = newFile.Sync()
	PrintFatalError(err)

}

func GenerateFileStatusReport(fname string) {

	filestats, err := os.Stat(fname)
	PrintFatalError(err)

	fmt.Println("What's the file name?", filestats.Name())
	fmt.Println("Am I a directory?", filestats.IsDir())
	fmt.Println("What are the permissions?", filestats.Mode())
	fmt.Println("What's the file size?", filestats.Size())
	fmt.Println("What's the last time the file was modified?", filestats.ModTime())

}
