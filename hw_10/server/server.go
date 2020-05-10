package server

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

// TCPServer tcp-server for helping debugging
type TCPServer struct {
	listener    net.Listener
	connections []net.Conn
}

// NewTCPServer returns instance of tcp-server
func NewTCPServer(addr string) (*TCPServer, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, fmt.Errorf("cannot listen: %v", err)
	}

	return &TCPServer{listener: l}, nil
}

// Listen starts getting request
func (ts *TCPServer) Listen() error {
	for {
		conn, err := ts.listener.Accept()
		if err != nil {
			return fmt.Errorf("cannot accept: %v", err)
		}
		ts.connections = append(ts.connections, conn)
		go ts.handleConnection(conn)
	}
}

// GracefulStop closing all active connections and close tcp server
func (ts *TCPServer) GracefulStop() error {
	for _, conn := range ts.connections {
		if err := conn.Close(); err != nil {
			return err
		}
	}
	return ts.listener.Close()
}

// closeConn closing active connections
func (ts *TCPServer) closeConn(conn net.Conn) error {
	for index, currConn := range ts.connections {
		if currConn == conn {
			ts.connections = append(ts.connections[:index], ts.connections[index+1:]...)
			return conn.Close()
		}
	}

	return fmt.Errorf("not found connection")
}

// handleConnection handling request
func (ts *TCPServer) handleConnection(conn net.Conn) {
	defer ts.closeConn(conn)
	conn.Write([]byte(fmt.Sprintf("Welcome to %s, friend from %s\n", conn.LocalAddr(), conn.RemoteAddr())))

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		log.Printf("RECEIVED: %s", text)
		if text == "quit" || text == "exit" {
			break
		}

		conn.Write([]byte(fmt.Sprintf("I have received '%s'\n", text)))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error happend on connection with %s: %v", conn.RemoteAddr(), err)
	}

	log.Printf("Closing connection with %s", conn.RemoteAddr())
}
