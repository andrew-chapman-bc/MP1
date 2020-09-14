#MP1
--- 
#To Run

In terms of dependencies, the only one that you need to download is the argparse golang package.  You can download this by running this command into one of your open terminals:
```
$ go get -u -v github.com/akamensky/argparse
``` 


The config is currently set up to handle multiple processes

Simply repeat steps in new terminals to have as many processes as needed

Open one terminal and enter
```bash
go run main.go --int 1
``` 
Then open up a second terminal and enter
```bash
go run main.go --int 2
```
Then open up a third terminal and enter
```bash
go run main.go --int 3
```
Then open up a fourth terminal and enter
```bash
go run main.go --int 4
```
To send a message, go to any terminal and send

```bash
send 2 hello
```

(the 2 and hello are obviously interchangable)

Should output on this terminal with a different time
```bash 
Sent message hello to destination 2 system time is: 14 Sep 20 18:19 EDT
```
To see this message go back to terminal 2

The  following output should be listed with different time
```bash
Received hello from process 1 system time is: 14 Sep 20 18:19 EDT
```
If you want to send a message back to terminal 1, input
```bash
send 2 hi
```

Should output on this terminal with a different time
```bash 
Sent message hi to destination 1 system time is: 14 Sep 20 18:19 EDT
```

The  following output should be printed on terminal 1 with different time
```bash
Received hi from process 2 system time is: 14 Sep 20 18:19 EDT
```

---
#Structure and design
###TCP Server
Each process starts off by initalizing a concurrent TCP server
The user's commandline input and config file are used to generate the port number


We have three structures designed to make passing data around easier and more readable.

They are as follows: 

The UserInput struct is used to easily access and pass around the destination, message and source (who sent what) throughout the program without sending strings everywhere

```
type UserInput struct {
	Destination string
	Message     string
	Source 		string
}
```

The Delay structure is to have the delay parameters saved to make it more readable and accessable in our codebase
```
type Delay struct {
	minDelay string
	maxDelay string
}
```

The connection struct was created to more easily pass around the IP/Port to connect to on the client side.  Additionally we realized that we also needed the source many times throughout our codebase, so we also included that in the struct.
```
type Connection struct {
	ip 		string
	port 	string
	source 	string
}
```

###Config file
The config file has the following format in a txt file
-----------------------------------------------------------------------------------------------    
min_delay(ms) max_delay(ms)

ID1 IP1 port1

ID2 IP2 port2

ID3 IP3 port3

ID4 IP4 port4

.... .... .......
-----------------------------------------------------------------------------------------------
To read the config file, it is read line by line, and uses whitespace to differentiate between the different values

To add more processes, add a new line, with an ID, IP, and port number

For example:
-----------------------------------------------------------------------------------------------    
10 15
1 127.0.0.1 1234
2 127.0.0.1 4567
-----------------------------------------------------------------------------------------------

Goes to 

-----------------------------------------------------------------------------------------------    
10 15
1 127.0.0.1 1234
2 127.0.0.1 4567
3 127.0.0.1 8543
4 127.0.0.1 1432
-----------------------------------------------------------------------------------------------

To go from 2 to 4 processes

It is all on local host right now, so the IP is repeated for all of the processes

Since the program is basic, it is all that was necessary 

In a more complex program, we would use a different file format of the config, i.e JSON

###Input
The user inputs three strings, : 
1. "Send"
2. Destination 
3. Message

The program reads each string individually, the first one is disregarded, the second one is used to find the port to send message through, and message is sent through a thread to the TCP server

Since the message is so small, we choose not to serialize into another format such as GOB or JSON

###Processes
The processes can be found in the unicast directory

In tcpC.go is where the message is sent out to the server

It takes in the parsed input from the main, and connects to the proper destination port

It then takes the system time and tells the user what was sent, to where, and at what time


In tcpS.go, 

we have a function ScanConfigForServer which has a string parameter for the source and returns a string and error.  This function scans the config file for the port that will be used, and returns an error "cannot find port" if the port is not found for user convenience.

The exported function CreateUserInputStruct is used to easily create the userInput struct that is needed throughout the code (this contains the destination, message and source)

The handleConnection function is used to receive the messages and is used as a goroutine so there is a separate thread for it.  
We call the getDelayParams function here outside of the for loop so if we ever decided to add support for multi-messaging the delay doesn't get meaninglessly calculated.  


The exported function ConnectToTCPClient is used call the net.Listen function to connect the TCP client and keep a TCP Client connection constantly open.  This is why we have the for loop and goroutine handleConnection call.


###Shortcomings and Potential Improvemnts 
As of right now, each process and only send out one message each

One way to improve this would be to have the user's input go through a goroutine which can constantly read messages
and parse them to send to the sever

We are also sending raw strings over TCP channels, which in any situation more complex than 
the one string we are sending, would be ineffiecnt 

Can improve by using JSON or GOB encoding

There could also be a stop command which closes the channels and does not require the user to use 
command c 
