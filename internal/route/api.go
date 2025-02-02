package route

import (
	"github.com/Yosorable/gomic/internal/controller"

	"github.com/gin-gonic/gin"
)

func setUpApi(r *gin.RouterGroup) {
	scanner := r.Group("/scanner")
	{
		ctr := controller.ScannerController
		scanner.POST("/start", ctr.Start)
		scanner.GET("/status", ctr.Status)
	}

	archive := r.Group("/archive")
	{
		ctr := controller.ArchiveController
		archive.GET("/authors", ctr.GetAllAuthorNames)
		archive.GET("/author/:name", ctr.GetArchivesByAuthorName)
		archive.GET("/all", ctr.GetAllArchives)
		archive.GET("/name/:name", ctr.GetArchiveByName)
	}
}
