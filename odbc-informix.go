package main
// isql -v Infdrv2 informix in4mix
import (
	"database/sql"
	"fmt"
	"log"
	
	 _ "github.com/alexbrainman/odbc"
)

func main() {
	db, err := sql.Open("odbc","DSN=Infdrv2;UID=informix;PWD=in4mix")
	if err != nil {
		log.Fatal(err)
	}

	var (
		tabname string
	)

	rows, err := db.Query("SELECT tabname FROM systables")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&tabname)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(tabname)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()
}
