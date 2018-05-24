package outcome

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func RetrieveAllOutcomes(c *gin.Context) {
	dataContext := c.MustGet("Db").(OutcomeDLInterface)
	outcomes, err := dataContext.RetrieveAllOutcomes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": http.StatusInternalServerError})
	} else {
		c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": outcomes})
	}
}

func RetrieveSingleOutcome(c *gin.Context) {
	dataContext := c.MustGet("Db").(OutcomeDLInterface)
	outcomeId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		o, err := dataContext.RetrieveSingleOutcome(outcomeId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "notfound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": o})
		}
	}
}
