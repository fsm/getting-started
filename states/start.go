package states

import (
	"github.com/fsm/fsm"
)

const varTurnCount = "turn-count"

func GetStartState(emitter fsm.Emitter, traverser fsm.Traverser) *fsm.State {
	return &fsm.State{
		Slug: "start",
		EntryAction: func() error {
			// Set the initial number of turns to 0
			traverser.Upsert(varTurnCount, 0)

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
