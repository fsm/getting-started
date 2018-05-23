package types

import "github.com/fsm/fsm"

var Coin = &fsm.Type{
	Slug: "coin",
	Options: []string{
		"coin",
		"quarter",
		"token",
	},
}
