package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"helpers"
	"log"
	"net/http"
)

type URL struct {
	ID    string `json:"id"`
	URL   string `json:"url"`
	Title string `json:"title"`
	Desc  string `json:"desc"` // description
	Cat   string `json:"cat"`  // category
}

var urls []URL = []URL{
	{"1", "flipkart.com", "Flipkart", "Online Shopping", "Shopping"},
	{"2", "zomato.com", "Zomato", "Order Food Online", "Food"},
}

func health(c *gin.Context) {
	fmt.Println("HEALTH : OK")
	c.IndentedJSON(http.StatusOK, "API IS HEALTHY AND RUNNING")
}

func getAllURLS(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, urls)
}

func addURL(c *gin.Context) {
	var newURL URL

	if err := c.BindJSON(&newURL); err != nil {
		return
	}

	urls = append(urls, newURL)

	c.IndentedJSON(http.StatusCreated, newURL)

	fmt.Printf("\nURL \"%v\" added successfully", newURL.URL)
}

func findURLbyID(id string) (*URL, error) { // helper function
	for index, url := range urls {
		if url.ID == id {
			return &urls[index], nil
		}
	}

	return nil, errors.New("404: URL not found")
}

func getURLbyID(c *gin.Context) {
	id := c.Param("i") // this "i" is what is inputted

	item, err := findURLbyID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, item)

}

// mai nhi kr rhA delete, db connect krne k baad krunga

func main() {

	dotenvErr := godotenv.Load(".env")

	if dotenvErr != nil {
		log.Fatal("Error loading .env file %f", dotenvErr)
	}

	// sample query for getting all users
	helpers.openDB("mysql", func(db *sql.DB) {
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

	// getting all cards
	helpers.openDB("mysql", func(db *sql.DB) {
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

	// router := gin.Default()

	// router.GET("/", health)

	// router.GET("/urls", getAllURLS)

	// router.GET("/urls/:i", getURLbyID)

	// router.POST("/urls", addURL)

	// router.Run("localhost:8080")

}
