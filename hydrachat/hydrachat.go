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

func Run(connection string) error {
	listener, err := net.Listen("tcp", connection)

	if err != nil {
		logger.Println("error connecting to chat client", err)
		return err
	}

	room := CreateRoom("HydraChat")

	go func() {
		//Monitoring this channel to see if a program is closing
		channel := make(chan os.Signal)
		//if the operating system throws SIGINT or SIGTERM, it will notify the channel
		signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM)
		//this will block the other code in this goRoutine, until the channel receives a terminate signal
		<-channel

		listener.Close()
		fmt.Println("Closing tcp connection")
		//Sends a notification to all functions listening to the room.Quit channel
		close(room.Quit)
		if room.ClCount() > 0 {
			<-room.Msgch
		}
		os.Exit(0)
	}()

	for {
		//listens for when someone tries to connect to our server
		conn, err := listener.Accept()
		if err != nil {
			logger.Println("Error accepting connection form chat client", err)
			break
		}
		go handleConnection(room, conn)
	}

	return err
}

func handleConnection(room *room, clientConnection net.Conn) {
	logger.Println("Received request from client", clientConnection.RemoteAddr())
	room.AddClient(clientConnection)
}
