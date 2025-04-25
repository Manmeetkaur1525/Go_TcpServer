package main

import ("fmt"
"net")

func (s *Server) ReadLoop(conn net.Conn){
	buf := make([]byte , 2048)
	defer conn.Close()
	for{
		n , err := conn.Read(buf)
		if err != nil{
			fmt.Println("Read error",err)
			continue
		}
		// msg := buf[:n] //amount we let 
		// fmt.Println(string(msg)); 
		//each time me receive a message 

		//from this message we also get who is this , who is doing it 
		s.msgch <- Message{
			from: conn.RemoteAddr().String() ,
			payload: buf[:n],
		}
		conn.Write([]byte("Thank you for message!"))
	}
}