package main
import (
	"./unicast"
	"bufio"
	"fmt"
	"os"
	"sync"
	"strings"
)

/*
	@function: getInput
	@description: gets the input entered through I/O and packages it into an array that will be used to create a {UserInput}
	@exported: False
	@params: N/A
	@returns: []string
*/
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

/*
	@function: parseInput
	@description: Parses the UserInput into a {UserInput} and calls ScanConfig() to parse the parameters of TCP connection into a {Connection}
	@exported: False
	@params: N/A
	@returns: {UserInput}, {Connection}
*/
func parseInput() (unicast.UserInput, unicast.Connection) {
	inputArray := getInput()
	inputStruct := unicast.CreateUserInputStruct(inputArray[1], inputArray[2], os.Args[1])
	connection := unicast.ScanConfigForClient(inputStruct)
	return inputStruct, connection
}

/*
	@function: openTCPServerConnections
	@description: Opens all of the ports defined in the config file using ScanConfigForServer() to get an array of ports 
					and ConnectToTCPClient() to open them
	@exported: False
	@params: N/A
	@returns: N/A
*/
func openTCPServerConnections() {
	openPortsArr := unicast.ScanConfigForServer()
	for _, port := range openPortsArr {
		unicast.ConnectToTCPClient(port)
	}
}

/*
	@function: unicast_send
	@description: function used as a goroutine to call SendMessage() to pass data from client to server, utilizes waitgroup
	@exported: False
	@params: {UserInput}, {Connection}, {WaitGroup}
	@returns: N/A
*/
func unicastSend(inputStruct unicast.UserInput, connection unicast.Connection, wg sync.WaitGroup) {
	defer wg.Done()
	unicast.SendMessage(inputStruct, connection)
}

func main() {
	var wg sync.WaitGroup
	openTCPServerConnections()
	inputStruct, connection := parseInput()
	wg.Add(2)

	go unicastSend(inputStruct, connection, wg)
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