package main

import (
	"flag"
	"log"
	"strings"

	"github.com/Hydra/hydracommlayer"
	"github.com/Hydra/hydracommlayer/hydraproto"
)

func main() {
	op := flag.String("type", "", "Server (s) or client (c) ?")
	address := flag.String("addr", ":8080", "address? host:port ")
	flag.Parse()

	switch strings.ToUpper(*op) {
	case "S":
		runServer(*address)
	case "C":
		runClient(*address)
	}
}

func runServer(dest string) {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	recvChan, err := c.ListenAndDecode(dest)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range recvChan {
		log.Println("received ", msg)
	}
}

func runClient(dest string) {
	c := hydracommlayer.NewConnection(hydracommlayer.Protobuf)
	ship := &hydraproto.Ship{
		Shipname:    "Hydra",
		CaptainName: "Jala",
		Crew: []*hydraproto.Ship_CrewMember{
			&hydraproto.Ship_CrewMember{Id: 1, Name: "Kevin", SecClearance: 5, Position: "Pilot"},
			&hydraproto.Ship_CrewMember{Id: 1, Name: "Jade", SecClearance: 4, Position: "Tech"},
			&hydraproto.Ship_CrewMember{Id: 1, Name: "Wally", SecClearance: 3, Position: "Enginneer"},
		},
	}
	if err := c.EncodeAndSend(ship, dest); err != nil {
		log.Println("Error occured while sending message", err)
	} else {
		log.Println("Send Operation Successful")
	}
}
