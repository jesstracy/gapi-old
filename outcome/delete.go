package outcome

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DeleteOutcome(c *gin.Context) {
	dataContext := c.MustGet("Db").(OutcomeDLInterface)
	outcomeId, err := strconv.Atoi(c.Param("Id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": "internalservererror"})
	} else {
		err := dataContext.DeleteOutcome(outcomeId)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"status": "notfound"})
		} else {
			c.JSON(http.StatusOK, gin.H{"status": "successful"})
		}
	}
}
