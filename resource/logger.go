package resource

import (
	"fmt"
	go_logger "github.com/phachon/go-logger"
)

var logger *go_logger.Logger

// InitLogger 初始化 logger
func InitLogger() {
	logger = go_logger.NewLogger()
	_ = logger.Detach("console")
	logFormat := "[%timestamp_format%] [%level_string%] %function%:%line% %body%"


	// file adapter config
	fileConfig := &go_logger.FileConfig{
		LevelFileName: map[int]string{
			logger.LoggerLevel("error"): fmt.Sprintf("%s/reviewer-crontab.log", "/Users/xixianbin/logs"),
			logger.LoggerLevel("info"):  fmt.Sprintf("%s/reviewer-crontab.log", "/Users/xixianbin/logs"),
			logger.LoggerLevel("debug"): fmt.Sprintf("%s/reviewer-crontab.log", "/Users/xixianbin/logs"),
		},
		MaxSize:    0,
		MaxLine:    1000000,
		DateSlice:  "y",
		JsonFormat: false,
		Format:     logFormat,
	}

	_ = logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
}

func GetLogger() *go_logger.Logger {
	return logger
}
