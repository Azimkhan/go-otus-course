package week4_log

import (
	"io"
	"fmt"
	"time"
)

var nowFunc func() time.Time

func setNowFunc(now func() time.Time)  {
	if now == nil {
		nowFunc = time.Now
		return
	}
	nowFunc = now
}

func init (){
	setNowFunc(nil)
}

type HwAccepted struct {
	Id    int
	Grade int
}


type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (e *HwAccepted) Log() string {
	return fmt.Sprintf(`accepted %d %d`, e.Id, e.Grade)
}

func (e *HwSubmitted) Log() string {
	return fmt.Sprintf(`submitted %d "%s"`, e.Id, e.Comment)
}

type OtusEvent interface {
	Log() string
}


func LogOtusEvent(e OtusEvent, w io.Writer) {
	currentTime := nowFunc()
	w.Write([]byte(fmt.Sprintln(currentTime.Format("2006-01-02"), e.Log())))
}
