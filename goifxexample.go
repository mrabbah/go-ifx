package main

import (
	"database/sql"
	"fmt"

	_ "github.com/ibmdb/go_ibm_db"
)

func main() {
	// Create a new connection to the Informix database

	conn, err := sql.Open("go_ibm_db", "HOSTNAME=ifx;PORT=9088;"+
		"DATABASE=sysmaster;UID=informix;PWD=in4mix")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()

	// Execute a SELECT statement and retrieve the results
	rows, err := conn.Query("SELECT tabid, tabname, nrows FROM systables")
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Print out the results
	for rows.Next() {
		var tabid int
		var tabname string
		var nrows int
		err = rows.Scan(&tabid, &tabname, &nrows)
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Table ID: %d\nTable Name: %s\nNumber of Rows: %d\n", tabid, tabname, nrows)
	}
}
