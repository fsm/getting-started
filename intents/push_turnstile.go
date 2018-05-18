package intents

import "github.com/fsm/fsm"

var PushTurnstile = &fsm.Intent{
	Slug: "push-turnstile",
	Utterances: []string{
		"push",
		"push turnstile",
		"push the turnstile",
	},
}
