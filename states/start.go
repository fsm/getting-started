package states

import (
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/intents"
)

const varTurnCount = "turn-count"

// GetStartState is a fsm.BuildState that returns a fsm.State that is
// the entry point for the state machine.
func GetStartState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "start",
		Entry: func(isReentry bool) error {
			// Set the initial number of turns to 0
			traverser.Upsert(varTurnCount, 0)

			// Say hello!
			emitter.Emit("Hello and welcome! I am Turnstile 3000.")
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return []*fsm.Intent{intents.Literal}
		},
		Transition: func(*fsm.Intent, map[string]string) *fsm.State {
			return GetLockedState(emitter, traverser)
		},
	}
}
