package lesson_golang_goroutines_pzn

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCPU)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine :", totalGoroutine)

	group.Wait()
}

func TestChangeThreadNumber(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(1 * time.Second)
			group.Done()
		}()
	}
	totalCPU := runtime.NumCPU()
	fmt.Println("Total CPU :", totalCPU)

	runtime.GOMAXPROCS(20)
	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Thread :", totalThread)

	totalGoroutine := runtime.NumGoroutine()
	fmt.Println("Total goroutine :", totalGoroutine)

	group.Wait()
}
