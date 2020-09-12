package main
import (
	"./unicast"
	"fmt"
	"strings"
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
	var input unicast.UserInput
	input.Destination = destination
	input.Message = message
}