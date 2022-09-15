package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	i := 0
	for b := 0; b < 10; b++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			i = i + 1
			fmt.Println(i)
		}()
	}
	wg.Wait()
}
