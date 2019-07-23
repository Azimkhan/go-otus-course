package week4_parallel_execution

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
	"time"
)

func makeTasks(n int, numErrors int) []func() error {
	out := make([]func() error, n)

	for i := 0; i < n; i++ {
		out[i] = func(i0 int) func() error {
			return func() error {
				ms := 5 + (i0%10)*10
				fmt.Printf("Task #%d: Sleeping for %d milliseconds\n", i0, ms)
				time.Sleep(time.Duration(ms) * time.Millisecond)
				if i0 < numErrors {
					fmt.Printf("Task #%d: Error\n", i0)
					return fmt.Errorf("error")
				}
				return nil
			}
		}(i)
	}
	return out

}

func Test_execute(t *testing.T) {

	var tasks []func() error
	var err error

	tasks = makeTasks(100, 10)
	err = execute(tasks, 10, 9)
	assert.Error(t, err, "too many errors")

	err = execute(tasks, 10, 11)
	assert.NilError(t, err)
}
