package loggerlib

import (
	"errors"
	"runtime"
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
	log := Log_t{item: data, next: nil}
	if headptr != nil {
		tailptr.next = &log
		tailptr = &log
	} else {
		headptr = &log
		tailptr = &log
	}
	return 0, nil
}

func Clearlog() {
	headptr = nil
	tailptr = nil
	runtime.GC()
}

func Getlog() (string, error) {
	result := ""
	if headptr != nil {
		curr := headptr
		for curr != nil {
			result += curr.item.Logged_time.String() + "\n" + curr.item.Str + "\n\n"
			curr = curr.next
		}
	} else {
		return "", errors.New("No log available")
	}
	return result, nil
}

func Savelog(filename string) error {
	if headptr != nil {
		str, err := Getlog()
		if err != nil {
			return err
		}

	} else {
		return errors.New("No log available")
	}
	return nil
}
