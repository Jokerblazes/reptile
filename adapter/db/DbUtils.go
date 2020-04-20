package db

import "database/sql"

func db() *sql.DB {
	db, err := sql.Open("mysql", "root:@/pokemon")
	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}
	return db
}
