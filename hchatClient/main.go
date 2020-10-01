package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	name := fmt.Sprintf("Anonymouse%d", rand.Intn(400))

	fmt.Println("Starting hydraChatCLient....")
	fmt.Println("what's your name?")
	fmt.Scanln(&name)

	fmt.Printf("Hello %s, connecting to the hydra chat system... \n", name)
	conn, err := net.Dial("tcp", "127.0.0.1:2300")
	if err != nil {
		log.Fatal("Could not connect to hydra chat system", err)
	}
	fmt.Println("Connected to hydra chat system")
	name += ":"

	defer conn.Close()

	go func() {
		scanner := bufio.NewScanner(conn)
		//Keeps litening and printing text for incomming messages till scanner.scan returns false
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
	}()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() && err == nil {
		msg := scanner.Text()
		_, err = fmt.Fprintf(conn, name+msg+"\n")
	}

}
