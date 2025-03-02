package controller

import (
	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model/database"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type fileController struct{}

func (*fileController) GetThumbByID(c *gin.Context, id string) {
	f := database.CacheFile{}
	err := global.DB.First(&f, id).Error
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(404)
	}

	c.File(f.Path)
}

func (*fileController) GetMediaByID(c *gin.Context, id string) {
	f := database.ArchiveFile{}
	err := global.DB.First(&f, id).Error
	if err != nil {
		logrus.Error(err)
		c.AbortWithStatus(404)
	}

	c.File(f.Path)
}
