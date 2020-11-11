package hydrachat

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Hydra/hlogger"
)

var logger = hlogger.GetInstance()

func Run() error {
	l, err := net.Listen("tcp", ":2100")
	r := CreateRoom("HydraChat")
	if err != nil {
		logger.Println("Error connecting to chat client", err)
		return err
	}

	go func() {
		// Handle SIGINT and SIGTERM.
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
		<-ch

		l.Close()
		fmt.Println("Closing tcp connection")
		close(r.Quit)
		if r.ClCount() > 0 {
			<-r.Msgch
		}
		os.Exit(0)
	}()

	for {
		conn, err := l.Accept()
		if err != nil {
			logger.Println("Error accepting connection from chat client", err)
			break
		}
		go handleConnection(r, conn)
	}

	return err
}

// Run - Starts the hydrachat server
// func Run(connection string) error {
// 	listener, err := net.Listen("tcp", connection)

// 	if err != nil {
// 		logger.Println("error connecting to chat client", err)
// 		return err
// 	}

// 	room := CreateRoom("HydraChat")

// 	//Monitors the OS to see if it should close the app
// 	go func() {
// 		//Monitoring this channel to see if a program is closing
// 		channel := make(chan os.Signal)
// 		//if the operating system throws SIGINT or SIGTERM, it will notify the channel
// 		signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
// 		//this will block the other code in this goRoutine, until the channel receives a terminate signal
// 		<-channel

// 		listener.Close()
// 		fmt.Println("Closing tcp connection")
// 		//Sends a notification to all functions listening to the room.Quit channel
// 		close(room.Quit)
// 		if room.ClCount() > 0 {
// 			<-room.Msgch
// 		}
// 		os.Exit(0)
// 	}()

// 	for {
// 		//Blocks the rest of the loop
// 		//listens for when someone tries to connect to our server
// 		conn, err := listener.Accept()

// 		if err != nil {
// 			logger.Println("Error accepting connection form chat client", err)
// 			break
// 		}
// 		go handleConnection(room, conn)
// 	}

// 	return err
// }

func handleConnection(room *room, clientConnection net.Conn) {
	logger.Println("Received request from client", clientConnection.RemoteAddr())
	room.AddClient(clientConnection)
}
