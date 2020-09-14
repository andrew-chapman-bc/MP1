package unicast

import (
	"fmt"
	"net"
	"bufio"
	"strings"
	"time"
	"os"
	"log"
	"errors"
)

/*
	@function: ScanConfigForServer
	@description: Scans the config file for all of the ports that will be used to open concurrent TCP Servers
	@exported: True
	@params: N/A
	@returns: []string
*/
func ScanConfigForServer(source string) (string, error) {
	config, err := os.Open("config.txt")
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)
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
		// don't check the first line
		if (counter != 0) {
			configArray := strings.Fields(scanner.Text())
			port := configArray[2]
			if (configArray[0] == source) {
				return port, nil
			}
		}
		counter++
	}
	return "", errors.New("Cannot find port")
}


/*
	@function: CreateUserInputStruct
	@description: Uses a destination, message and source string to construct a UserInput struct to send and receive the message across the server/client
	@exported: True
	@params: string, string, string
	@returns: {UserInput}
*/
func CreateUserInputStruct(destination, message, source string) UserInput {
	var input UserInput
	input.Destination = destination
	input.Message = message
	input.Source = source
	return input
}


/*
	@function: handleConnection
	@description: handles connections to the concurrent TCP client and receives messages that are sent over through a goroutine in connectToTCPClient
	@exported: False
	@params: net.Conn
	@returns: N/A
*/
func handleConnection(c net.Conn) {
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		fmt.Println("NETDATA: ", netData)
        if err != nil {
            fmt.Println(err)
            return
		}
		netArray := strings.Fields(netData)
		timeOfReceive := time.Now().Format("02 Jan 06 15:04 MST")
		fmt.Println("Received " + netArray[0] + " from process " + netArray[1] + "system time is: " + timeOfReceive)
		// c.Write([]byte("Received " + netArray[0] + " from process " + netArray[1] + "system time is: " + timeOfReceive))
		if netArray[0] == "STOP" {
			break
		}
	}
	c.Close()
}

/*
	@function: connectToTCPClient
	@description: Opens a concurrent TCP Server and calls the net.Listen function to connect to the TCP Client
	@exported: True
	@params: string
	@returns: N/A
*/
func ConnectToTCPClient(PORT string) {
	// listen/connect to the tcp client
	l, err := net.Listen("tcp4", ":" + PORT)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("port open: ", PORT)
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
		}
		go handleConnection(c)
	}
}