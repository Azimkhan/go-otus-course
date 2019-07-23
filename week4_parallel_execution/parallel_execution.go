package week4_parallel_execution

import "fmt"

func execute(tasks []func() error, maxConcurrentTasks int, maxErrors int) error {
	resChan := make(chan error, maxConcurrentTasks)
	errors := 0

	taskLen := len(tasks)
	lastIndex := 0
	if taskLen < maxConcurrentTasks {
		maxConcurrentTasks = taskLen
	}

	executeTask := func(task func() error) {
		resChan <- task()
	}
	go func() {
		for i := 0; i < maxConcurrentTasks; i++ {
			lastIndex = i
			go executeTask(tasks[i])
		}
	}()

	for err := range resChan {
		if err != nil {
			errors += 1
		}
		if errors > maxErrors {
			return fmt.Errorf("too many errors")
		}
		if lastIndex >= taskLen-1 {
			close(resChan)
			break
		}
		lastIndex += 1
		go executeTask(tasks[lastIndex])
	}

	return nil
}
