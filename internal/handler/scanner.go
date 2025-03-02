package handler

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yosorable/gomic/internal/global"
	"github.com/Yosorable/gomic/internal/model/database"
	"github.com/Yosorable/gomic/internal/utils"
	"github.com/sirupsen/logrus"
)

func ScanArchives() error {
	if !global.SCANNING_MUTEX.TryLock() {
		if global.IS_SERVER_SCANNING {
			return errors.New("scanning")
		}
		return errors.New("server error")
	}
	global.IS_SERVER_SCANNING = true

	go func() {
		var err error = nil

		defer global.SCANNING_MUTEX.Unlock()
		defer func() { global.IS_SERVER_SCANNING = false }()

		start := time.Now()
		defer func() {
			if err == nil {
				duration := time.Since(start)
				logrus.Infof("Scan file success, cost: %v", duration)
				global.TMP_SYNC_RECORD = append(global.TMP_SYNC_RECORD, fmt.Sprintf(
					"start: %s, cost: %v",
					start.Format("2006-01-02 15:04:05"),
					duration,
				))
			} else {
				logrus.Infof("Scan file failed: %v", err)
			}
		}()

		files, err := os.ReadDir(global.CONFIG.MediaPath)

		if err != nil {
			return
		}

		for _, dir := range files {
			if !dir.IsDir() {
				continue
			}
			err = scanForAuthor(dir.Name())
			if err != nil {
				return
			}
		}

	}()

	return nil
}

func scanForAuthor(name string) error {
	db := global.DB
	authorPath := filepath.Join(global.CONFIG.MediaPath, name)
	var author database.Author

	// find or create author
	if err := db.Limit(1).Find(&author, "name = ?", name).Error; err != nil {
		return err
	}

	if author.ID == 0 {
		author = database.Author{
			Name: name,
		}
		if err := db.Create(&author).Error; err != nil {
			return err
		}
	}

	// find author's archives
	archives := []database.Archive{}
	if err := db.Find(&archives, "author_id = ?", author.ID).Error; err != nil {
		return err
	}
	archiveNameMap := map[string]database.Archive{}
	for _, ele := range archives {
		archiveNameMap[ele.Name] = ele
	}

	// scan author folder's folders (archives on disk)
	filesMap := make(map[string][]string) // archive->files
	var scan func(string, int)
	scan = func(path string, level int) {
		files, err := os.ReadDir(path)
		if err != nil {
			logrus.Errorf("Error reading directory: %v", err)
			return
		}

		var fileList []string
		for _, file := range files {
			if file.IsDir() {
				if level == 0 {
					filesMap[file.Name()] = []string{}
				}
				subDir := filepath.Join(path, file.Name())
				scan(subDir, level+1)
			} else {
				// ignore level 0's files
				if level > 0 {
					fileList = append(fileList, file.Name())
				}
			}
		}

		if level > 0 {
			if len(fileList) > 0 {
				curr := strings.TrimPrefix(strings.TrimPrefix(path, authorPath), string(os.PathSeparator))
				archiveName := strings.Split(curr, string(os.PathSeparator))[0]
				// natural sort files
				utils.NaturalSort(fileList)
				// video and picture files sort
				images := []string{}
				videos := []string{}
				for _, ele := range fileList {
					filePath := filepath.Join(path, ele)
					nameInArchive := strings.TrimPrefix(
						strings.TrimPrefix(filePath, filepath.Join(authorPath, archiveName)),
						string(os.PathSeparator),
					)
					if utils.IsPicture(ele) {
						images = append(images, nameInArchive)
					} else if utils.IsVideo(ele) {
						videos = append(videos, nameInArchive)
					}
				}
				var finalFiles = make([]string, 0, len(images)+len(videos))
				finalFiles = append(finalFiles, images...)
				finalFiles = append(finalFiles, videos...)
				filesMap[archiveName] = append(filesMap[archiveName], finalFiles...)
			}
		}
	}
	scan(authorPath, 0)

	// natural sort archive (on disk) names
	archiveOnDiskNames := make([]string, 0, len(filesMap))
	for k, _ := range filesMap {
		archiveOnDiskNames = append(archiveOnDiskNames, k)
	}
	utils.NaturalSort(archiveOnDiskNames)

	// handle db
	for _, aname := range archiveOnDiskNames {
		arcOnDB, exist := archiveNameMap[aname]
		arcFilesOnDisk := filesMap[aname]

		// new archive with new files
		if !exist {
			var auID uint = author.ID
			newArchive := &database.Archive{
				Name:     aname,
				AuthorID: &auID,
			}
			if err := db.Create(newArchive).Error; err != nil {
				return err
			}

			var arcFiles = make([]*database.ArchiveFile, 0, len(arcFilesOnDisk))
			for _, item := range arcFilesOnDisk {
				newArchiveFile := &database.ArchiveFile{
					Name:      item,
					Path:      filepath.Join(authorPath, aname, item),
					ArchiveID: newArchive.ID,
				}
				if utils.IsPicture(item) {
					newArchiveFile.FileType = database.PICTURE_FILE
				} else if utils.IsVideo(item) {
					newArchiveFile.FileType = database.VIDEO_FILE
				} else {
					// unsupported files
					continue
				}
				arcFiles = append(arcFiles, newArchiveFile)
			}

			db.Create(arcFiles)

			continue
		}

		// exist archive
		arcFilesOnDB := []*database.ArchiveFile{}
		if err := db.Find(&arcFilesOnDB, "archive_id = ?", arcOnDB.ID).Error; err != nil {
			return err
		}

		arcFilesToAddFromDisk := []*database.ArchiveFile{}
		arcFileIDsToDelete := []uint{}
		arcFileDBNameToID := map[string]uint{}
		arcFileDiskNameMap := map[string]bool{}

		for _, ele := range arcFilesOnDB {
			arcFileDBNameToID[ele.Name] = ele.ID
		}
		for _, ele := range arcFilesOnDisk {
			arcFileDiskNameMap[ele] = true
		}

		for _, ele := range arcFilesOnDisk {
			_, exist := arcFileDBNameToID[ele]
			if exist {
				continue
			}
			arcFilesToAddFromDisk = append(arcFilesToAddFromDisk, &database.ArchiveFile{
				Name:      ele,
				Path:      filepath.Join(authorPath, aname, ele),
				ArchiveID: arcOnDB.ID,
			})
		}

		for _, ele := range arcFilesOnDB {
			_, exist := arcFileDiskNameMap[ele.Name]
			if exist {
				continue
			}
			arcFileIDsToDelete = append(arcFileIDsToDelete, ele.ID)
		}

		if len(arcFilesToAddFromDisk) > 0 {
			if err := db.Create(arcFilesToAddFromDisk).Error; err != nil {
				return err
			}
		}
		if len(arcFileIDsToDelete) > 0 {
			if err := db.Delete([]database.ArchiveFile{}, arcFileIDsToDelete).Error; err != nil {
				return err
			}
		}
	}
	for _, ele := range archives {
		aname := ele.Name
		_, exist := filesMap[aname]
		if !exist {
			db.Delete(&database.ArchiveFile{}, "archive_id = ?", ele.ID)
			db.Delete(&ele)
		}
	}

	generateCoversForAuthor(author.ID)

	return nil
}

func generateCoversForAuthor(authorID uint) {
	db := global.DB

	author := database.Author{}
	if err := db.Where("id = ?", authorID).Take(&author).Error; err != nil {
		logrus.Errorf("generateCoversForAuthor error: %v", err)
		return
	}

	archives := []database.Archive{}
	if err := db.Where("author_id = ?", authorID).Find(&archives).Error; err != nil {
		logrus.Errorf("generateCoversForAuthor error: %v", err)
		return
	}
	for _, ele := range archives {
		if ele.CoverFileID == nil {
			var coverFilePath string
			db.Model(&database.ArchiveFile{}).Select("path").Where("archive_id = ?", ele.ID).Limit(1).Find(&coverFilePath)
			if utils.IsPicture(coverFilePath) {
				ext := filepath.Ext(coverFilePath)
				hash, err := utils.CalculateStringMD5(coverFilePath)
				if err == nil {
					finalPath := filepath.Join(global.CONFIG.DataPath, "thumb", hash) + ext
					err = utils.CreateImageThumb(
						coverFilePath,
						finalPath,
					)
					if err == nil {
						cacheFile := database.CacheFile{
							Path: finalPath,
						}
						err = db.Create(&cacheFile).Error
						if err == nil {
							cID := cacheFile.ID
							ele.CoverFileID = &cID
							db.Model(&ele).Update("cover_file_id", cID)
						} else {
							logrus.Errorf("generateCoversForAuthor error: %v", err)
						}
					} else {
						logrus.Errorf("generateCoversForAuthor error: %v", err)
					}
				} else {
					logrus.Errorf("generateCoversForAuthor error: %v", err)
				}
			}
		}

		if author.CoverFileID == nil && ele.CoverFileID != nil {
			cID := ele.CoverFileID
			author.CoverFileID = cID
			err := db.Model(&author).Update("cover_file_id", cID).Error
			if err != nil {
				logrus.Errorf("generateCoversForAuthor error: %v", err)
			}
		}
	}
}
