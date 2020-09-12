package main

// "./unicast/receive.go"
//	"./unicast/send.go"
//	"./unicast/tpcC.go"
//	"./unicast/tpcS.go"
import (
	"fmt"
	"strings"
	"unicast/receive.go"
	"unicast/send.go"
	"unicast/tpcC.go"
	"unicast/tpcS.go"
)

func getInput() {
	fmt.Println("Enter input >> ")
	var input string
	fmt.Scanln(&input)
	// input: "send destination message"
	// input array: [send, destination, message]
	inputArray := strings.Fields(input)
	createStruct(inputArray[1], inputArray[2])

}

func createStruct(destination, message string) {
	var input userInput
	input.destination = destination
	input.message = message
}
