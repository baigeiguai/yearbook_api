package routers

import (
	"BaigeiCode/yearbook_api/handler"
	"BaigeiCode/yearbook_api/middlewares"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetupRouters() error {
	r := gin.New()
	r.Use(middlewares.Cors())
	apiGroup := r.Group("/api/")
	{
		apiGroup.GET("/echo", handler.HandlerEcho)
		apiGroup.GET("/year2events", handler.HanlderGetEventByYear)
		apiGroup.POST("/crawl", handler.HandlerCrawlSudaEvents)
	}
	if err := http.ListenAndServe(":8000", r); err != nil {
		return err
	}
	return nil
}
