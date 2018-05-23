package intents

import (
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/types"
)

var InsertCoin = &fsm.Intent{
	Slug: "insert-coin",
	Slots: map[string]*fsm.Type{
		"coin": types.Coin,
	},
	Utterances: []string{
		"Insert {coin}.",
		"Insert a {coin}",
		"Insert the {coin}.",
	},
}
