package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type AppForm struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}
type updateAppForm struct {
	Name    string `json:"name" form:"name"`
	Email   string `json:"email" form:"email"`
	Address string `json:"address" form:"address"`
}

func getCorsConfig() gin.HandlerFunc {
	var origins = []string{"*"}

	return cors.New(cors.Config{

		AllowOrigins: origins,
		AllowMethods: []string{"GET", "POST", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders: []string{"Authorization", "Accept", "Accept-Encoding",
			"Accept-Language", "Connection", "Content-Length",
			"Content-Type", "Host", "Origin", "Referer", "User-Agent", "transformRequest"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}

var forms = []AppForm{
	{
		Name:    "",
		Email:   "",
		Address: "",
	},
}

func main() {
	router := gin.Default()
	router.Use(getCorsConfig())
	db, err := sql.Open("mysql", "root:Prabhat@2022@tcp(127.0.0.1:3306)/crudreact?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Connection Successfully!!!")
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "It Works")
	})

	router.POST("/person", func(c *gin.Context) {
		var user AppForm
		c.Bind(&user)
		log.Println(user)
		if user.Name != "" && user.Email != "" && user.Address != "" {
			if insert, _ := db.Exec(`INSERT INTO crud(Name,Email,Address) VALUES(?,?,?)`, user.Name, user.Email, user.Address); insert != nil {
				_, err := insert.LastInsertId()
				if err == nil {
					content := &AppForm{
						Name:    user.Name,
						Email:   user.Email,
						Address: user.Address,
					}
					c.JSON(http.StatusOK, gin.H{
						"status": "ok",
						"data":   content,
					})
				}
			}
		}
	})

	router.PATCH("/persons/id", func(c *gin.Context) {
		fmt.Println("id")
	})

	router.DELETE("/person/:id", func(c *gin.Context) {
		var user AppForm
		fmt.Println(user)
	})
	router.Run(":8000")
}
