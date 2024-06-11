package logs

import (
	"daltondiaz/async-jobs/conf"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	JobLog   *log.Logger
	Job      string
)

type LogFile struct {
	TypeLog string
	JobName string
}

func init() {
	logFile, err := createLog(LogFile{"info", ""})
	if err != nil {
		log.Fatal("error to create info log", err)
	}
	InfoLog = log.New(logFile, "INFO ", log.Ldate|log.Ltime|log.Lshortfile)

	ErrorFile, err := createLog(LogFile{"error", ""})
	if err != nil {
		log.Fatal("error to create debug log")
	}
	ErrorLog = log.New(ErrorFile, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	JobFile, err := createLog(LogFile{"job", Job})
	if err != nil {
		log.Fatal("error to create debug log")
	}
	JobLog = log.New(JobFile, "JOB ", log.Ldate|log.Ltime|log.Lshortfile)
}

func createLog(logFile LogFile) (*os.File, error) {
	conf.LoadEnv()
	logsDir := os.Getenv("LOGS_DIR")
	nameFile := fmt.Sprintf("%s/%s_%d.log", logsDir, logFile.TypeLog, time.Now().Unix())
	return os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}
