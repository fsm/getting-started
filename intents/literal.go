package intents

import "github.com/fsm/fsm"

var Literal = &fsm.Intent{
	Slug: "literal",
	Slots: map[string]*fsm.Type{
		"any": &fsm.Type{
			Slug: "literal",
			IsValid: func(string) bool {
				return true
			},
		},
	},
	Utterances: []string{
		"{any}",
	},
}
