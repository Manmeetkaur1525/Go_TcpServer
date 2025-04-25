package main

func NewServer(listenAddr string)*Server{
	return &Server{
		listenAddr: listenAddr,
		quitch: make(chan struct{}),
		msgch: make(chan Message,10), //buffer it 

		 
	}

}