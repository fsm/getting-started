package states

import (
	"github.com/fsm/fsm"
)

const varTurnCount = "turn-count"

// GetStartState is a fsm.BuildState that returns a fsm.State that is
// the entry point for the state machine.
func GetStartState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "start",
		EntryAction: func() error {
			// Set the initial number of turns to 0
			traverser.Upsert(varTurnCount, 0)

			// Say hello!
			emitter.Emit("Hello and welcome! I am Turnstile 3000.")
			return nil
		},
		ReentryAction: func() error {
			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			return GetLockedState(emitter, traverser)
		},
	}
}
