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
	"errors"
)

type UserInput struct {
	Destination string
	Message     string
	Source 		string
}

//keeps track of delay bounds from config
type Delay struct {
	min_delay string
	max_delay string
}

type Connection struct {
	ip 		string
	port 	string
	source 	string
}


func ScanConfig(userInput UserInput) Connection {

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
	for {
		success := scanner.Scan()
		if success == false {
			err = scanner.Err()
			if err == nil {
				fmt.Println("Scan completed and reached EOF")
				break
			} else {
				log.Fatal(err)
				break
			}
		}
		configArray := strings.Fields(scanner.Text())
		if configArray[0] == destination {
			connection.ip = configArray[1]
			connection.port = configArray[2]
			connection.source = userInput.Source
		}
		counter++
	}
	return connection
}


func connectToTCPServer(connect string) (net.Conn, error) {
	// Dial in to the TCP Server, return the connection to it
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return nil, nil
	}

	return c, errors.New("Error connecting to server")
}

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
	delayStruct.min_delay = delays[0]
	delayStruct.max_delay = delays[1]
	return delayStruct, errors.New("Error: Cannot fetch delay params")
} 

func generateDelay (delay Delay) {
	rand.Seed(time.Now().UnixNano())
	min, _ := strconv.Atoi(delay.min_delay)
	max, _ := strconv.Atoi(delay.max_delay)
	delayTime := rand.Intn(max - min + 1) + min
	//TODO: Decide if we want this here or in other file
	time.Sleep(time.Duration(delayTime))
}


func SendMessage( messageParams UserInput, connection Connection ) {
	connectionString := connection.ip + ":" + connection.port
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

