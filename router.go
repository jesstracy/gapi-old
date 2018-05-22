package main

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jesstracy/gapi/game"
	"github.com/jesstracy/gapi/outcome"
	"github.com/jesstracy/gapi/player"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.StaticFS("/frontend", http.Dir("./frontend"))
	playerRouter := router.Group("/players")
	playerRouter.Use(PlayerDataContextMW())
	{
		playerRouter.POST("/", player.CreatePlayer)
		playerRouter.GET("/", player.RetrieveAllPlayers)
		playerRouter.GET("/:Id", player.RetrieveSinglePlayer)
		playerRouter.DELETE("/:Id", player.DeletePlayer)
	}
	gameRouter := router.Group("/games")
	gameRouter.Use(GameDataContextMW())
	{
		gameRouter.POST("/", game.CreateGame)
		gameRouter.GET("/", game.RetrieveAllGames)
		gameRouter.GET("/:Id", game.RetrieveSingleGame)
		gameRouter.DELETE("/:Id", game.DeleteGame)
	}
	outcomeRouter := router.Group("/outcomes")
	{
		outcomeRouter.POST("/", outcome.CreateOutcome)
		outcomeRouter.GET("/", outcome.OutcomeIndex)
		outcomeRouter.GET("/:Id", outcome.ShowOutcome)
		outcomeRouter.DELETE("/:Id", outcome.DeleteOutcome)
	}
	//router.Use(GameDataContextMW())
	//router.Use(PlayerDataContextMW())
}

func GameDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		gameDl := &game.GameDLGorm{}
		c.Set("Db", gameDl)
		c.Next()
	}
}

func PlayerDataContextMW() gin.HandlerFunc {
	return func(c *gin.Context) {
		playerDl := &player.PlayerDLGorm{}
		c.Set("Db", playerDl)
		c.Next()
	}
}
