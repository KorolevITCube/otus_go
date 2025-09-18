package main

import (
	"errors"
	"fmt"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(n, m int) error {
	fmt.Println("__________Start task worker___________")
	wg := sync.WaitGroup{}
	errorCh := make(chan error)
	stopCh := make(chan struct{})
	tasksCh := make(chan int, 10)
	doneCh := make(chan struct{})

	for t := range 10 {
		tasksCh <- t
	}
	close(tasksCh)

	for i := range n {
		wg.Add(1)
		fmt.Printf("Create goroutine %d\n", i)
		go func() {
			fmt.Printf("Started goroutine %d\n", i)
			defer func() {
				fmt.Printf("Deferred wg.Done from goroutine %d \n", i)
				wg.Done()
			}()
			worker(i, stopCh, errorCh, tasksCh)
			fmt.Printf("Close goroutine %d\n", i)
		}()
	}
	// вспомогательная горутина, которая блокируется и ожидает завершения всех воркеров
	go func() {
		fmt.Println("System goroutine started")
		wg.Wait()
		fmt.Println("System goroutine got wg.wait")
		close(doneCh)
	}()

	close(stopCh)

	errCount := 0
	for {
		select {
		case <-errorCh:
			errCount++
			if errCount >= m {
				fmt.Println("Got max errors, close workers")
				close(stopCh)
				fmt.Println("Blocks to wait all goroutines")
				<-doneCh
				fmt.Println("Exit")
				return ErrErrorsLimitExceeded
			}
		case <-doneCh:
			return nil
		}
	}
}

func worker(id int, stopCh <-chan struct{}, errorCh chan<- error, t <-chan int) {
	for {
		select {
		case <-stopCh:
			fmt.Printf("Worker %d stoped\n", id)
			return
		case i, ok := <-t:
			if !ok {
				return
			}
			fmt.Printf("Worker %d got new task %d\n", id, i)
		default:
			fmt.Printf("Worker %d case default", id)
		}
	}
}

func main() {
	Run(1, 10)
}
