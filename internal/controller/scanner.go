package controller

import (
	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/handler"
	"github.com/Yosorable/gomic/internal/model/response"

	"github.com/gin-gonic/gin"
)

type scannerController struct{}

func (*scannerController) Start(c *gin.Context) {
	if err := handler.ScanDirs(); err != nil {
		response.FailWithError(err, c)
		return
	}
	response.OkWithMessage("创建扫描任务成功", c)
}

func (*scannerController) Status(c *gin.Context) {
	response.OkWithData(gin.H{
		"status": global.IS_SERVER_SCANNING,
		"records": global.TMP_SYNC_RECORD,
	}, c)
}
