package main

//to add: curl -H "Content-Type: application/json" -d '{"name": "Camel Up"}' http://localhost:8080/games
//to delete: curl -X "DELETE" http://localhost:8080/players/4

import (
	"fmt"

	"github.com/jesstracy/gapi/db"
	"github.com/jesstracy/gapi/game"
	"github.com/jesstracy/gapi/outcome"
	"github.com/jesstracy/gapi/player"
	"github.com/gin-gonic/gin"
)

func main() {

	db, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(player.Player{})
	db.AutoMigrate(game.Game{})
	db.AutoMigrate(outcome.Outcome{})

	router := gin.Default()
	router.Use(gin.Recovery())
	Routes(router)

	router.Run(":8080")
}
