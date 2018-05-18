package intents

import "github.com/fsm/fsm"

var InsertCoin = &fsm.Intent{
	Slug: "insert-coin",
	Slots: map[string]*fsm.Type{
		"coin": &fsm.Type{
			Slug: "coin",
			Options: []string{
				"coin",
				"quarter",
				"token",
			},
		},
	},
	Utterances: []string{
		"Insert {coin}.",
		"Insert a {coin}",
		"Insert the {coin}.",
	},
}
