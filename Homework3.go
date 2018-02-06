package loggerlib

import (
	"time"
)

type Data_t struct {
	Logged_time time.Time
	Str         string
}

type Log_t struct {
	item Data_t
	next *Log_t
}

var headptr *Log_t
var tailptr *Log_t

func Addmsg(data Data_t) (int, error) {

}

func Clearlog() {

}

func Getlog() (string, error) {

}

func Savelog(filename string) error {

}
