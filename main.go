package main
import (
	"./unicast"
	"fmt"
	"strings"
	"bufio"
	"os"
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

func createStruct(destination, message string) unicast.UserInput {
	var input unicast.UserInput
	input.Destination = destination
	input.Message = message
	return input
}

func main() {
	inputArray := getInput() 
	inputStruct := createStruct(inputArray[1], inputArray[2])
	unicast_send(inputStruct.Destination, inputStruct.Message)
}