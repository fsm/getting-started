package states

import (
	"fmt"

	"github.com/fsm/emitable"
	"github.com/fsm/fsm"
)

func GetLockedState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "locked",
		EntryAction: func() error {
			// Let user know how many times it's been turned
			turnCount, err := traverser.Fetch(varTurnCount)
			if err != nil {
				return err
			}

			// If we haven't met, introduce Turnstile 3000
			if turnCount.(int) == 0 {
				emitter.Emit("Hello and welcome! I am Turnstile 3000.")
			} else {
				// If it's been turned before, let the user know how many times
				emitter.Emit(fmt.Sprintf("Thank you for using Turnstile 3000."))
				if turnCount.(int) == 1 {
					emitter.Emit(fmt.Sprintf("So far I have been turned %v time.", turnCount))
				} else {
					emitter.Emit(fmt.Sprintf("So far I have been turned %v times.", turnCount))
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
		ReentryAction: func() error {
			// If the user said anything other than "insert coin", prompt them again
			emitter.Emit(emitable.QuickReply{
				Message:       "Sorry, I don't quite understand.",
				RepliesFormat: "You can say %v to unlock me.",
				Replies:       []string{"insert coin"},
			})

			return nil
		},
		Transition: func(input interface{}) *fsm.State {
			switch v := input.(type) {
			case string:
				if v == "insert coin" {
					// Go to unlocked
					return GetUnlockedState(emitter, traverser)
				}
			}

			return nil
		},
	}
}
