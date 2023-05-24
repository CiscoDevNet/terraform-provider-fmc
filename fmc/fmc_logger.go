package fmc

import (
	"fmt"
	"log"
	"os"
)

type logger struct {
	Logger          *log.Logger
	urlLogger       *log.Logger
	requestsLogger  *log.Logger
	responsesLogger *log.Logger
}

var Log *logger

func init() {
	var err error
	Log, err = newLogger("", log.Ldate|log.Ltime)
	if err != nil {
		log.Fatal(err)
	}
}

func newLogger(prefix string, flag int) (*logger, error) {
	if LogEnabled() {
		logDir := "outputs"
		err := os.MkdirAll(logDir, os.ModePerm)
		if err != nil {
			return nil, err
		}

		reqResLogFileName := fmt.Sprintf("%s/%s", logDir, "reqres")
		reqResFile, err := os.OpenFile(reqResLogFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return nil, err
		}

		return &logger{
			Logger:          log.New(reqResFile, prefix, flag),
			urlLogger:       log.New(reqResFile, prefix, flag),
			requestsLogger:  log.New(reqResFile, prefix, flag),
			responsesLogger: log.New(reqResFile, prefix, flag),
		}, nil
	}
	return nil, nil
}

func LogEnabled() bool {
	return os.Getenv("LOG") == "true"
}

func (l *logger) debug(v ...interface{}) {
	if LogEnabled() {
		if (v[len(v)-1]).(string) == "request" {
			l.requestsLogger.SetPrefix("[DEBUG REQ] ")
			l.requestsLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "response" {
			l.responsesLogger.SetPrefix("[DEBUG RES] ")
			l.responsesLogger.Println(v...)
			l.responsesLogger.SetPrefix("")
			l.responsesLogger.SetFlags(0)
			l.responsesLogger.Println()
		} else if (v[len(v)-1]).(string) == "url" {
			l.urlLogger.SetPrefix("[DEBUG URL] ")
			l.urlLogger.Println(v...)
		}
	}
}

func (l *logger) info(v ...interface{}) {
	if LogEnabled() {
		if (v[len(v)-1]).(string) == "request" {
			l.requestsLogger.SetPrefix("[INFO REQ] ")
			l.requestsLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "response" {
			l.responsesLogger.SetPrefix("[INFO RES] ")
			l.responsesLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "url" {
			l.urlLogger.SetPrefix("[INFO URL] ")
			l.urlLogger.Println(v...)
		}
	}
}

func (l *logger) warn(v ...interface{}) {
	if LogEnabled() {
		if (v[len(v)-1]).(string) == "request" {
			l.requestsLogger.SetPrefix("[WARN REQ] ")
			l.requestsLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "response" {
			l.responsesLogger.SetPrefix("[WARN RES] ")
			l.responsesLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "url" {
			l.urlLogger.SetPrefix("[WARN URL] ")
			l.urlLogger.Println(v...)
		}
	}
}

func (l *logger) error(v ...interface{}) {
	if LogEnabled() {
		if (v[len(v)-1]).(string) == "request" {
			l.requestsLogger.SetPrefix("[ERROR REQ] ")
			l.requestsLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "response" {
			l.responsesLogger.SetPrefix("[ERROR RES] ")
			l.responsesLogger.Println(v...)
		} else if (v[len(v)-1]).(string) == "url" {
			l.urlLogger.SetPrefix("[ERROR URL] ")
			l.urlLogger.Println(v...)
		}
	}
}

func (l *logger) line() {
	if LogEnabled() {
		l.Logger.SetFlags(0)
		l.Logger.Print("================================================================================================================\n")
	}
}
