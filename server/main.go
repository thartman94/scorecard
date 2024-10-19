package main

import (

	// "fmt"

	"log"
	"net/http"

	// "net/http"
	// "strings"

	"github.com/gin-gonic/gin"
	// "github.com/thartman94/scorecard/models"
)

func main() {
	var err error

	// db, err = sql.Open("postgres", "postgres://postgres:postgres@localhost/mydb?sslmode=disable")
	router := gin.Default()

	if err != nil {
		log.Fatal(err)
	}

	router.GET("/ping", func(context *gin.Context) {
		context.JSON(http.StatusOK, "pong")
	})

	router.GET("/pong", func(context *gin.Context) {
		context.JSON(http.StatusOK, "poop")
	})

	// router.GET("/player", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, players)
	// })

	// router.GET("/player/:lastName", func(c *gin.Context) {
	// 	lastName := c.Param("lastName")
	// 	player, err := players.find(lastName)
	// 	if err != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	}

	// 	c.JSON(http.StatusOK, player)
	// })

	// r.POST("/player", func(c *gin.Context) {
	// 	var player models.Player
	// 	if err := c.ShouldBindJSON(&player); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}

	// 	players = append(players, player)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Player created successfully",
	// 		"player":  player,
	// 	})

	// })

	// r.PUT("/player/:lastName", func(c *gin.Context) {
	// 	lastName := c.Param("lastName")
	// 	player, err := players.find(lastName)
	// 	if err != nil {
	// 		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	// 	}

	// 	if err := c.ShouldBindJSON(&player); err != nil {
	// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 		return
	// 	}
	// })

	router.Run()
}
