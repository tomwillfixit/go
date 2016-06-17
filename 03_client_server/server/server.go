//usr/bin/env go run $0 $@; exit //allows you to run ./server.go

// Based on : https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

package main 

import "net" 
import "fmt" 
import "bufio" 

// Simple Server that accepts a single client connection
func main() {   

	fmt.Println("Server running ...")   

	// listen on all interfaces. _ means check for the present of a value but not the actual value 
	ln, _ := net.Listen("tcp", ":8081")   

	// accept connection on port   
	conn, _ := ln.Accept()   

	// run loop forever (or until ctrl-c)   
	for {     

	// will listen for message to process ending in newline (\n)     
	message, _ := bufio.NewReader(conn).ReadString('\n')     

	// output message received     
	fmt.Print("Message Received from Client :", string(message))     

	// send new string back to client     
	conn.Write([]byte(message + "\n"))   
} 
}
