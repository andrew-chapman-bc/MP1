package main
import (
	"./unicast"
	"bufio"
	"fmt"
	"os"
	"sync"
	"strings"
	"github.com/akamensky/argparse"
	"strconv"
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
func parseInput(source *string) (unicast.UserInput, unicast.Connection) {
	inputArray := getInput()
	inputStruct := unicast.CreateUserInputStruct(inputArray[1], inputArray[2], *source)
	connection := unicast.ScanConfigForClient(inputStruct)
	return inputStruct, connection
}

/*
	@function: openTCPServerConnections
	@description: Opens all of the ports defined in the config file using ScanConfigForServer() to get an array of ports 
					and ConnectToTCPClient() to open them
	@exported: False
	@params: {WaitGroup}
	@returns: N/A
*/
func openTCPServerConnections( source *string, wg *sync.WaitGroup) {
	openPort, err := unicast.ScanConfigForServer(*source)
	if openPort == "" {
		fmt.Println(err)
	}
	unicast.ConnectToTCPClient(openPort)
}

/*
	@function: unicast_send
	@description: function used as a goroutine to call SendMessage() to pass data from client to server, utilizes waitgroup
	@exported: False
	@params: {UserInput}, {Connection}, {WaitGroup}
	@returns: N/A
*/
func unicastSend(inputStruct unicast.UserInput, connection unicast.Connection, wg *sync.WaitGroup) {
	defer wg.Done()
	unicast.SendMessage(inputStruct, connection)
}

func main() {
	parser := argparse.NewParser("", "Concurrent TCP Channels")
	i := parser.Int("i", "int", &argparse.Options{Required: true, Help: "Source destination/identifiers"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
	}
	s := strconv.Itoa(*i)

	var wg sync.WaitGroup
	wg.Add(1)
	go openTCPServerConnections(&s, &wg)
	inputStruct, connection := parseInput(&s)
	wg.Add(1)
	go unicastSend(inputStruct, connection, &wg)
	wg.Wait()
}