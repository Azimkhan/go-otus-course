package week2

import (
	"fmt"
	"testing"

	"gotest.tools/assert"
)

func TestUnpackString(t *testing.T) {
	testData := []struct {
		arg            string
		expectedResult string
		expectError    bool
	}{
		{"a4bc2d5e", "aaaabccddddde", false},
		{"abcd", "abcd", false},
		{"45", "", true},
		{`qwe\4\5`, `qwe45`, false},
		{`qwe\45`, `qwe44444`, false},
		{`qwe\\5`, `qwe\\\\\`, false},
		{`привет5`, `приветтттт`, false},
		{`привет\5`, `привет5`, false},
	}
	var res string
	var err error

	for _, data := range testData {
		res, err = UnpackString(data.arg)
		fmt.Printf("Testing %v\n", data.arg)
		if data.expectError {
			assert.Equal(t, true, err != nil)
		} else {
			assert.Equal(t, data.expectedResult, res)
		}
	}

}
