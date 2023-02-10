package handler

import (
	"BaigeiCode/yearbook_api/crawl"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandlerCrawlSudaEvents(c *gin.Context) {
	yearStr := c.Query("year")
	year, err := strconv.ParseInt(yearStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = crawl.CrawlSudaEvents(year)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, "")
}
