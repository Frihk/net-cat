package test
package test

import (
	"net"
	"testing"
	"time"
)

func TestClientJoin(t *testing.T) {
	server := NewServer(2)
	go server.Run()

	alice := &Client{
		Conn:     nil,
		Name:     "Alice",
		Messages: make(chan string, 10),
	}

	bob := &Client{
		Conn:     nil,
		Name:     "Bob",
		Messages: make(chan string, 10),
	}

	server.Join <- alice
	server.Join <- bob

	time.Sleep(50 * time.Millisecond)

	if len(server.Clients) != 2 {
		t.Fatalf("expected 2 clients, got %d", len(server.Clients))
	}

	// Bob joining should generate system message to Alice
	select {
	case msg := <-alice.Messages:
		if msg == "" {
			t.Fatal("Alice received empty join message")
		}
	case <-time.After(100 * time.Millisecond):
		t.Fatal("Alice did not receive join broadcast")
	}
}