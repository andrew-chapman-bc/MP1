package unicast

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func createStruct(destination, message string) UserInput {
	var input UserInput
	input.Destination = destination
	input.Message = message
	return input
}

func connectToTCPClient(PORT string) net.Conn {
	// listen/connect to the tcp client
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer l.Close()

	c, err := l.Accept()
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return c
}

func ReceiveMessage(source string) UserInput{
	if source == "" {
		fmt.Println("Port number not provided... Exiting")
		return createStruct("","")
	}

	Connect := connectToTCPClient(source)
	if Connect == nil {
		fmt.Println("An error has occurred... Exiting")
		return createStruct("","")
	}

	netData, err := bufio.NewReader(Connect).ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return createStruct("","")
	}

	netArray := strings.Fields(netData)
	dataStruct := createStruct(netArray[1], netArray[2])

	//letting client know server is exiting
	//Connect.Write([]byte("Exiting TCP Server..."))
	return dataStruct

}