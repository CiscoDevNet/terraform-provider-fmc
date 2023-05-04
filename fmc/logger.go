package fmc

import (
	"fmt"
	"log"
	"os"
)

type logger struct {
	logger *log.Logger
}

var Log *logger

func init() {
	var err error
	Log, err = newLogger("", log.Ldate|log.Ltime|log.Lshortfile)
	if err != nil {
		log.Fatal(err)
	}
}

func newLogger(prefix string, flag int) (*logger, error) {
	logDir := "logs"
	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		return nil, err
	}
	logFileName := fmt.Sprintf("%s/%s.logs", logDir, "finally")
	file, err := os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return &logger{
		logger: log.New(file, prefix, flag),
	}, nil
}

func (l *logger) Debug(v ...interface{}) {
	l.logger.SetPrefix("[DEBUG] ")
	l.logger.Println(v...)
}

func (l *logger) Info(v ...interface{}) {
	l.logger.SetPrefix("[INFO] ")
	l.logger.Println(v...)
}

func (l *logger) Warning(v ...interface{}) {
	l.logger.SetPrefix("[WARNING] ")
	l.logger.Println(v...)
}

func (l *logger) Error(v ...interface{}) {
	l.logger.SetPrefix("[ERROR] ")
	l.logger.Println(v...)
}
