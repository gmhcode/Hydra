package main

import (
	"log"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type crewMember struct {
	ID           int    `bson:"id"`
	Name         string `bson:"name"`
	SecClearance int    `bson:"security clearance"`
	Position     string `bson:"position"`
}

type Crew []crewMember

func main() {
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()
	//Get collection
	personnel := session.DB("Hydra").C("Personnel")

	//Get number of documents in the collection
	n, _ := personnel.Count()
	log.Println("Number of personnel is ", n)

	//Perform simple query
	readCrewMember(personnel)
	// prints
	// crew member {3 Lorette Gee   2 Assistant Pilot}

	//Query with expression
	query := bson.M{
		"security clearance": bson.M{
			"$gt": 3,
		},
		"position": bson.M{
			"$in": []string{"Mechanic", "Biologist"},
		},
	}

	var crew Crew
	err = personnel.Find(query).All(&crew)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Query results: ", crew)
	// prints
	//  [{4 Telma Rosas   5 Mechanic} {6 Shanika Reller   5 Biologist} {11 Noble Luczynski   4 Mechanic}]

	//this is called an anonymouns struct
	names := []struct {
		Name string `bson:"name"`
	}{}

	err = personnel.Find(query).Select(bson.M{"name": 1}).All(&names)
	if err != nil {
		log.Fatal(err)
	}
	// prints the names of what we found in "query"
	log.Println(names)
	// prints
	// [{Telma Rosas  } {Shanika Reller  } {Noble Luczynski  }]

	insertCrewMember(personnel)
	updateExisting(personnel)
	// removeCrewmember(personnel)

	concurrentRead(personnel, session)
}

/*
..######..########..########....###....########.########
.##....##.##.....##.##.........##.##......##....##......
.##.......##.....##.##........##...##.....##....##......
.##.......########..######...##.....##....##....######..
.##.......##...##...##.......#########....##....##......
.##....##.##....##..##.......##.....##....##....##......
..######..##.....##.########.##.....##....##....########
*/
func insertCrewMember(personnel *mgo.Collection) {
	newCrewMember := crewMember{ID: 18, Name: "Kaya Gal", SecClearance: 4, Position: "Biologist"}
	if err := personnel.Insert(newCrewMember); err != nil {
		log.Fatal(err)
	}

	//Finding inserted crewMember to make sure it worked
	cm := crewMember{}
	personnel.Find(bson.M{"id": 18}).One(&cm)
	log.Println("newCrewMember: ", cm)
	// prints
	// newCrewMember:  {18 Kaya Gal 4 Biologist}
}

/*
.########..########....###....########.
.##.....##.##.........##.##...##.....##
.##.....##.##........##...##..##.....##
.########..######...##.....##.##.....##
.##...##...##.......#########.##.....##
.##....##..##.......##.....##.##.....##
.##.....##.########.##.....##.########.
*/
func readCrewMember(personnel *mgo.Collection) {
	cm := crewMember{}
	personnel.Find(bson.M{"id": 3}).One(&cm)

	log.Println("crew member", cm)
	// prints
	// crew member {3 Lorette Gee   2 Assistant Pilot}
}

/*
.##.....##.########..########.....###....########.########
.##.....##.##.....##.##.....##...##.##......##....##......
.##.....##.##.....##.##.....##..##...##.....##....##......
.##.....##.########..##.....##.##.....##....##....######..
.##.....##.##........##.....##.#########....##....##......
.##.....##.##........##.....##.##.....##....##....##......
..#######..##........########..##.....##....##....########
*/
func updateExisting(personnel *mgo.Collection) {
	err := personnel.Update(bson.M{"id": 18}, bson.M{"$set": bson.M{"position": "Engineer III"}})
	if err != nil {
		log.Fatal(err)
	}

	cm := crewMember{}
	personnel.Find(bson.M{"id": 18}).One(&cm)
	log.Println("updated crew member: ", cm)
	// prints
	// updated crew member:  {18 Kaya Gal 4 Engineer III}
}

/*
.########..########.##.......########.########.########
.##.....##.##.......##.......##..........##....##......
.##.....##.##.......##.......##..........##....##......
.##.....##.######...##.......######......##....######..
.##.....##.##.......##.......##..........##....##......
.##.....##.##.......##.......##..........##....##......
.########..########.########.########....##....########
*/
func removeCrewmember(personnel *mgo.Collection) {
	var crew Crew
	_ = personnel.Find(bson.M{"id": 18}).All(&crew)

	log.Println("number of crewmembers Found", len(crew))

	for range crew {
		var crew2 Crew
		if err := personnel.Remove(bson.M{"id": 18}); err != nil {
			log.Fatal(err)
		}
		_ = personnel.Find(bson.M{"id": 18}).All(&crew2)
		log.Println("remaining crewmembers", len(crew2))
	}

	cm := crewMember{}
	personnel.Find(bson.M{"id": 18}).One(&cm)
	log.Println("deleted crew member id 18: ", cm)
}

/*
..######...#######..##....##..######.....########..########....###....########.
.##....##.##.....##.###...##.##....##....##.....##.##.........##.##...##.....##
.##.......##.....##.####..##.##..........##.....##.##........##...##..##.....##
.##.......##.....##.##.##.##.##..........########..######...##.....##.##.....##
.##.......##.....##.##..####.##..........##...##...##.......#########.##.....##
.##....##.##.....##.##...###.##....##....##....##..##.......##.....##.##.....##
..######...#######..##....##..######.....##.....##.########.##.....##.########.
*/
// concurrentRead - allows for multiple concurrent queries
func concurrentRead(personnel *mgo.Collection, session *mgo.Session) {
	var wg sync.WaitGroup
	count, _ := personnel.Count()
	wg.Add(count)
	for i := 1; i <= count; i++ {
		go readID(i, session.Copy(), &wg)
	}
	// prints
	// all the crewMembers out of order

	//waits unwill all the wait groups are close in readID()
	wg.Wait()
}

func readID(id int, sessionCopy *mgo.Session, wg *sync.WaitGroup) {
	defer func() {
		sessionCopy.Close()
		wg.Done()
	}()

	p := sessionCopy.DB("Hydra").C("Personnel")
	cm := crewMember{}
	err := p.Find(bson.M{"id": id}).One(&cm)
	log.Println("concurrent crew member", cm)
	if err != nil {
		return
	}

}

// func s() {
// 	file, err := os.Open("Crews.csv")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
