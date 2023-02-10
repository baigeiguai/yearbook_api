package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerEcho(c *gin.Context) {
	x := c.Query("x")
	c.String(http.StatusOK, x)
}
