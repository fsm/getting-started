package main_test

import (
	"testing"

	"github.com/fsm/emitable"
	bot "github.com/fsm/getting-started"
	"github.com/fsm/test"
	"github.com/stretchr/testify/assert"
)

func TestChatbot(t *testing.T) {
	traverser := test.New(bot.GetStateMachine(), bot.GetStore())

	// Send initial message
	traverser.Send("Hi")

	// Verify we're in the proper state
	assert.Equal(t, traverser.CurrentState(), "locked")

	// Validate turn-count variable
	turnCount := traverser.Fetch("turn-count")
	assert.Equal(t, turnCount.(int), 0)

	// Validate that messages are being emitted received
	assert.Equal(t,
		"Hello and welcome! I am Turnstile 3000.",
		traverser.GetReceived(),
	)
	assert.Equal(t,
		"I am currently locked, and need a coin to be unlocked.",
		traverser.GetReceived().(emitable.QuickReply).Message,
	)

	// Ensure we don't switch states with invalid input
	traverser.Send("HELLO")
	assert.Equal(t,
		"Sorry, I don't quite understand.",
		traverser.GetReceived(),
	)
	assert.Equal(t, traverser.CurrentState(), "locked")
	traverser.GetAllReceived() // Flush

	// Go through the turnstyle 5 times
	for i := 1; i <= 5; i++ {
		// Insert a coin
		traverser.Send("insert coin")
		assert.Equal(t, traverser.CurrentState(), "unlocked")
		traverser.GetAllReceived() // Flush

		// Ensure invalid input won't switch states
		traverser.Send("hello")
		assert.Equal(t,
			"Sorry, I don't quite understand.",
			traverser.GetReceived(),
		)
		assert.Equal(t, traverser.CurrentState(), "unlocked")
		traverser.GetAllReceived() // Flush

		// Push the turnstyle
		traverser.Send("push")
		assert.Equal(t, traverser.CurrentState(), "locked")
		turnCount = traverser.Fetch("turn-count")
		assert.Equal(t, turnCount.(int), i)
	}

	// Turn count should be 5 now
	turnCount = traverser.Fetch("turn-count")
	assert.Equal(t, turnCount.(int), 5)
}
