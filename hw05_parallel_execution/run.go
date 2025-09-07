package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, n, m int) error {
	wg := sync.WaitGroup{}
	errorCh := make(chan error, len(tasks))
	stopCh := make(chan struct{})
	tasksCh := make(chan Task, len(tasks))
	doneCh := make(chan struct{})

	for _, t := range tasks {
		tasksCh <- t
	}
	close(tasksCh)

	for range n {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			worker(stopCh, errorCh, tasksCh)
		}()
	}
	// вспомогательная горутина, которая блокируется и ожидает завершения всех воркеров
	go func() {
		wg.Wait()
		close(doneCh)
	}()

	errCount := 0
	for {
		select {
		case <-errorCh:
			errCount++
			if errCount >= m {
				close(stopCh)
				<-doneCh
				return ErrErrorsLimitExceeded
			}
		case <-doneCh:
			return nil
		}
	}
}

func worker(stopCh <-chan struct{}, errorCh chan<- error, t <-chan Task) {
	// вариант с range и селектом внутри, тк селект из двух каналов
	// в цикле не дает гарантии из какого канала вычитает данные, что приводит к
	// взятию в работу лишних задач
	// Можно реализовать через context, вероятно, это будет более красивое решение
	for task := range t {
		select {
		case <-stopCh:
			return
		default:
			if err := task(); err != nil {
				errorCh <- err
			}
		}
	}
}
