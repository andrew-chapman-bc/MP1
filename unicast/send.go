package unicast

import (
	"fmt"
	"sync"
	"time"
)

func unicast_send(destination, message string) {

}





//THIS IS FOR sync.waitgroup

func sleepFun(sec time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(sec * time.Second)
	fmt.Println("goroutine exit")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(2)
	go sleepFun(1, &wg)
	go sleepFun(3, &wg)
	wg.Wait()
	fmt.Println("Main goroutine exit")

}