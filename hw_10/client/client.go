package client

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

// TelnetClient instance of tcp-client as telnet
type TelnetClient struct {
	conn          net.Conn
	stdinScanner  *bufio.Scanner
	serverScanner *bufio.Scanner
	serverCh      chan string
	stdinCh       chan string
}

// NewTelnetClient getting instance of telnet
func NewTelnetClient(addr string, timeOut time.Duration, reader io.Reader) (*TelnetClient, error) {
	dialer := &net.Dialer{}
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TelnetClient{
		conn:          conn,
		stdinScanner:  bufio.NewScanner(reader),
		serverScanner: bufio.NewScanner(conn),
		serverCh:      make(chan string, 1),
		stdinCh:       make(chan string, 1),
	}, nil
}

// scanStdin scans IOF of io.Reader of STDIN
func (t *TelnetClient) scanStdin() {
	for {
		if !t.stdinScanner.Scan() {
			t.stdinCh <- "Bye!\n"
			return
		}
		msg := t.stdinScanner.Text()
		if _, err := t.conn.Write([]byte(fmt.Sprintf("%s\n", msg))); err != nil {
			t.stdinCh <- fmt.Sprintf("server error: %s\n", err.Error())
			return
		}
	}
}

// scanStdin scans IOF of io.Reader of tcp-server
func (t *TelnetClient) scanTCPServer() {
	for {
		if !t.serverScanner.Scan() {
			t.serverCh <- "Server close connection!\nBye!\n"
			return
		}
		if _, err := os.Stdout.WriteString(t.serverScanner.Text() + "\n"); err != nil {
			t.serverCh <- err.Error()
		}
	}
}

// Run launches telnet client
func (t *TelnetClient) Run() error {
	defer t.conn.Close()
	defer os.Stdin.Close()
	defer os.Stdout.Close()

	go t.scanStdin()
	go t.scanTCPServer()

	for {
		select {
		case msg := <-t.serverCh:
			_, err := os.Stdout.WriteString(msg)

			return err
		case msg := <-t.stdinCh:
			_, err := os.Stdout.WriteString(msg)

			return err
		default:
			time.Sleep(10 * time.Millisecond)
		}
	}
}
