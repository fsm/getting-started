<a href="https://github.com/fsm"><p align="center"><img src="https://user-images.githubusercontent.com/2105067/35464215-a014d512-02a9-11e8-8913-63a066f6064e.png" alt="FSM" width="350px" align="center;"/></p></a>

# Getting Started

[![Version](https://img.shields.io/github/tag/fsm/getting-started.svg)](https://github.com/fsm/getting-started/releases)
[![MIT License](https://img.shields.io/badge/License-MIT-blue.svg)](https://github.com/fsm/getting-started/blob/master/LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/fsm/getting-started)](https://goreportcard.com/report/github.com/fsm/getting-started)
[![Gitter](https://img.shields.io/gitter/room/nwjs/nw.js.svg)](https://gitter.im/fsm/Lobby)

This package contains an example project / scaffolding for [FSM](https://github.com/fsm/fsm).

The example included in this repository follows the classic turnstile state machine:

![State Machine Example](https://user-images.githubusercontent.com/2105067/35996819-8b99dcae-0ce5-11e8-8765-d63c9b070232.png)

# Setting Up

The first thing you're going to want to do is clone down the repository to the location you would like your project to exist.  This should be somewhere in your `$GOPATH`.

After you've cloned it down, you're going to need to update the package on [this line](https://github.com/fsm/getting-started/blob/master/main.go#L8), and [this line](https://github.com/fsm/getting-started/blob/master/bot_test.go#L9) that reference packages within this repo.  It will need to be updated to reflect the location where you cloned the repository.

# Building & Running

After your project is set up as specified in the above section, you'll need to install dependencies.

Assuming you have [dep](https://github.com/golang/dep) installed, you can run the following command in the root of the repository:

```sh
dep ensure
```

Now you can build and run the project.

```sh
go build -o bot && ./bot
```

# Understanding FSM

In order to successfully build something with FSM, you must understand three fundamental concepts:

- [Finite State Machines](#finite-state-machines)
- [States](#states)
- [Transitions](#transitions)

## Finite State Machines

> [A finite-state machine] is an abstract machine that can be in exactly one of a finite number of states at any given time. The FSM can change from one state to another in response to some external inputs; the change from one state to another is called a transition. An FSM is defined by a list of its states, its initial state, and the conditions for each transition.
>
> Wikipedia

The FSM library was built as purely against the definition of a formal finite-state machine as possible, so the Wikipedia definition holds true for this library.

> Looking at the diagram at the [top of this README](#getting-started), the entire diagram is the State Machine.

## States

[States](https://github.com/fsm/fsm/blob/master/fsm.go#L15-L21) are individual nodes within a [StateMachine](https://github.com/fsm/fsm/blob/master/fsm.go#L3-L4) that represent a meaningful point in the machine.

When a [Traverser](https://github.com/fsm/fsm/blob/master/fsm.go#L36-L47) enters a [State](https://github.com/fsm/fsm/blob/master/fsm.go#L15-L21), its [EntryAction](https://github.com/fsm/fsm/blob/master/fsm.go#L18) is immediately called.

> Looking at the diagram at the [top of this README](#getting-started), states are the yellow circles.

## Transitions

Transitions are the elements that make it possible for a traverser to switch their state as they traverse the state machine.

A [Transition](https://github.com/fsm/fsm/blob/master/fsm.go#L20) is defined in each state. This describes how to handle incoming data for a state.

If that data tells you that the [Traverser](https://github.com/fsm/fsm/blob/master/fsm.go#L36-L47) should switch states, just return the appropriate [BuildState](https://github.com/fsm/fsm/blob/master/fsm.go#L11-L13) function.

If the data tells you that you should not switch states, return nothing and the [ReentryAction](https://github.com/fsm/fsm/blob/master/fsm.go#L19) will be called, which is a great opportunity to reprompt for a response that will keep the conversation flowing.

> Looking at the diagram at the [top of this README](#getting-started), transitions are the black arrows.

# Beyond the CLI

If you're building a more serious bot, you're more likely going to want to use something other than [fsm/cli](https://github.com/fsm/cli) as your deployment target, and [fsm/cache-store](https://github.com/fsm/cache-store) as your data store.

All FSM targets can be viewed [here](https://github.com/search?q=topic%3Afsm-target+org%3Afsm&type=Repositories).

All FSM stores can be viewed [here](https://github.com/search?q=topic%3Afsm-store+org%3Afsm&type=Repositories).

View their respective READMEs for usage information.

# License

[MIT](LICENSE.md)
