package main

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type crewMember struct {
	id           int
	name         string
	secClearance int
	position     string
}

type Crew []crewMember

func main() {
	db, err := sql.Open("mysql", "root:password@/hydra?parseTime=true")
	if err != nil {
		log.Fatal("Could not connect, error", err.Error())
	}
	defer db.Close()

	cw := GetCrewByPositions(db, []string{"'Mechanic'", "'Biologist'"})
	fmt.Println(cw)

	// fmt.Println(GetCrewByPositions(db, 11))
}

func GetCrewByPositions(db *sql.DB, positions []string) Crew {
	queryString := fmt.Sprintf("Select id,Name,SecurityClearance,Position from Personnel where Position in (%s);", strings.Join(positions, ","))

	rows, err := db.Query(queryString)
	if err != nil {
		log.Fatal("Could not get data from the Personnel table ", err)
	}
	defer rows.Close()

	retVal := Crew{}
	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	for rows.Next() {
		member := crewMember{}
		err = rows.Scan(&member.id, &member.name, &member.secClearance, &member.position)
		if err != nil {
			log.Fatal("Error scanning row", err)
		}
		retVal = append(retVal, member)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return retVal
}
