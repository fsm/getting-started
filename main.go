package main

import (
	"github.com/fsm/cli"

	cachestore "github.com/fsm/cache-store"
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/states"
)

func main() {
	cli.Start(GetStateMachine(), GetStore())
}

func GetStateMachine() fsm.StateMachine {
	return fsm.StateMachine{
		states.GetStartState,
		states.GetLockedState,
		states.GetUnlockedState,
	}
}

func GetStore() fsm.Store {
	return cachestore.New()
}
