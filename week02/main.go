package main

import (
	"fmt"
	"sync"
)

type Calculator struct {
	Result int
	Mutex  sync.Mutex
}

// add
func (c *Calculator) Add(num int) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Result += num
}

// subtract
func (c *Calculator) Subtract(num int) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Result -= num
}

// multiply
func (c *Calculator) Multiply(num int) {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Result *= num
}

func main() {
	fmt.Println("Hello week02")
	var wg sync.WaitGroup

	counter := &Calculator{Result: 1}

	wg.Add(3)

	go func() {
		counter.Add(10)
		wg.Done()
	}()

	go func() {
		counter.Subtract(2)
		wg.Done()
	}()

	go func() {
		counter.Multiply(4)
		wg.Done()
	}()

	wg.Wait()
	fmt.Println("Final Value:", counter.Result)
}
