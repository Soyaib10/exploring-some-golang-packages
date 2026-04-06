package main

import (
	"fmt"

	"go.uber.org/fx"
)

// 1. Define a simple struct
type Message struct {
	Text string
}

// 2. A constructor function that returns the struct
func NewMessage() *Message {
	return &Message{Text: "Hello from Uber FX!"}
}

// 3. A function that uses the struct
func PrintMessage(msg *Message) {
	fmt.Println(msg.Text)
}

// normal way of calling
func NormalWay() {
	msg := NewMessage()
	PrintMessage(msg)
}

func FxWay() {
	// fx.New starts the application container
	app := fx.New(
		// Provide the constructor
		fx.Provide(NewMessage),

		// Invoke the function that needs the Message
		fx.Invoke(PrintMessage),
	)
	app.Run()
}