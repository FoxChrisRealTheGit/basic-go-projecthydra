package main

import(
	"database/sql"
	"log"

	_"github.com/go-sql-driver/mysql"
	// "encoding/csv"
	"fmt"
	// "io"
	// "os"
	"strings"
)

type crewMember struct{
	id	int
	name	string
	secClearance	int
	position	string
}

type Crew []crewMember

func main(){
	db, err := sql.Open("mysql", "gouser:gouser@/Hydra?parseTime=true")
	if err != nil{
		log.Fatal("Could not connect, error", err.Error())
	}
	defer db.Close()

	cw := GetCrewByPositions(db, []string{"'Mechanic'", "'Biologist'"})
	fmt.Println(cw)

	fmt.Println(GetCrewMemberById(db, 11))

	//AddCrewMember(db, crewMember{name: "name person", setClearance: 4, position: "Biologist"})

	//fmt.Println()

	/*
	cr:= Crew{
		crewMember{name: "some name", setClearance: 4, psoition: "Chemist"},
		crewMember{name: "another name", setClearance: 5, position: "Biologist"}
	}
	CreateCrewMembersByTransaction(db, cr)
	*/
}

func GetCrewByPositions(db *sql.DB, positions []string) Crew{
	Qs := fmt.Sprintf("SELECT id, Name, SecurityClearance, Position from Personal where Position in (%s);", strings.Join(positions, ","))

	rows, err := db.Query(Qs)
	if err != nil{
		log.Fatal("Could not get data from the Personal table ", err)
	}
	defer rows.Close()

	retVal := Crew{}
	cols, _ := rows.Columns()
	fmt.Println("Columns detected: ", cols)

	for rows.Next() {
		member := crewMember{}
		err = rows.Scan(&member.id, &member.name, &member.secClearance, &member.position)
		if err != nil{
			log.Fatal("Error scanning row", err)
		}
		retVal = append(retVal, member)
	}
	return retVal
}

func GetCrewMemberById(db *sql.DB, id int) (cm crewMember){
	row := db.QueryRow("Select * from Personal where id = ?", id)

	err := row.Scan(&id, &cm.name, &cm.secClearance, &cm.position)
	if err != nil{
		log.Fatal(err)
	}
	return
}

func AddCrewMember(db *sql.DB, cm crewMember) int64{
	res, err := db.Exec("")
	if err != nil{
		log.Fatal(err)
	}
	ra, _ := res.RowsAffected()
	id, _:= res.LastInsertId()

	log.Println("Rows Affected", ra, "Last inserted id", id)
	return id
}

// func GetCrewMemberByPosition(db *sql.DB, position string) (cr Crew){
// 	stmt, err := db.Prepare("Select * from Personal where Position = ?")
// 	if err != nil{
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()

// }