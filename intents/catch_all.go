package intents

import (
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/types"
)

var CatchAll = &fsm.Intent{
	Slug: "catch-all",
	Slots: map[string]*fsm.Type{
		"input": types.Literal,
	},
	Utterances: []string{
		"{input}",
	},
}
