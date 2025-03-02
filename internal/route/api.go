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
		scanner.POST("/status", ctr.Status)
	}

	archive := r.Group("/archive")
	{
		ctr := controller.ArchiveController
		archive.POST("/authors", ctr.GetAllAuthors)
		archive.POST("/author/:name", ctr.GetArchivesByAuthorName)
		archive.POST("/all", ctr.GetAllArchives)
		archive.POST("/files/:archive_id", ctr.GetArchiveFilesByID)
	}
}
