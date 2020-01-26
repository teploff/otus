package concurency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// TestCase checking invalid value of N
func TestIncorrectNValueInput(t *testing.T) {

	tasks := []Task{func() error {
		return nil
	}}
	threadCount := 0
	errorLimit := 1
	err := Run(tasks, threadCount, errorLimit)

	assert.Error(t, err)
}

// TestCase checking invalid value of M
func TestIncorrectMValueInput(t *testing.T) {

	tasks := []Task{func() error {
		return nil
	}}
	threadCount := 1
	errorLimit := 0
	err := Run(tasks, threadCount, errorLimit)

	assert.Error(t, err)
}

// TestCase checking that there is no >= M errors
func TestAllTasksFinishWithoutErrors(t *testing.T) {

	tasks := []Task{func() error {
		return nil
	}, func() error {
		return nil
	}, func() error {
		return nil
	}, func() error {
		return nil
	}, func() error {
		return nil
	}}
	threadCount := 5
	errorLimit := 5
	err := Run(tasks, threadCount, errorLimit)

	assert.Nil(t, err)
}

// TestCase checking that launching not more N + M tasks
func TestAllTasksFinishWithErrors(t *testing.T) {

	workTime := time.Second
	tasks := []Task{func() error {
		time.Sleep(workTime)
		return fmt.Errorf("error from task 1")
	}, func() error {
		time.Sleep(workTime)
		return fmt.Errorf("error from task 2")
	}, func() error {
		time.Sleep(workTime)
		return fmt.Errorf("error from task 3")
	}, func() error {
		time.Sleep(workTime)
		return fmt.Errorf("error from task 4")
	}, func() error {
		time.Sleep(workTime)
		return fmt.Errorf("error from task 5")
	}}
	threadCount := 3
	errorLimit := 4
	startTime := time.Now()
	err := Run(tasks, threadCount, errorLimit)
	endTime := time.Now()

	assert.Nil(t, err)
	assert.Equal(t, int(workTime.Seconds())*2, endTime.Second()-startTime.Second())
}

// TestCase checking that all tasks are launched concurrency (in different coroutines)
func TestAllTasksLaunchParallel(t *testing.T) {

	workTime := time.Second
	tasks := []Task{func() error {
		time.Sleep(workTime)
		return nil
	}, func() error {
		time.Sleep(workTime)
		return nil
	}, func() error {
		time.Sleep(workTime)
		return nil
	}, func() error {
		time.Sleep(workTime)
		return nil
	}}
	threadCount := 4
	errorLimit := 1
	startTime := time.Now()
	err := Run(tasks, threadCount, errorLimit)
	endTime := time.Now()

	assert.Nil(t, err)
	assert.Equal(t, int(workTime.Seconds()), endTime.Second()-startTime.Second())
}
