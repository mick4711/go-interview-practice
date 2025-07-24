// Package challenge8 contains the solution for Challenge 8: Chat Server with Channels.
package challenge8

import (
	"errors"
	"sync"
)

// Client represents a connected chat client
type Client struct {
	// username, message channel, mutex, disconnected flag
	username     string
	messages     chan string
	server       *ChatServer
	mu           sync.Mutex
	disconnected bool
}

// Send sends a message to the client
func (c *Client) Send(message string) {
	// TODO: Implement this method
	// Hint: thread-safe, non-blocking send
	c.messages <- message
}

// Receive returns the next message for the client (blocking)
func (c *Client) Receive() string {
	// TODO: Implement this method
	// Hint: read from channel, handle closed channel
	return <-c.messages
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
	if _, ok := s.clients[username]; ok {
		return nil, ErrUsernameAlreadyTaken
	}

	client := &Client{
		username:     username,
		messages:     make(chan string),
		server:       s,
		mu:           sync.Mutex{},
		disconnected: false,
	}
	s.clients[username] = client
	// TODO  s.connect <- client

	return client, nil
}

// Disconnect removes a client from the chat server
func (s *ChatServer) Disconnect(client *Client) {
	// TODO: Implement this method
	// Hint: remove from map, close channels
	delete(s.clients, client.username)
}

// Broadcast sends a message to all connected clients
func (s *ChatServer) Broadcast(sender *Client, message string) {
	// TODO: Implement this method
	// Hint: format message, send to all clients
}

// PrivateMessage sends a message to a specific client
func (s *ChatServer) PrivateMessage(sender *Client, recipient string, message string) error {
	// TODO: Implement this method
	// Hint: find recipient, check errors, send message
	return nil
}

// Common errors that can be returned by the Chat Server
var (
	ErrUsernameAlreadyTaken = errors.New("username already taken")
	ErrRecipientNotFound    = errors.New("recipient not found")
	ErrClientDisconnected   = errors.New("client disconnected")
	// Add more error types as needed
)
