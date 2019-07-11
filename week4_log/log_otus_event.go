package week4_log

import "io"

type HwAccepted struct {
	Id    int
	Grade int
}

func (e *HwAccepted) Log() string {
	panic("implement me")
}

type HwSubmitted struct {
	Id      int
	Code    string
	Comment string
}

func (e *HwSubmitted) Log() string {
	panic("implement me")
}

type OtusEvent interface {
	Log() string
}

func LogOtusEvent(e OtusEvent, w io.Writer) {

}
