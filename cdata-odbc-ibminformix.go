package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "github.com/alexbrainman/odbc"
// )

// func main() {
// 	db, err := sql.Open("odbc",
// 		"DSN=CData Informix Sys")
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var (
// 		tabname string
// 	)

// 	rows, err := db.Query("SELECT * FROM newtable")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&tabname)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		fmt.Println(tabname)
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer db.Close()
// }
