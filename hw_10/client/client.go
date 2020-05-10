package client

import (
	"bufio"
	"context"
	"fmt"
	"net"
	"os"
	"time"
)

type TelnetClient struct {
	conn          net.Conn
	stdinScanner  *bufio.Scanner
	serverScanner *bufio.Scanner
	serverCh      chan string
	stdinCh       chan string
}

func NewTelnetClient(addr string, timeOut time.Duration) (*TelnetClient, error) {
	dialer := &net.Dialer{}
	ctx, _ := context.WithTimeout(context.Background(), timeOut)

	conn, err := dialer.DialContext(ctx, "tcp", addr)
	if err != nil {
		return nil, err
	}

	return &TelnetClient{
		conn:          conn,
		stdinScanner:  bufio.NewScanner(os.Stdin),
		serverScanner: bufio.NewScanner(conn),
		serverCh:      make(chan string, 1),
		stdinCh:       make(chan string, 1),
	}, nil
}

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

func (t *TelnetClient) scanTCPServer() {
	for {
		if !t.serverScanner.Scan() {
			t.serverCh <- "Bye!\n"
			return
		}
		if _, err := os.Stdout.WriteString(t.serverScanner.Text() + "\n"); err != nil {
			t.serverCh <- err.Error()
		}
	}
}

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
