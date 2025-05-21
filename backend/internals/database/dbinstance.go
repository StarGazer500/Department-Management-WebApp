package database

import (
	"database/sql"
	"fmt"
)
var Dbinstance *sql.DB

func SetDb(db *sql.DB ){
	Dbinstance = db
	fmt.Println("instance",Dbinstance)
}
