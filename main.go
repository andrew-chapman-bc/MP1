package main
import (
	"./unicast"
	"bufio"
	"fmt"
	"os"
	"sync"
)

func getInput() []string {
	fmt.Println("Enter input >> ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Print("this is input: " + input + "\n")
	// input: "send destination message"
	// input array: [send, destination, message]
	inputArray := strings.Fields(input)
	return inputArray

}

func parseInput() (unicast.UserInput, unicast.Connection) {
	inputArray := getInput()
	inputStruct := unicast.CreateUserInputStruct(inputArray[1], inputArray[2], os.Args[1])
	connection := unicast.ScanConfig(inputStruct)
	return inputStruct, connection
}

func unicast_send(inputStruct, connection) {
	defer wg.Done()
	unicast.SendMessage(inputStruct, connection)
}

func main() {
	var wg sync.WaitGroup
	inputStruct, connection := parseInput()
	wg.Add(2)

	go unicast_send(inputStruct, connection)
}


/* 
//	Throwing this here for now
// 
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
*/