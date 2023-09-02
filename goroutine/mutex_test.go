package goroutine

import (
	"fmt"
	"sync"
	"testing"
)

func TestDoublyNode(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	go worker(&wg)
	go worker(&wg)
	wg.Wait()
	fmt.Println(total.value)
	var wg2 sync.WaitGroup
	wg2.Add(2)
	go worker2(&wg2)
	go worker2(&wg2)
	wg2.Wait()
	fmt.Println(total2)
}
