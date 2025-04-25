package main

import (
	"fmt"
)

func (s *Server) AcceptLoop(){
	for{
		conn , err := s.ln.Accept()
		if err != nil {
			fmt.Println("accept error" , err)
			continue
		}
		fmt.Println("new connection to the server",conn.RemoteAddr())
		go s.ReadLoop(conn)
	}
}