package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"example.com/openapi"
)

func Index(c *gin.Context) {
	session := sessions.Default(c)
	var mysession int
	if v := session.Get("mysession"); v != nil {
		switch v.(type) {
		case string:
			mysession, _ = strconv.Atoi(v.(string))
			mysession++
			session.Set("mysession", strconv.Itoa(mysession))

			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}
		default:
			session.Delete("mysession")
			session.Set("mysession", strconv.Itoa(1))
			mysession = 1
			if err := session.Save(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
				return
			}
		}
	} else {
		fmt.Println("no session key.")
		session.Set("mysession", strconv.Itoa(1))

		if err := session.Save(); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save session"})
			return
		}
		mysession = 1
	}

	str := fmt.Sprintf("this is server 1, counter = %d\n", mysession)

	c.JSON(http.StatusOK, str)
}

func main() {
	log.Printf("Server started")

	// DefaultApiService := openapi.NewDefaultApiService()
	MyApiService := openapi.NewMyApiService()
	DefaultApiController := openapi.NewDefaultApiController(MyApiService)

	router := openapi.NewRouter(DefaultApiController)

	log.Fatal(http.ListenAndServe(":8080", router))

	// r := gin.Default()
	// store, _ := redis.NewStore(10, "tcp", "redis:6379", "", []byte("secret"))
	// r.Use(sessions.Sessions("session", store))
	// r.GET("/", Index)

	// r.Run(":8080")
}
