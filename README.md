# Simple TCP Server in Go

This project is a basic TCP server implemented in Go.

## Features
- Listens on a TCP port (`:3000`)
- Accepts multiple client connections
- Reads messages sent by clients
- Sends back a confirmation response

## File Structure
- `main.go`: Entry point of the server
- `newServer.go`: Defines the Server struct and constructor
- `accept.go`: Handles incoming connections
- `read.go`: Reads messages from clients
- `start.go`: Starts the server and manages shutdown

## How to Run
```bash
go run main.go
```

## Testing with Telnet
Make sure your server is running, then open a new terminal and connect using:
```bash
telnet localhost 3000
```

Once connected, type a message and press Enter. You should receive:
```
Thank you for message!
```

Server log output will look like:
```
Started at :3000
new connection to the server 127.0.0.1:50432
received message from connection(127.0.0.1:50432): hello server
```

---


