package main

import (
	"fmt"
	"sync"
)

func level2(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	ch <- "level-2"
}

func level1(ch chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var wg1 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg1.Add(1)
		go level2(ch, &wg1)
	}
	ch <- "level-1"
	wg1.Wait()
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go level1(ch, &wg)
	}
	done := make(chan bool)
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
		done <- true
		// for {
		// 	fmt.Println(<-ch)
		// }
	}()
	wg.Wait()
	close(ch)
	<-done

}
