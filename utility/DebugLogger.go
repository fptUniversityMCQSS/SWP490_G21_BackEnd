package utility

import (
	"github.com/natefinch/lumberjack"
	"log"
	"os"
	"strconv"
)

var DebugLog DebugLogger = DebugLogger{}

type DebugLogger struct {
	Logger           *log.Logger
	CurrentRequestID int64
}

func (logger *DebugLogger) Print(path string, end bool, id int64) int64 {
	var requestID int64
	endString := "START"
	if end {
		requestID = id
		endString = "FINISH"
	} else {
		requestID = logger.CurrentRequestID
		logger.CurrentRequestID++
	}
	DebugLog.Logger.Println("Request ID: " + strconv.FormatInt(requestID, 10) + " ; Request: " + path + " ; " + endString)
	return requestID
}

func init() {
	err := os.MkdirAll("log", os.ModePerm)
	if err != nil {
		log.Print(err)
	}
	logFile, err := os.OpenFile("log/debug.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		log.Print("Can not open or create log/debug.log: " + err.Error())
	}
	DebugLog.Logger = log.New(logFile, "", log.Ldate|log.Ltime)
	DebugLog.Logger.SetOutput(&lumberjack.Logger{
		Filename:   "log/debug.log",
		MaxSize:    1,  // megabytes after which new file is created
		MaxBackups: 3,  // number of backups
		MaxAge:     28, //days
	})
}
