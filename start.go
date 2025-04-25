package main

import (
	"fmt"
	"net"
)

func (s *Server) Start() error{
	ln , err := net.Listen("tcp",s.listenAddr)
	fmt.Println("Started at",s.listenAddr)
	if err != nil{
	   return err 
	}
	defer ln.Close()
	s.ln = ln ; 
	go s.AcceptLoop()
	<-s.quitch // wth is this 

	//each time  we close the server  also cloce the channel 
	close(s.msgch);

	return  nil 
}

