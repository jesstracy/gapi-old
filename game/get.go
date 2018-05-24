package game

import (
	//"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RetrieveAllGames(c *gin.Context) {
	dataContext := c.MustGet("Db").(GameDLInterface) // Db is set in router.go
	games, err := dataContext.RetrieveAllGames()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": games})
	}
}

func RetrieveSingleGame(c *gin.Context) {
	dataContext := c.MustGet("Db").(GameDLInterface)
	gameId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		g, err := dataContext.RetrieveSingleGame(gameId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "notfound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": g})
		}
	}
}
