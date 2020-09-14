package unicast

import (
	"fmt"
	"net"
	"errors"
	"bufio"
	"strings"
	"time"
)

func CreateUserInputStruct(destination, message, source string) UserInput {
	var input UserInput
	input.Destination = destination
	input.Message = message
	input.Source = source
	return input
}

func handleConnection(c net.Conn) {
	for {
		// string message, string source
		netData, err := bufio.NewReader(c).ReadString('\n')
        if err != nil {
            fmt.Println(err)
            return
		}
		netArray := strings.Fields(netData)
		timeOfReceive := time.Now().Format("02 Jan 06 15:04 MST")
		fmt.Println("Received " + netArray[0] + " from process " + netArray[1] + "system time is: " + timeOfReceive)
	}
}


func connectToTCPClient(PORT string) {
	// listen/connect to the tcp client
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
	}
	defer l.Close()
	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			go handleConnection(c)
		}
	}
}


// Deprecated?
// func ReceiveMessage( messageParams UserInput, connection Connection) {
// 	source := connection.port
// 	if source == "" {
// 		fmt.Println("Port number not provided... Exiting")
// 	}
// 	sender := connection.source
// 	if sender == "" {
// 		fmt.Println("Error: No Sender")
// 	}

// 	Connect, err := connectToTCPClient(source)
// 	if err != nil {
// 		fmt.Println("Error: ", err)
		
// 	}

// 	netData, err := bufio.NewReader(Connect).ReadString('\n')
// 	if err != nil {
// 		fmt.Println("Error writing data: ", err)
// 	}

// 	netArray := strings.Fields(netData)
// 	fmt.Println(netArray)


// }