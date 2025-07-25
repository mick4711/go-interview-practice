// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	"fmt"
	"sync"
)

// Client represents a connected chat client
type Client struct {
	// username, message channel, mutex, disconnected flag
	username     string
	messages     chan string
	mu           sync.Mutex
	disconnected bool
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	if c.disconnected {
		return
	}

	// thread-safe
	c.mu.Lock()
	defer c.mu.Unlock()

	// non-blocking
	select {
	case c.messages <- message:
		fmt.Printf("Client [%s] has been SENT message:[%s]\n", c.username, message)
	default:
	}
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// read from channel, handle closed channel
	if message, ok := <-c.messages; ok {
		fmt.Printf("Client [%s] RECEIVED message:[%s]\n", c.username, message)
		return message
	}

	fmt.Printf("Client [%s] no more messages\n", c.username)
	return "no more messages"
}

type BroadcastMessage struct {
	sender *Client
	msg    string
}

// ChatServer manages client connections and message routing
type ChatServer struct {
	// clients map, mutex
	clients    map[string]*Client
	broadcast  chan BroadcastMessage
	connect    chan *Client
	disconnect chan *Client
	mu         sync.Mutex
}

// NewChatServer creates a new chat server instance
func NewChatServer() *ChatServer {
	return &ChatServer{
		clients: make(map[string]*Client),
	}
}

// Connect adds a new client to the chat server
func (s *ChatServer) Connect(username string) (*Client, error) {
	// check username, create client, add to map
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.clients[username]; ok {
		return nil, ErrUsernameAlreadyTaken
	}

	client := &Client{
		username:     username,
		messages:     make(chan string,1), //TODO see if working with no buffer
		disconnected: false,
	}
	s.clients[username] = client

	return client, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// remove from map, close channels
	delete(s.clients, client.username)
	close(client.messages)
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	// TODO: mutex?
	// format message, send to all clients
	for _, client := range s.clients {
		client.Send(fmt.Sprintf("From:%s\nMessage:%s\n", sender.username, message))
	}
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	// find recipient, check errors, send message
	receiver, ok := s.clients[recipient]
	if !ok {
		return ErrRecipientNotFound
	}

	fmt.Printf("Sender [%s] sending to:[%s] private message: [%s]\n", sender.username, recipient, message)
	receiver.Send(fmt.Sprintf("PRIVATE message from [%s] to:[%s] content: %s", sender.username, recipient, message))

	return nil
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
)
