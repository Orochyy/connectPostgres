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

var price = map[int]float32{
	1: 100,
	2: 300,
	3: 500,
	4: 1000,
	5: 5000,
}
var typep = map[int]string{
	1: "laser",
	2: "jet",
	3: "matrix",
}
var color = map[int]bool{
	1: false,
	2: true,
}
var screen = map[int]float32{
	1: 100,
	2: 300,
	3: 500,
	4: 1000,
	5: 5000,
}

var hd = map[int]float32{
	1: 100,
	2: 300,
	3: 500,
	4: 1000,
	5: 5000,
}
var ram = map[int]float32{
	1: 100,
	2: 300,
	3: 500,
	4: 1000,
	5: 5000,
}
var cd = map[int]float32{
	1: 100,
	2: 300,
	3: 500,
	4: 1000,
	5: 5000,
}
var speed = map[int]float32{
	1: 100,
	2: 300,
	3: 500,
	4: 1000,
	5: 5000,
}
var maker = map[int]string{
	1: "yura",
	2: "vlad",
}
var model = map[int]string{
	1: "laptop",
	2: "pc",
	3: "printer",
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

	insertPc := `insert into "pc"("model", "speed","ram","hd","cd","price") values($1,$2,$3,$4,$5,$6)`
	_, e := db.Exec(insertPc, model[1], speed[1], ram[1], hd[1], cd[1], price[1])
	CheckError(e)

	insertLaptop := `insert into "laptop"("model", "speed","ram","hd","price","screen") values($1,$2,$3,$4,$5,$6)`
	_, e = db.Exec(insertLaptop, model[1], speed[1], ram[1], hd[1], price[1], screen[1])
	CheckError(e)

}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
