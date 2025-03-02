package controller

import (
	"fmt"
	"strconv"

	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model"
	"github.com/Yosorable/gomic/internal/model/database"
	"github.com/Yosorable/gomic/internal/model/response"
	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type archiveController struct{}

func (*archiveController) GetAllAuthors(c *gin.Context) {
	authors := []database.Author{}
	global.DB.Find(&authors)
	res := []model.AuthorMedia{}
	for _, ele := range authors {
		a := model.AuthorMedia{
			ID:   int(ele.ID),
			Name: ele.Name,
		}
		if ele.CoverFileID != nil {
			a.CoverURL = fmt.Sprintf("/thumb/%d", *ele.CoverFileID)
		}
		res = append(res, a)
	}
	response.OkWithData(res, c)
}

func (*archiveController) GetArchivesByAuthorName(c *gin.Context) {
	name := c.Param("name")
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "50")
	pageNum := 1
	limitNum := 50

	if n, err := strconv.Atoi(page); err == nil && n > 0 {
		pageNum = n
	}
	if n, err := strconv.Atoi(limit); err == nil && n > 0 {
		limitNum = n
	}

	authorID := 0

	global.DB.Model(&database.Author{}).Where("name = ?", name).Select("id").Find(&authorID)

	archives := []database.Archive{}
	global.DB.Order("id DESC").Where("author_id = ?", authorID).Offset((pageNum - 1) * limitNum).Limit(limitNum).Find(&archives)

	var res = make([]model.Archive, 0, len(archives))
	for i := 0; i < len(archives); i++ {
		arc := archives[i]
		marc := model.Archive{
			ID:   int(arc.ID),
			Name: arc.Name,
		}
		if arc.CoverFileID != nil {
			marc.CoverURL = fmt.Sprintf("/thumb/%d", *arc.CoverFileID)
		}
		res = append(res, marc)
	}
	response.OkWithData(res, c)
}

func (*archiveController) GetAllArchives(c *gin.Context) {

	page := c.Query("page")
	limit := c.Query("limit")
	pageNum := 1
	limitNum := 50

	if n, err := strconv.Atoi(page); err == nil && n > 0 {
		pageNum = n
	}
	if n, err := strconv.Atoi(limit); err == nil && n > 0 {
		limitNum = n
	}

	archives := []database.Archive{}
	global.DB.Order("id DESC").Offset((pageNum - 1) * limitNum).Limit(limitNum).Find(&archives)

	var res = make([]model.Archive, 0, len(archives))
	for i := 0; i < len(archives); i++ {
		arc := archives[i]
		marc := model.Archive{
			ID:   int(arc.ID),
			Name: arc.Name,
		}
		if arc.CoverFileID != nil {
			marc.CoverURL = fmt.Sprintf("/thumb/%d", *arc.CoverFileID)
		}
		res = append(res, marc)
	}
	response.OkWithData(res, c)
}

func (*archiveController) GetArchiveFilesByID(c *gin.Context) {
	id := c.Param("archive_id")
	archiveFiles := []database.ArchiveFile{}
	var archive *database.Archive
	err := global.DB.Where("id = ?", id).Limit(1).Find(&archive).Error
	if err != nil || archive == nil {
		if err != nil {
			logrus.Error(err)
		}
		response.FailWithMessage("Archive not found", c)
		return
	}
	err = global.DB.Where("archive_id = ?", id).Find(&archiveFiles).Error
	if err != nil || len(archiveFiles) == 0 {
		if err != nil {
			logrus.Error(err)
		}
		response.FailWithMessage("Archive not found", c)
	}
	res := model.Archive{
		ID:    int(archive.ID),
		Name:  archive.Name,
		Files: []model.ArchiveFile{},
	}
	if archive.CoverFileID != nil {
		res.CoverURL = fmt.Sprintf("/thumb/%d", *archive.CoverFileID)
	}

	for _, ele := range archiveFiles {
		res.Files = append(res.Files, model.ArchiveFile{
			ID:   int(ele.ID),
			URL:  fmt.Sprintf("/media/%d", ele.ID),
			Name: ele.Name,
		})
	}

	response.OkWithData(res, c)
}
