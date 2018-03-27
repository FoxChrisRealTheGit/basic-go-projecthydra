package main

import (
	// "encoding/csv"
	// "io"
	"log"
	// "os"
	// "strconv"
	"sync"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type crewMember struct {
	ID int `bson:"int"`
	Name string `bson"name"`
	SecClearance int `bson:"security clearance"`
	Position string `bson:"position"`
}

type Crew []crewMember

func main(){
	//session, err := mgo.Dial("localhost")
	session, err := mgo.Dial("mongodb://127.0.0.1")
	if err != nil{
		log.Fatal(err)
	}
	defer session.Close()

	//get collection
	personnel := session.DB("Hydra").C("Personnel")

		//get number of documents in the collection
	n, _ := personnel.Count()
	log.Println("Number of personnel is ", n)
//perform simple query
cm := crewMember{}
personnel.Find(bson.M{"id":3}).One(&cm)
log.Println(cm)

//query with expression
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
	if err != nil{
		log.Fatal(err)
	}

	//use select to get names only
// names if of type []struct{Name string}
	names := []struct {
		Name string `bson:"name"`
	}{}

	err = personnel.Find(query).Select(bson.M{"name": 1}).All(&names)
	if err != nil{
		log.Fatal(err)
	}


	//insert
	// newcr := crewMember{ID: 18, Name: "Some Name", SecClearance: 4, Position: "Biologist"}
	// if err := personnel.Insert(newcr); err !=nil{
	// log.Fatal(err)
	// }

//update
err = personnel.Update(bson.M{"id": 16}, bson.M{"$set": bson.M{"position": "Engineer III"}})
if err != nil{
	log.Fatal(err)
}

	//remove
// if err := personnel.Remove(bson.M{"id": 18}); err != nil{
// 	log.Fatal(err)
// }


	//concurrent access
	var wg sync.WaitGroup
	count, _ := personnel.Count()
	wg.Add(count)
	for i := 1; 1 <= count; i++{
		go readId(i, session.Copy(), &wg)
	}
	wg.Wait()

	//CSVToMongo(session.DB("Hydra").C("Personnel"))
}

func readId(id int, sessionCopy *mgo.Session, wg *sync.WaitGroup){
	defer func(){
		sessionCopy.Close()
		wg.Done()
	}()
	p:= sessionCopy.DB("Hydra").C("Personnel")
	cm := crewMember{}
	err := p.Find(bson.M{"id": id}).One(&cm)
	if err != nil{
		return
	}

}

// func CSVToMongo(c *mgo.Collection){
// 	file, err := os.Open("Crews.csv")
// 	defer file.Close()
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	r := csv.NewReader(file)
// 	r.Comment = '#'
// 	var crew[]interface{}
// }

