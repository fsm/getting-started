package main_test

import (
	"testing"
)

func TestChatbot(t *testing.T) {
	/*
		traverser := test.New(bot.GetStateMachine(), bot.GetStore())

		// Send initial message
		traverser.Send("Hi")

		// Verify we're in the proper state
		assert.Equal(t, traverser.CurrentState(), "locked")

		// Validate turn-count variable
		turnCount, _ := traverser.Fetch("turn-count")
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

		// Go through the turnstyle 5 times
		for i := 1; i <= 5; i++ {
			// Insert a coin
			traverser.Send("insert coin")
			assert.Equal(t, traverser.CurrentState(), "unlocked")

			// Push the turnstyle
			traverser.Send("push")
			assert.Equal(t, traverser.CurrentState(), "locked")
			turnCount, _ = traverser.Fetch("turn-count")
			assert.Equal(t, turnCount.(int), i)
		}

		// Turn count should be 5 now
		turnCount, _ = traverser.Fetch("turn-count")
		assert.Equal(t, turnCount.(int), 5)
	*/
}
