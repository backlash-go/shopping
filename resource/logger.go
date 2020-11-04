package resource

import (
	"fmt"
	"log"

	"github.com/spf13/viper"

	go_logger "github.com/phachon/go-logger"
)

var (
	Logger         *go_logger.Logger
	ProgressLogger *go_logger.Logger
)

// InitLogger 初始化 logger
func InitLogger(env string) {
	InitProgressLoggerLogger()
	log.Println("init logger start")
	Logger = go_logger.NewLogger()
	logFormat := "[%timestamp_format%] [%level_string%] %function%:%line% %body%"

	if env == "dev" {
		_ = Logger.Detach("console")
		// console adapter config
		consoleConfig := &go_logger.ConsoleConfig{
			Color:      true,      // Does the text display the color
			JsonFormat: false,     // Whether or not formatted into a JSON string
			Format:     logFormat, // JsonFormat is false, logger message output to console format string
		}
		// add output to the console
		_ = Logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)
		log.Println("log to : console")
		return
	}
	logDir := viper.GetString("logDir")
	// file adapter config
	fileConfig := &go_logger.FileConfig{
		//Filename: "./reviewer-test.log", // The file name of the logger output, does not exist automatically
		// If you want to separate separate logs into files, configure LevelFileName parameters.
		LevelFileName: map[int]string{
			Logger.LoggerLevel("error"): fmt.Sprintf("%s/go-course.log", logDir), // The error level log is written to the error.log file.
			Logger.LoggerLevel("info"):  fmt.Sprintf("%s/go-course.log", logDir), // The info level log is written to the info.log file.
			Logger.LoggerLevel("debug"): fmt.Sprintf("%s/go-course.log", logDir), // The debug level log is written to the debug.log file.
		},
		MaxSize:    0,         // File maximum (KB), default 0 is not limited
		MaxLine:    0,         // The maximum number of lines in the file, the default 0 is not limited
		DateSlice:  "y",       // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
		JsonFormat: false,     // Whether the file data is written to JSON formatting
		Format:     logFormat, // JsonFormat is false, logger message written to file format string
	}
	// add output to the file
	_ = Logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	log.Println("log to dir : " + logDir)
}

// 本来记录到progress_log 表的用户行为数据，记录到log，再收集到大数据里
func InitProgressLoggerLogger() {
	log.Println(" InitProgressLoggerLogger start")
	ProgressLogger = go_logger.NewLogger()
	logFormat := "[%timestamp_format%] [%level_string%] %function%:%line% %body%"

	logDir := viper.GetString("logDir")
	// file adapter config
	fileConfig := &go_logger.FileConfig{
		//Filename: "./reviewer-test.log", // The file name of the logger output, does not exist automatically
		// If you want to separate separate logs into files, configure LevelFileName parameters.
		LevelFileName: map[int]string{
			Logger.LoggerLevel("error"): fmt.Sprintf("%s/go-course-progress.log", logDir), // The error level log is written to the error.log file.
			Logger.LoggerLevel("info"):  fmt.Sprintf("%s/go-course-progress.log", logDir), // The info level log is written to the info.log file.
			Logger.LoggerLevel("debug"): fmt.Sprintf("%s/go-course-progress.log", logDir), // The debug level log is written to the debug.log file.
		},
		MaxSize:    0,         // File maximum (KB), default 0 is not limited
		MaxLine:    0,         // The maximum number of lines in the file, the default 0 is not limited
		DateSlice:  "y",       // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
		JsonFormat: false,     // Whether the file data is written to JSON formatting
		Format:     logFormat, // JsonFormat is false, logger message written to file format string
	}
	// add output to the file
	_ = ProgressLogger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)
	log.Println("log to dir : " + logDir)
}
