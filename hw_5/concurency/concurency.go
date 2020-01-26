package concurency

import (
	"fmt"
	"sync"
)

// Task is
type Task func() error

// ErrorCounter is concurrent safe error counter
type ErrorCounter struct {
	sync.Mutex
	Count int
	limit int
}

// Run launches tasks in N goroutines. N and M must are more or equal 1. Tasks works while there is not M errors
func Run(tasks []Task, N int, M int) error {
	if N < 1 || M < 1 {
		return fmt.Errorf("parameters N and M must are > 1")
	}

	taskCh := make(chan Task, len(tasks))
	wg := sync.WaitGroup{}
	wg.Add(N)
	errorCounter := ErrorCounter{limit: M}

	for concurTask := 0; concurTask < N; concurTask++ {
		go func() {
			defer wg.Done()
			worker(taskCh, &errorCounter)
		}()
	}

	for _, task := range tasks {
		taskCh <- task
	}
	close(taskCh)
	wg.Wait()

	return nil
}

// worker getting task from the channel taskCh and launching it.
func worker(taskCh <-chan Task, errorCounter *ErrorCounter) {

	for task := range taskCh {
		err := task()

		if err != nil {
			if err = incrementErrors(errorCounter); err != nil {
				return
			}
		}

	}
}

// incrementErrors counting errors and compare with permissible errorLimit
func incrementErrors(errorCounter *ErrorCounter) error {
	errorCounter.Lock()
	defer errorCounter.Unlock()

	errorCounter.Count++
	if errorCounter.Count >= errorCounter.limit {
		return fmt.Errorf("error count is M")
	}

	return nil
}
