package handler

import (
	"BaigeiCode/yearbook_api/crawl"
	"BaigeiCode/yearbook_api/db"
	"BaigeiCode/yearbook_api/rpc"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"events": evnets,
	})
}
func HandlerEventsDetail(c *gin.Context) {
	uri := c.Query("uri")
	fmt.Println(uri)
	content, err := crawl.CrawlEventDetail(uri)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	req, _ := json.Marshal(map[string]string{
		"content": content,
	})
	resp, err := rpc.DoBytesPost("http://10.37.156.42:8001/labapi/entity_extract", req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}
	c.String(http.StatusOK, string(resp))
}
