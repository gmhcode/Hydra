package main

import (
	"encoding/xml"
	"log"
	"os"
)

// attr makes it print <Captain name="Jaro">
type CrewMember struct {
	XMLName           xml.Name `xml:"member"`
	ID                int      `xml:"id,omitempty"`
	Name              string   `xml:"name,attr"`
	SecurityClearance int      `xml:"clearance,attr"`
	AccessCodes       []string `xml:"accessCodes>code"`
	// prints
	//  <accessCodes>
	// 		<code>ADA</code>
}

type ShipInfo struct {
	// tag will now be <ship> instead of <ShipInfo>
	XMLName   xml.Name `xml:"ship"`
	ShipID    int      `xml:"ShipDetails>ShipId"`
	ShipClass string   `xml:"ShipDetails>ShipClass"`
	Captain   CrewMember
}

func main() {
	file, err := os.Create("xmlfile.xml")
	PrintFatalError(err)
	defer file.Close()

	cm := CrewMember{Name: "Jaro", SecurityClearance: 10, AccessCodes: []string{"ADA", "LAL"}}
	si := ShipInfo{ShipID: 1, ShipClass: "Fighter", Captain: cm}

	enc := xml.NewEncoder(file)
	enc.Indent(" ", "	")
	enc.Encode(si)

	if err != nil {
		log.Fatal("Could not encode xml file", err)
	}

	err = xml.NewEncoder(file).Encode(&si)
	PrintFatalError(err)
}

func PrintFatalError(err error) {
	if err != nil {
		log.Fatal("Error happened while processing file", err)
	}
}
