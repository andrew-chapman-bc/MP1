package unicast

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

// UserInput holds user input such as message, destination and source
type UserInput struct {
	Destination string
	Message     string
	Source 		string
}

// Delay keeps track of delay bounds from config
type Delay struct {
	minDelay string
	maxDelay string
}
// Connection holds the ip/port and source
type Connection struct {
	ip 		string
	port 	string
	source 	string
}

/*
	@function: ScanConfigForClient
	@description: Scans the config file using the user input destination and retrieves the ip/port that will later be used to connect to the TCP server
	@exported: True
	@params: {userInput} 
	@returns: {Connection}
*/
func ScanConfigForClient(userInput UserInput) Connection {

	destination := userInput.Destination
	
	// Open up config file
	// TODO: create a variable which holds the destination of config file instead of hardcoding here
	config, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)
	var connection Connection
	counter := 0
	for {
		success := scanner.Scan()
		if success == false {
			err = scanner.Err()
			if err == nil {
				// fmt.Println("Scan completed and reached EOF")
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		if counter != 0 {
			configArray := strings.Fields(scanner.Text())
			if configArray[0] == destination {
				connection.ip = configArray[1]
				connection.port = configArray[2]
				connection.source = userInput.Source
			}
		}
		counter++
	}
	return connection
} 

/*
	@function: connectToTCPServer
	@description:	Connects to the TCP server with the ip/port obtained from config file as a parameter and 
					returns the connection to the server which will later be used to write to the server
	@exported: false
	@params: string 
	@returns: net.Conn, err
*/
func connectToTCPServer(connect string) (net.Conn, error) {
	// Dial in to the TCP Server, return the connection to it
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return c, err
} 


/*
	@function: getDelayParams
	@description: Scans the config file for the first line to get the delay parameters that will be used to simulate the network delay
	@exported: false
	@params: N/A 
	@returns: Delay, error
*/
func getDelayParams() (Delay, error) {
	config, err := os.Open("config.txt")
	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)
	success := scanner.Scan()
	if success == false {
		err = scanner.Err()
		if err == nil {
			fmt.Println("Scanned first line")
		} else {
			log.Fatal(err)
		}
	}
	delays := strings.Fields(scanner.Text())
	var delayStruct Delay
	delayStruct.minDelay = delays[0]
	delayStruct.maxDelay = delays[1]
	return delayStruct, err
} 

/*
	@function: generateDelay
	@description: Uses the delay parameters obtained from getDelayParams() to generate the delay that will be used in sendMessage function
	@exported: false
	@params: Delay
	@returns: N/A
*/
func generateDelay (delay Delay) {
	rand.Seed(time.Now().UnixNano())
	min, _ := strconv.Atoi(delay.minDelay)
	max, _ := strconv.Atoi(delay.maxDelay)
	delayTime := rand.Intn(max - min + 1) + min
	//TODO: Decide if we want this here or in other file
	time.Sleep(time.Duration(delayTime))
} 

/*
	@function: SendMessage
	@description: 	SendMessage sends the message from TCPClient to TCPServer by connecting to the server and 
					using the Fprintf function to send the message.  After it does this, it calls generateDelay to simulate a network delay
	@exported: True
	@params: {UserInput}, {Connection}
	@returns: N/A
*/
func SendMessage( messageParams UserInput, connection Connection ) {
	connectionString := connection.ip + ":" + connection.port
	fmt.Println(connectionString)
	c, err := connectToTCPServer(connectionString)
	if (err != nil) {
		fmt.Println("Network Error: ", err)
	}
	
	delay, err := getDelayParams()
	if (err != nil) {
		fmt.Println("Error: ", err)
	}
	
	// Sending the message to TCP Server
	fmt.Fprintf(c, messageParams.Message, messageParams.Source)
	timeOfSend := time.Now().Format("02 Jan 06 15:04 MST")
	fmt.Println("Sent message " + messageParams.Message + " to destination " + messageParams.Destination + " system time is: " + timeOfSend)
	
	// Generate Delay
	generateDelay(delay)
} 

