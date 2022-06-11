package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "uraaruta"
	password = "6193"
	dbname   = "pggolang"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	// insert
	// hardcoded
	insertStmt := `insert into "students"("name", "roll") values('John', 1)`
	_, e := db.Exec(insertStmt)
	CheckError(e)

	// dynamic
	insertDynStmt := `insert into "students"("name", "roll") values($1,$2)`
	_, e = db.Exec(insertDynStmt, "Jane", 2)
	CheckError(e)

	//// update
	//updateStmt := `update "Students" set "Name"=$1, "Roll"=$2 where "id"=$3`
	//_, e = db.Exec(updateStmt, "Mary", 3, 2)
	//CheckError(e)
	//
	//deleteStmt := `delete from "Students" where id=$4`
	//_, e = db.Exec(deleteStmt, 1)
	//CheckError(e)

	//rows, err := db.Query(`SELECT "id","name", "roll" FROM "students"`)
	//CheckError(err)
	//
	//defer rows.Close()
	//for rows.Next() {
	//	var id int
	//	var name string
	//	var roll int
	//
	//	err = rows.Scan(&id, &name, &roll)
	//	CheckError(err)
	//
	//	fmt.Println(id, name, roll)
	//}

	CheckError(err)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
