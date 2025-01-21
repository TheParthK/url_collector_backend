package helpers

import (
	"database/sql"
	"fmt"
	"os"
)

func openDB(driverName string, runnable func(db *sql.DB)) {
	tidb_user := os.Getenv("TIDB_USER")
	tidb_password := os.Getenv("TIDB_PASSWORD")
	tidb_host := os.Getenv("TIDB_HOST")
	tidb_port := os.Getenv("TIDB_PORT")
	tidb_db_name := os.Getenv("TIDB_DB_NAME")
	use_ssl := os.Getenv("USE_SSL")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&tls=%s",
		tidb_user, tidb_password, tidb_host, tidb_port, tidb_db_name, use_ssl)
	db, err := sql.Open(driverName, dsn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	runnable(db)
}

func QueryAllUsers() {
	openDB("mysql", func(db *sql.DB) {
		sqlQuery := "SELECT * FROM users"
		rows, err := db.Query(sqlQuery)

		if err != nil {
			fmt.Printf("Error with db: %s", err)
		}

		defer rows.Close()

		uid, user_name := 0, ""
		for rows.Next() {
			err = rows.Scan(&uid, &user_name)
			if err == nil {
				fmt.Printf("uid: %d\tuser_name: %s\n", uid, user_name)
			}
		}
	})
}

func QueryAllCards() {
	openDB("mysql", func(db *sql.DB) {
		sqlQuery := "SELECT * FROM cards"
		rows, err := db.Query(sqlQuery)

		if err != nil {
			fmt.Printf("Error with db: %s", err)
		}

		defer rows.Close()

		cid, uid, title, description, category, url := 0, 0, "", "", "", ""
		for rows.Next() {
			err = rows.Scan(&cid, &uid, &title, &description, &category, &url)
			if err == nil {
				fmt.Printf("cid: %d, uid: %d, title: %s, description: %s, category: %s, url: %s\n", cid, uid, title, description, category, url)
			}
		}
	})
}
