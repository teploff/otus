package client

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/Netflix/go-expect"
	"github.com/stretchr/testify/assert"
	"github.com/teploff/otus/hw_10/server"
)

const tcpAddr = "0.0.0.0:12555"

// TestTCPServerNotAvailable check available connection of the tcp server
func TestTCPServerNotAvailable(t *testing.T) {
	tn, err := NewTelnetClient(tcpAddr, 10*time.Millisecond, os.Stdin)
	assert.Error(t, err)
	assert.Nil(t, tn)
}

// TestStdinStopTelnet check pressing ctrl+D by user
func TestStdinStopTelnet(t *testing.T) {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout), expect.WithStdin(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	defer assert.NoError(t, c.Close())

	srv, err := server.NewTCPServer(tcpAddr)
	assert.NoError(t, err)
	assert.NotNil(t, srv)

	stdin := c.Tty()
	tn, err := NewTelnetClient(tcpAddr, 100*time.Millisecond, stdin)
	assert.NoError(t, err)
	assert.NotNil(t, tn)

	go srv.Listen()
	defer srv.GracefulStop()

	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
	OUTER:
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
				stdin.Close()
				break OUTER
			default:
				c.SendLine("hello")
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()
	assert.NoError(t, tn.Run())
}

// TestTCPServerCloseConnection check closing connection from server
func TestTCPServerCloseConnection(t *testing.T) {
	c, err := expect.NewConsole(expect.WithStdout(os.Stdout), expect.WithStdin(os.Stdin))
	if err != nil {
		log.Fatal(err)
	}
	defer assert.NoError(t, c.Close())

	srv, err := server.NewTCPServer(tcpAddr)
	if err != nil {
		log.Fatalln(err)
	}
	go srv.Listen()

	stdin := c.Tty()
	tn, err := NewTelnetClient(tcpAddr, 100*time.Millisecond, stdin)
	if err != nil {
		log.Fatalln(err)
	}

	ticker := time.NewTicker(time.Millisecond * 100)
	go func() {
	EXIST:
		for {
			select {
			case <-ticker.C:
				ticker.Stop()
				srv.GracefulStop()
				break EXIST
			default:
				c.SendLine("how are you?")
				time.Sleep(time.Millisecond * 10)
			}
		}
	}()

	log.Println(tn.Run())
}
