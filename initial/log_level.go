package initial

import (
	"strings"

	"github.com/Yosorable/gomic/internal/global"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func SetLogrusAndGinFromConfigLogLevel() {
	logLevel := strings.ToLower(strings.TrimSpace(global.CONFIG.LogLevel))

	lvl := logrus.DebugLevel

	if logLevel == "panic" {
		lvl = logrus.PanicLevel
	} else if logLevel == "fatal" {
		lvl = logrus.FatalLevel
	} else if logLevel == "error" {
		lvl = logrus.ErrorLevel
	} else if logLevel == "info" {
		lvl = logrus.InfoLevel
	} else if logLevel == "debug" {
		lvl = logrus.DebugLevel
	}

	logrus.SetLevel(lvl)

	if lvl <= logrus.ErrorLevel {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}
}
