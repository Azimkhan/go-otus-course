package week4_log

import (
	"testing"
	"bytes"
	"gotest.tools/assert"
	"time"
)

var fakeNow = func() time.Time {
	return time.Date(2019,1,1,0,0,0,0, time.UTC)
}

func TestLogOtusEventSubmitted(t *testing.T) {
	setNowFunc(fakeNow)

	event := &HwSubmitted{
		Id:      3456,
		Code:    "unknown",
		Comment: "please take a look at my homework",
	}

	var buffer = &bytes.Buffer{}

	buffer.Reset()
	shouldBe := "2019-01-01 submitted 3456 \"please take a look at my homework\"\n"

		LogOtusEvent(event, buffer)
	assert.Equal(t, shouldBe, string(buffer.Bytes()))
}

func TestLogOtusEventAccepted(t *testing.T) {

	setNowFunc(fakeNow)
	event := &HwAccepted{
		Id:    3456,
		Grade: 4,
	}

	var buffer = &bytes.Buffer{}
	shouldBe := "2019-01-01 accepted 3456 4\n"
	LogOtusEvent(event, buffer)
	assert.Equal(t, shouldBe, string(buffer.Bytes()))

}
