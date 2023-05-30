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
var allLog bool
var reqResLog bool
var userLog bool

func init() {
	var err error
	Log, err = newLogger("", log.Ldate|log.Ltime)
	if err != nil {
		log.Fatal(err)
	}
}

func newLogger(prefix string, flag int) (*logger, error) {
	allLog = allLogEnabled()
	reqResLog = reqResLogEnabled()
	userLog = userLogEnabled()
	if allLog || userLog || reqResLog {
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

func reqResLogEnabled() bool {
	return os.Getenv("LOGS") == "REQ_RES"
}

func userLogEnabled() bool {
	return os.Getenv("LOGS") == "USER"
}

func allLogEnabled() bool {
	return os.Getenv("LOGS") == "ALL"
}

func (l *logger) debug(v ...interface{}) {
	if allLog || reqResLog || userLog {
		switch v[len(v)-1] {
		case "request":
			if allLog || reqResLog {
				l.requestsLogger.SetPrefix("[DEBUG REQ] ")
				l.requestsLogger.Println(v...)
			}
		case "response":
			if allLog || reqResLog {
				l.responsesLogger.SetPrefix("[DEBUG RES] ")
				l.responsesLogger.Println(v...)
				l.responsesLogger.SetPrefix("")
				l.responsesLogger.SetFlags(0)
				l.responsesLogger.Println()
			}
		default:
			if allLog || userLog {
				l.urlLogger.SetPrefix("[DEBUG USER] ")
				l.urlLogger.Println(v...)
				l.Logger.SetFlags(0)
				l.Logger.Print("\n================================================================================================================")
			}
		}
	}
}

func (l *logger) info(v ...interface{}) {
	if allLog || reqResLog || userLog {
		switch v[len(v)-1] {
		case "request":
			if allLog || reqResLog {
				l.requestsLogger.SetPrefix("[INFO REQ] ")
				l.requestsLogger.Println(v...)
			}
		case "response":
			if allLog || reqResLog {
				l.responsesLogger.SetPrefix("[INFO RES] ")
				l.responsesLogger.Println(v...)
				l.responsesLogger.SetPrefix("")
				l.responsesLogger.SetFlags(0)
				l.responsesLogger.Println()
			}
		default:
			if allLog || userLog {
				l.urlLogger.SetPrefix("[INFO USER] ")
				l.urlLogger.Println(v...)
				l.Logger.SetFlags(0)
				l.Logger.Print("\n================================================================================================================")
			}
		}
	}
}

func (l *logger) warn(v ...interface{}) {
	if allLog || reqResLog || userLog {
		switch v[len(v)-1] {
		case "request":
			if allLog || reqResLog {
				l.requestsLogger.SetPrefix("[WARN REQ] ")
				l.requestsLogger.Println(v...)
			}
		case "response":
			if allLog || reqResLog {
				l.responsesLogger.SetPrefix("[WARN RES] ")
				l.responsesLogger.Println(v...)
				l.responsesLogger.SetPrefix("")
				l.responsesLogger.SetFlags(0)
				l.responsesLogger.Println()
			}
		default:
			if allLog || userLog {
				l.urlLogger.SetPrefix("[WARN USER] ")
				l.urlLogger.Println(v...)
				l.Logger.SetFlags(0)
				l.Logger.Print("\n================================================================================================================")
			}
		}
	}
}

func (l *logger) error(v ...interface{}) {
	if allLog || reqResLog || userLog {
		switch v[len(v)-1] {
		case "request":
			if allLog || reqResLog {
				l.requestsLogger.SetPrefix("[ERROR REQ] ")
				l.requestsLogger.Println(v...)
			}
		case "response":
			if allLog || reqResLog {
				l.responsesLogger.SetPrefix("[ERROR RES] ")
				l.responsesLogger.Println(v...)
				l.responsesLogger.SetPrefix("")
				l.responsesLogger.SetFlags(0)
				l.responsesLogger.Println()
			}
		default:
			if allLog || userLog {
				l.urlLogger.SetPrefix("[ERROR USER] ")
				l.urlLogger.Println(v...)
				l.Logger.SetFlags(0)
				l.Logger.Print("\n================================================================================================================")
			}
		}
	}
}

func (l *logger) line() {
	if allLog || reqResLog {
		l.Logger.SetFlags(0)
		l.Logger.Print("================================================================================================================\n")
	}
}
