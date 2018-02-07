/******************************************************************************
/  filename: loggerlib.go
/
/ description: Implements a logger library in go
/
/  author: Miller, Nathaniel
/  login id: SP_18_CPS444_02
/
/  class: CPS 444
/  instructor: Perugini
/  assignment: Homework #3
/
/  assigned: Febuary 5, 2018
/  due: Febuary 9, 2018
/
/****************************************************************************/
package loggerlib

import (
	"errors"
	"io/ioutil"
	"runtime"
	"strings"
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
			result += formattime(curr.item.Logged_time) + "\n" + curr.item.Str + "\n\n"
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
		err = ioutil.WriteFile(filename, []byte(str), 0774)
		if err != nil {
			return err
		}
	} else {
		return errors.New("No log available")
	}
	return nil
}

func formattime(t time.Time) string {
	var timestring string
	months := make(map[string]string)

	months["Jan"] = "01"
	months["Feb"] = "02"
	months["Mar"] = "03"
	months["Apr"] = "04"
	months["May"] = "05"
	months["Jun"] = "06"
	months["Jul"] = "07"
	months["Aug"] = "08"
	months["Sep"] = "09"
	months["Oct"] = "10"
	months["Nov"] = "11"
	months["Dec"] = "12"

	const layout = "Jan 2 2006 15:04:05"
	timeslice := strings.Split(t.Format(layout), " ")
	timestring = months[timeslice[0]] + "/"
	if len(timeslice[1]) == 1 {
		timestring += "0"
	}
	timestring += timeslice[1] + "/" + timeslice[2] + " " + timeslice[3]
	return timestring
}
