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
	"sync"
	"time"
)

type UserInput struct {
	Destination string
	Message     string
}

type Delay struct {
	min_delay string
	max_delay string
}

func connectToTCPServer(connect string) net.Conn {
	// Dial in to the TCP Server, return the connection to it
	c, err := net.Dial("tcp", connect)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return c
}

func getDelay() Delay{
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
	return delayStruct
} 

func scanConfig(destination string) []string {
	
	getDelay()
	// Open up config file
	// TODO: create a variable which holds the destination of config file instead of hardcoding here
	config, err := os.Open("config.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(config)
	scanner.Split(bufio.ScanLines)
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
			return configArray
		}
	}
	return nil
}

func sendMessage(destination, message string) {
	
	configArray := scanConfig(destination)
	connection := configArray[2] 
	c := connectToTCPServer(connection)
	if c == nil {
		fmt.Println("Error connection to server... Exiting")
	}
	fmt.Fprintf(c, message)
	timeOfSend := time.Now().Format("02 Jan 06 15:04 MST")
	fmt.Println("Sent message " + message + " to destination " + destination + " system time is: " + timeOfSend)
	netDelay := getDelay()
	rand.Seed(time.Now().UnixNano())
	min, _ := strconv.Atoi(netDelay.min_delay)
	max, _ := strconv.Atoi(netDelay.max_delay)
	delayTime := rand.Intn(max - min + 1) + min
	//TODO: Decide if we want this here or in other file
	time.Sleep(time.Duration(delayTime))
}
