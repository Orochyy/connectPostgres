package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"os"
)

const (
	port = 5432
)

var categoryName = map[int]string{
	1: "laptop",
	2: "pc",
	3: "printer",
	4: "potato",
	5: "apple",
}
var shopName = map[int]string{
	1: "ashan",
	2: "mcdonalds",
	3: "jabko",
	4: "atb",
	5: "blizenko",
}
var imageName = map[int]string{
	1: "imageName1",
	2: "imageName2",
	3: "imageName3",
	4: "imageName4",
	5: "imageName5",
}
var email = map[int]string{
	1: "email1@gmail.com",
	2: "email2@gmail.com",
	3: "email3@gmail.com",
	4: "email4@gmail.com",
	5: "email5@gmail.com",
}

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	var (
		host     = os.Getenv("HOST")
		user     = os.Getenv("USER")
		password = os.Getenv("PASS")
		dbname   = os.Getenv("DBNAME")
	)

	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)

	defer db.Close()

	insertIntoShops := `insert into "shops"("name","email") values($1,$2)`
	_, e := db.Exec(insertIntoShops, shopName[5], email[5])
	CheckError(e)

	insertIntoCategory := `insert into "categories"("name") values($1)`
	_, e = db.Exec(insertIntoCategory, categoryName[5])
	CheckError(e)

	insertIntoImage := `insert into "images"("alt") values($1)`
	_, e = db.Exec(insertIntoImage, imageName[5])
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
