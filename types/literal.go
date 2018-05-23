package types

import "github.com/fsm/fsm"

var Literal = &fsm.Type{
	Slug: "literal",
	IsValid: func(string) bool {
		return true
	},
}
