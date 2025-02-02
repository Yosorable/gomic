package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model"
	"github.com/Yosorable/gomic/internal/utils"

	"github.com/sirupsen/logrus"
)

func ScanDirs() error {
	if !global.SCANNING_MUTEX.TryLock() {
		if global.IS_SERVER_SCANNING {
			return errors.New("正在执行扫描")
		}
		return errors.New("服务器错误")
	}
	global.IS_SERVER_SCANNING = true
	go func() {
		defer global.SCANNING_MUTEX.Unlock()
		defer func() { global.IS_SERVER_SCANNING = false }()

		// scan dirs
		start := time.Now()
		files, err := os.ReadDir(global.CONFIG.Media)
		if err != nil {
			logrus.Error("scan files error", err)
			return
		}
		var allMedias = []*model.AuthorMedia{}
		for _, dir := range files {
			if !dir.IsDir() {
				continue
			}
			if r, err := scanDirectory(filepath.Join(global.CONFIG.Media, dir.Name())); err == nil {
				allMedias = append(allMedias, r)
			}
		}

		if !utils.FileExists(filepath.Join(global.CONFIG.Data, "thumb")) {
			os.MkdirAll(filepath.Join(global.CONFIG.Data, "thumb"), 0700)
		}

		auid := 0
		arid := 0
		fid := 0
		for i := 0; i < len(allMedias); i++ {
			auid++
			allMedias[i].ID = auid
			for j := 0; j < len(allMedias[i].Archives); j++ {
				arid++
				allMedias[i].Archives[j].ID = arid
				for k := 0; k < len(allMedias[i].Archives[j].Files); k++ {
					if allMedias[i].CoverURL == "" && utils.IsPicture(allMedias[i].Archives[j].Files[k].Name) {
						file := allMedias[i].Archives[j].Files[k]
						ext := filepath.Ext(file.Path)
						coverPath := filepath.Join(global.CONFIG.Data, "thumb", file.UUID) + ext
						coverURL := file.URL
						if !utils.FileExists(coverPath) {
							if err := utils.CreateImageThumb(file.Path, coverPath); err == nil {
								coverURL = "/thumb/" + file.UUID + ext
							}
						} else {
							coverURL = "/thumb/" + file.UUID + ext
						}
						allMedias[i].CoverURL = coverURL
					}
					if allMedias[i].Archives[j].CoverURL == "" && utils.IsPicture(allMedias[i].Archives[j].Files[k].Name) {
						file := allMedias[i].Archives[j].Files[k]
						ext := filepath.Ext(file.Path)
						coverPath := filepath.Join(global.CONFIG.Data, "thumb", file.UUID) + ext
						coverURL := file.URL
						if !utils.FileExists(coverPath) {
							if err := utils.CreateImageThumb(file.Path, coverPath); err == nil {
								coverURL = "/thumb/" + file.UUID + ext
							}
						} else {
							coverURL = "/thumb/" + file.UUID + ext
						}
						allMedias[i].Archives[j].CoverURL = coverURL
					}
					fid++
					allMedias[i].Archives[j].Files[k].ID = fid
				}
			}
		}

		duration := time.Now().Sub(start)
		logrus.Infof("Scan file success, cost: %v", duration)
		global.TMP_SYNC_RECORD = append(
			global.TMP_SYNC_RECORD, 
			fmt.Sprintf("start: %s, cost: %v", start.Format("2006-01-02 15:04:05"), duration),
		)

		global.TMP_MEMORY_DB = allMedias
		if data, err := json.MarshalIndent(allMedias, "", "  "); err == nil {
			if err := os.WriteFile(filepath.Join(global.CONFIG.Data, "db.json"), data, 0700); err == nil {
				logrus.Info("Write to json file success")
			}
		}

	}()

	return nil
}

func scanDirectory(dir string) (*model.AuthorMedia, error) {
	filesMap := make(map[string]map[string][]string)

	var authorName = filepath.Base(dir)
	var scan func(string, int)
	scan = func(path string, level int) {
		files, err := os.ReadDir(path)
		if err != nil {
			logrus.Error("Error reading directory:", err)
			return
		}

		dirName := filepath.Base(path)

		if level == 0 {
			filesMap[dirName] = make(map[string][]string)
		}

		var fileList []string
		for _, file := range files {
			if file.IsDir() {
				subDir := filepath.Join(path, file.Name())
				scan(subDir, level+1)
			} else {
				fileList = append(fileList, file.Name())
			}
		}

		if level > 0 {
			if len(fileList) > 0 {
				curr := strings.TrimPrefix(strings.TrimPrefix(path, dir), string(os.PathSeparator))
				// natural sort files
				utils.NaturalSort(fileList)

				// video and picture files sort
				images := []string{}
				videos := []string{}
				for _, ele := range fileList {
					if utils.IsPicture(ele) {
						images = append(images, ele)
					} else if utils.IsVideo(ele) {
						videos = append(videos, ele)
					}
				}
				var finalFiles = make([]string, 0, len(images)+len(videos))
				finalFiles = append(finalFiles, images...)
				finalFiles = append(finalFiles, videos...)
				filesMap[authorName][curr] = finalFiles
			}
		}
	}

	scan(dir, 0)

	// transform to struct and natural sort folders
	mp := filesMap[authorName]
	res := &model.AuthorMedia{
		Name:     authorName,
		Archives: make([]model.Archive, 0, len(mp)),
	}
	existedMap := make(map[string]int, len(mp))

	archiveName := make([]string, 0, len(mp))
	for k := range mp {
		archiveName = append(archiveName, k)
	}
	utils.NaturalSort(archiveName)
	for _, ele := range archiveName {
		secondFolderName := ele
		archiveName := strings.Split(ele, string(os.PathSeparator))[0]

		idx, exist := existedMap[archiveName]
		if !exist {
			idx = len(res.Archives)
			existedMap[archiveName] = idx
			archive := model.Archive{
				Name:  archiveName,
				Files: make([]model.ArchiveFile, 0, len(mp[ele])),
			}
			res.Archives = append(res.Archives, archive)
		}

		for _, f := range mp[ele] {
			path := filepath.Join(dir, string(os.PathSeparator)+secondFolderName+string(os.PathSeparator)+f)

			md5, err := utils.CalculateStringMD5(path)
			if err != nil {
				logrus.Error("Calculate file md5 error:", err)
				continue
			}
			res.Archives[idx].Files = append(res.Archives[idx].Files, model.ArchiveFile{
				UUID: md5,
				Name: f,
				Path: path,
				URL:  global.MEDIA_SERVER_PREFIX + strings.ReplaceAll(strings.TrimPrefix(path, global.CONFIG.Media), string(os.PathSeparator), "/"),
			})
		}
	}

	return res, nil
}
