package main

import (
	"database/sql"
	"encoding/csv"
	_ "encoding/csv"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type user struct {
	Uid     int `gorm:"primaryKey"`
	Name    string
	Phone   string
	Email   string
	HashPwd string `gorm:"columns:password"`
}

func (user) TableName() string {
	return "user"
}

var dbDSN = "host=localhost port=5432 user=admin password=password dbname=tmpdb sslmode=disable"

func queryUsersByStdLib() []user {
	db, err := sql.Open("postgres", dbDSN)
	if err != nil {
		log.Fatalln(err)
	}
	db.Close()

	rows, err := db.Query(`SELECT uid, name, phone, email, password FROM "user"`)
	if err != nil {
		log.Fatalln(err)
	}
	defer rows.Close()

	users := []user{}
	u := user{}
	for rows.Next() {
		err := rows.Scan(&u.Uid, &u.Name, &u.Phone, &u.Email, &u.HashPwd)
		if err != nil {
			log.Fatalln(err)
		}
		users = append(users, u)
	}
	return users
}

func queryUsersByOrmLib() []user {
	gormDB, err := gorm.Open(postgres.Open(dbDSN), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}

	users := []user{}
	gormDB.Model(&user{}).Find(&users)
	return users
}

func exportCsv(to string, table [][]string) {
	f, err := os.Create(to)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	w.WriteAll(table)
}

func main() {
	rows := [][]string{{"ID", "Name", "Phone", "Mail", "Hash"}}
	// users := queryUsersByStdLib()
	users := queryUsersByOrmLib()
	for _, u := range users {
		columns := []string{}
		columns = append(columns, strconv.Itoa(u.Uid))
		columns = append(columns, u.Name)
		columns = append(columns, u.Phone)
		columns = append(columns, u.Email)
		columns = append(columns, u.HashPwd)
		rows = append(rows, columns)
	}
	exportCsv("./gen/exportUsers.csv", rows)
}
