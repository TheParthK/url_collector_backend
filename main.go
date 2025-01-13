package main

import (
	// "context"
	"errors"
	"fmt"
	// "log"
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

// CONNECTING DB

// func connectToDB() {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

// 	defer cancel()
// 	connectionString := "mongodb://url-manager:pA9UUH0klrDEklpEK4jP75tjFSfdbzJkUkD5rmeTKAnbne9EjI1TvpG24D7EFVG0f1qjwV1swT0FACDbETcySQ==@url-manager.mongo.cosmos.azure.com:10255/?ssl=true&replicaSet=globaldb&retrywrites=false&maxIdleTimeMS=120000&appName=@url-manager@"

// 	var err error

// 	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))

// 	if err != nil {
// 		log.Fatalf("Failed to connect to DB: %v", err)
// 	}

// 	fmt.Println("Connected to DB successfully")
// }

// BACKEND

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

func main() {

	// connectToDB()

	router := gin.Default()

	router.GET("/", health)

	router.GET("/urls", getAllURLS)

	router.GET("/urls/:i", getURLbyID)

	router.POST("/urls", addURL)

	router.Run("localhost:8080")

}
