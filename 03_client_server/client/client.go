//usr/bin/env go run $0 $@; exit //allows you to run ./client.go

// Based on : https://systembash.com/a-simple-go-tcp-server-and-tcp-client/

package main 

import "net" 
import "fmt" 
import "bufio" 
import "os" 

func main() {   

	// connect to this socket   
	conn, _ := net.Dial("tcp", "server:8081")  // server is the container name in docker-compose.yml 

	for {     
		// read in input from stdin    
		reader := bufio.NewReader(os.Stdin)     

		fmt.Print("Type a message and hit enter to send to server : ")     
		text, _ := reader.ReadString('\n')     

		if(text == "exit\n") {
        		fmt.Print("Exiting client")
			os.Exit(0)
    		}
		// send to socket     
		fmt.Fprintf(conn, text + "\n")     

		// listen for reply     
		message, _ := bufio.NewReader(conn).ReadString('\n')     
		fmt.Print("Server successfully received : "+ message + "\n")   
	} 
}
