package memberships

import (
	"database/sql"
	"log"
)

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	rows, err := db.Query("SELECT id, email FROM users")
	if err != nil {
		log.Println("error query", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int64
		var email string
		err = rows.Scan(&id, &email)
		if err != nil {
			log.Println("error scan", err)
		}
		log.Printf("id: %d, email: %s \n\n", id, email)
	}
	return &repository{
		db: db,
	}
}