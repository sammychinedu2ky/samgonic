package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sammychinedu2ky/samgonic/controller"
	"github.com/sammychinedu2ky/samgonic/middlewares"
	"github.com/sammychinedu2ky/samgonic/service"
	gindump "github.com/tpkeeper/gin-dump"
)

// import "fmt"
var (
	videoService    service.VideoService       = service.New()
	VideoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.New()
	server.Static("/css","templates/css")
	server.LoadHTMLGlob("templates/*.html")
	apiRoutes := server.Group("/api", middlewares.BasicAuth())
	{
		apiRoutes.GET("/test", func(ctx *gin.Context) {
			ctx.JSON(200, gin.H{
				"message": "Ok!!",
			})
		})
	
		apiRoutes.GET("/posts", func(ctx *gin.Context) {
			ctx.JSON(200, VideoController.FindAll())
		})
	
		apiRoutes.POST("/videos", func(ctx *gin.Context) {
			err := VideoController.Save(ctx)
			if err != nil {
	
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": err})
			}
		})
	}

	viewRoute := server.Group("/view")
	{
		viewRoute.GET("/videos",VideoController.ShowAll)
	}
	server.Use(gin.Recovery(), middlewares.Logger(), gindump.Dump())
	port := os.Getenv("PORT")
	if port == "" {
		port = "5000"
	}
	server.Run((":" + port))

}
