package states

import (
	"fmt"

	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/intents"
)

// GetLockedState is a fsm.BuildState that returns a fsm.State that
// represents the state of "locked" for a turnstyle machine.
func GetLockedState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "locked",
		Entry: func(isReentry bool) error {
			// If we're kicked back to this state, it means that we don't understand the
			// users intent. In this state, the user can only insert a coin.
			if isReentry {
				emitter.Emit("Sorry, I don't quite understand.")
			} else {
				// Let user know how many times it's been turned
				turnCount := traverser.Fetch(varTurnCount)
				if turnCount.(int) != 0 {
					emitter.Emit(fmt.Sprintf("My turn counter now reads '%v'!", turnCount))
				}
			}

			// Ask the user for a coin
			emitter.Emit(emitable.QuickReply{
				Message:       "I am currently locked, and need a coin to be unlocked.",
				RepliesFormat: "You can say %v to unlock me.",
				Replies:       []string{"insert coin"},
			})
			return nil
		},
		ValidIntents: func() []*fsm.Intent {
			return []*fsm.Intent{
				intents.InsertCoin,
			}
		},
		Transition: func(intent *fsm.Intent, params map[string]string) *fsm.State {
			switch intent {
			case intents.InsertCoin:
				return GetUnlockedState(emitter, traverser)
			}
			return nil
		},
	}
}
