package states

import (
	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
)

// GetUnlockedState is a fsm.BuildState that returns a fsm.State that
// represents the state of "unlocked" for a turnstyle machine.
func GetUnlockedState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "unlocked",
		EntryAction: func() error {
			// Ask user to push the turnstile
			emitter.Emit(emitable.QuickReply{
				Message:       "I am now unlocked, and awaiting to be pushed.",
				RepliesFormat: "You can say %v to turn me.",
				Replies:       []string{"push"},
			})

			return nil
		},
		ReentryAction: func() error {
			// If the user said anything other than "push", prompt them again
			emitter.Emit(emitable.QuickReply{
				Message:       "Sorry, I don't quite understand.",
				RepliesFormat: "You can say %v to turn me.",
				Replies:       []string{"push"},
			})

			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			switch v := input.(type) {
			case string:
				if v == "push" {
					// Update turn count
					turnCount, _ := traverser.Fetch(varTurnCount)
					traverser.Upsert(varTurnCount, turnCount.(int)+1)

					// Go back to locked
					return GetLockedState(emitter, traverser)
				}
			}

			return nil
		},
	}
}
