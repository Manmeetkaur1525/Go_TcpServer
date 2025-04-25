package main

import (
	"fmt"
	"log"
	"net"

)

type Message struct{
	from string
	payload []byte
}
type Server struct{
	listenAddr string
	ln  net.Listener
	quitch chan struct{} //what is channel why is it used 
	msgch  chan Message
	//map of connection 
	// peerMap map[net.Addr]
	
}

func main(){
 server := NewServer(":3000")
 go func(){ //making a go routine 
	for msg := range server.msgch{
		fmt.Println("received message from connection(%s)%s\n" ,msg.from,string(msg.payload) )
	}
 }()
 log.Fatal(server.Start())
}