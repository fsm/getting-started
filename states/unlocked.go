package states

import (
	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/intents"
)

// GetUnlockedState is a fsm.BuildState that returns a fsm.State that
// represents the state of "unlocked" for a turnstile machine.
func GetUnlockedState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "unlocked",
		Entry: func(isReentry bool) error {
			if isReentry {
				emitter.Emit("Sorry, I don't quite understand.")
			}
			emitter.Emit(emitable.QuickReply{
				Message:       "I am now unlocked, and awaiting to be pushed.",
				RepliesFormat: "You can say %v to turn me.",
				Replies:       []string{"push"},
			})
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return []*fsm.Intent{
				intents.PushTurnstile,
			}
		},
		Transition: func(intent *fsm.Intent, params map[string]string) *fsm.State {
			switch intent {
			case intents.PushTurnstile:
				// Update turn count
				turnCount := traverser.Fetch(varTurnCount)
				traverser.Upsert(varTurnCount, turnCount.(int)+1)

				// Go back to locked
				return GetLockedState(emitter, traverser)
			}
			return nil
		},
	}
}
