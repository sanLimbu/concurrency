package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	now := time.Now()
	userID := 10
	respch := make(chan string, 100)
	wg := &sync.WaitGroup{}
	wg.Add(3)
	go fetchUserData(userID, respch, wg)
	go fetchUserData2(userID, respch, wg)
	go fetchUserData3(userID, respch, wg)

	go func() {
		wg.Wait()
		close(respch)
	}()

	for resp := range respch {
		fmt.Println(resp)
	}

	fmt.Println(time.Since(now))

}

func fetchUserData(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	respch <- "user data"
	wg.Done()
}

func fetchUserData2(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(10 * time.Millisecond)
	respch <- "user data2"
	wg.Done()
}

func fetchUserData3(userID int, respch chan string, wg *sync.WaitGroup) {
	time.Sleep(140 * time.Millisecond)
	respch <- "user data3"
	wg.Done()
}
