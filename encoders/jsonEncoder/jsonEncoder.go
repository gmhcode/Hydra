package main

import (
	"encoding/json"
	"log"
	"os"
)

type CrewMember struct {
	ID                int      `json:"id,omitempty"`
	Name              string   `json:"name"`
	SecurityClearance int      `json:"clearancelevel"`
	AccessCodes       []string `json:"accessCodes"`
}

type ShipInfo struct {
	ShipId    int
	ShipClass string
	Captain   CrewMember
}

func main() {
	f, err := os.Create("jFile.json")
	PrintFatalError(err)
	defer f.Close()

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{1, "Fighter", cm}

	err = json.NewEncoder(f).Encode(&si)
	// writes to file "jFile.json"
	// {"ShipId":1,"ShipClass":"Fighter","Captain":{"name":"Jaro","clearancelevel":10,"accessCodes":["ADA","LAL"]}}
	PrintFatalError(err)
}

// PrintFatalError - Prints an error
func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
