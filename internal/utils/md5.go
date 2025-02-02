package utils

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"strings"
)

func CalculateFileMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hasher := md5.New()
	buffer := make([]byte, 1024*1024)
	for {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				break
			}
			return "", err
		}
		hasher.Write(buffer[:n])
	}

	hashBytes := hasher.Sum(nil)
	md5Hash := hex.EncodeToString(hashBytes)
	return md5Hash, nil
}

func CalculateStringMD5(s string) (string, error) {
	r := strings.NewReader(s)

	hasher := md5.New()
	_, err := io.Copy(hasher, r)
	if err != nil {
		return "", err
	}

	hashBytes := hasher.Sum(nil)
	md5Hash := hex.EncodeToString(hashBytes)
	return md5Hash, nil
}
