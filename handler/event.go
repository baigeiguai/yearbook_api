package handler

import (
	"BaigeiCode/yearbook_api/db"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HanlderGetEventByYear(c *gin.Context) {
	yearStr := c.Query("year")
	year, err := strconv.ParseInt(yearStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	evnets, err := db.MultiGetEvents(year)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"events": evnets,
	})
}
