package logs

import (
	"daltondiaz/async-jobs/conf"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	// Error log
	ErrorLog *log.Logger
	// Job log
	JobLog *log.Logger
	// Job name to be used in the future to possible separations of
	// log by job
	Job string
)

// Log File struct to be used like parameter to create new Log
type LogFile struct {
	// Type of log error, job
	TypeLog string
	// Job name
	JobName string
}

func init() {
	ErrorFile, err := CreateLog(LogFile{"error", ""})
	if err != nil {
		log.Fatal("error to create error log")
	}
	ErrorLog = log.New(ErrorFile, "ERROR ", log.Ldate|log.Ltime|log.Lshortfile)

	JobFile, err := CreateLog(LogFile{"job", Job})
	if err != nil {
		log.Fatal("error to create job log")
	}
	JobLog = log.New(JobFile, "JOB ", log.Ldate|log.Ltime|log.Lshortfile)
}

// Create a new log and return a os.File using the default directory
// and error if happend something
func CreateLog(logFile LogFile) (*os.File, error) {
	conf.LoadEnv()
	logsDir := os.Getenv("LOGS_DIR")
	nameFile := fmt.Sprintf("%s/%s_%d.log", logsDir, logFile.TypeLog, time.Now().Unix())
	return os.OpenFile(nameFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
}
