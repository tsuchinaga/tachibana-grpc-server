package tachibana_grpc_server

import (
	"fmt"
	"log"
	"os"
)

type iLogger interface {
	error(v ...interface{})
	scheduler(v ...interface{})
	request(v ...interface{})
}

type logger struct {
	errorLog     *log.Logger
	schedulerLog *log.Logger
	requestLog   *log.Logger
}

func (l *logger) error(v ...interface{}) {
	if l.errorLog == nil {
		errorLog, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalln(err)
		}
		l.errorLog = log.New(errorLog, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	}

	_ = l.errorLog.Output(2, fmt.Sprintln(v...))
}

func (l *logger) scheduler(v ...interface{}) {
	if l.schedulerLog == nil {
		schedulerLog, err := os.OpenFile("logs/scheduler.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			if l.errorLog != nil {
				l.error(err, v)
			} else {
				log.Fatalln(err)
			}
		}
		l.schedulerLog = log.New(schedulerLog, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	}

	_ = l.schedulerLog.Output(2, fmt.Sprintln(v...))
}

func (l *logger) request(v ...interface{}) {
	if l.requestLog == nil {
		requestLog, err := os.OpenFile("logs/request.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			if l.errorLog != nil {
				l.error(err, v)
			} else {
				log.Fatalln(err)
			}
		}
		l.requestLog = log.New(requestLog, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	}

	_ = l.requestLog.Output(2, fmt.Sprintln(v...))
}
