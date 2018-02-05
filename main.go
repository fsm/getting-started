package main

import (
	cachestore "github.com/fsm/cache-store"
	"github.com/fsm/cli"
	"github.com/fsm/fsm"
	"github.com/fsm/getting-started/states"
)

func main() {
	cli.Start(getStateMachine(), getStore())
}

func getStateMachine() fsm.StateMachine {
	return fsm.StateMachine{
		states.GetStartState,
		states.GetLockedState,
		states.GetUnlockedState,
	}
}

func getStore() fsm.Store {
	return cachestore.New()
}
