package controller

import (
	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model"
	"github.com/Yosorable/gomic/internal/model/response"
	"github.com/Yosorable/gomic/internal/utils"

	"github.com/gin-gonic/gin"
)

type archiveController struct{}

func (*archiveController) GetAllAuthorNames(c *gin.Context) {
	names := make([]string, 0, len(global.TMP_MEMORY_DB))
	for _, ele := range global.TMP_MEMORY_DB {
		names = append(names, ele.Name)
	}
	utils.NaturalSort(names)

	response.OkWithData(names, c)
}

func (*archiveController) GetArchivesByAuthorName(c *gin.Context) {
	name := c.Param("name")
	for _, ele := range global.TMP_MEMORY_DB {
		if ele.Name == name {
			var res = make([]model.Archive, 0, len(ele.Archives))
			for i := len(ele.Archives) - 1; i >= 0; i-- {
				arc := ele.Archives[i]
				arc.Files = nil
				res = append(res, arc)
			}
			response.OkWithData(res, c)
			return
		}
	}
	response.FailWithMessage("Author not found", c)
}

func (*archiveController) GetAllArchives(c *gin.Context) {
	res := []model.Archive{}

	db := global.TMP_MEMORY_DB

	for i := len(db) - 1; i >= 0; i-- {
		for j := len(db[i].Archives) - 1; j >= 0; j-- {
			arc := db[i].Archives[j]
			arc.Files = nil
			res = append(res, arc)
		}
	}

	response.OkWithData(res, c)
}

func (*archiveController) GetArchiveByName(c *gin.Context) {
	name := c.Param("name")

	db := global.TMP_MEMORY_DB

	for i := len(db) - 1; i >= 0; i-- {
		for j := len(db[i].Archives) - 1; j >= 0; j-- {
			arc := db[i].Archives[j]
			if arc.Name == name {
				response.OkWithData(arc, c)
				return
			}
		}
	}
	response.FailWithMessage("Archive not found", c)
}
