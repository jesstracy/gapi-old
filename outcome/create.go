package outcome

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// one problem here, BindJSON doesn't like being passed an id as a string
func CreateOutcome(c *gin.Context) {
	if c.Param("GameId") == "" || c.Param("PlayerId") == "" || c.Param("Result") == "" {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"status": "unprocessable"})
	}
	dataContext := c.MustGet("Db").(OutcomeDLInterface)
	gameId, _ := strconv.Atoi(c.Param("GameId"))
	playerId, _ := strconv.Atoi(c.Param("PlayerId"))
	result := c.Param("Result")
	date := c.Param("Date")
	score, _ := strconv.Atoi(c.Param("Score"))
	if outcome, err := dataContext.CreateOutcome(gameId, playerId, result, date, score); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		c.JSON(http.StatusOK, gin.H{"Id": outcome.Id})
	}
}
