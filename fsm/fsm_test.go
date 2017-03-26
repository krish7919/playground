package fsm

import (
	"log"
	"testing"
)

/*
A valid test fsm lokos like this

                                  +-----+                     +-----+
                     +-----e2---->+ s3  +-------e4----------->+ s5  |
                     |            |     |                     |     |
                     |            +-----+                     +-----+
                 +---+-+
    +-----e0---->+ s1  |
    |            |     |
    |            +---+-+
+---+-+              |            +-----+                    +-----+
| s0  |              +-----e3---->+ s4  +-------e5---------->+ s6  |
|     |                           |     |                    |     |
+---+-+                           +-----+                    +--+--+
    |                                                           ^
    |            +-----+                                        |
    +-----e1---->+ s2  +------------------------e6--------------+
                 |     |
                 +-----+
*/

var (
	s0 = &State{
		Name:      "s0",
		isInitial: true,
		isFinal:   false,
	}
	s1 = &State{
		Name:      "s1",
		isInitial: false,
		isFinal:   false,
	}
	s2 = &State{
		Name:      "s2",
		isInitial: false,
		isFinal:   false,
	}
	s3 = &State{
		Name:      "s3",
		isInitial: false,
		isFinal:   false,
	}
	s4 = &State{
		Name:      "s4",
		isInitial: false,
		isFinal:   false,
	}
	s5 = &State{
		Name:      "s5",
		isInitial: false,
		isFinal:   true,
	}
	s6 = &State{
		Name:      "s6",
		isInitial: false,
		isFinal:   true,
	}
	e0 = &Event{
		Name:      "e0",
		FromState: s0,
		ToState:   s1,
		transitionFunction: func() *State {
			log.Println("s0->s1")
			return s1
		},
	}
	e1 = &Event{
		Name:      "e1",
		FromState: s0,
		ToState:   s2,
		transitionFunction: func() *State {
			log.Println("s0->s2")
			return s2
		},
	}
	e2 = &Event{
		Name:      "e2",
		FromState: s1,
		ToState:   s3,
		transitionFunction: func() *State {
			log.Println("s1->s3")
			return s3
		},
	}
	e3 = &Event{
		Name:      "e3",
		FromState: s1,
		ToState:   s4,
		transitionFunction: func() *State {
			log.Println("s1->s4")
			return s4
		},
	}
	e4 = &Event{
		Name:      "e4",
		FromState: s3,
		ToState:   s5,
		transitionFunction: func() *State {
			log.Println("s3->s5")
			return s5
		},
	}
	e5 = &Event{
		Name:      "e5",
		FromState: s4,
		ToState:   s6,
		transitionFunction: func() *State {
			log.Println("s4->s6")
			return s6
		},
	}
	e6 = &Event{
		Name:      "e6",
		FromState: s2,
		ToState:   s6,
		transitionFunction: func() *State {
			log.Println("s2->s6")
			return s6
		},
	}
)

func getValidFSM() *FSM {
	states := []*State{s6, s5, s4, s3, s2, s1, s0}
	events := []*Event{e6, e5, e4, e3, e2, e1, e0}
	f, err := New("valid-test-fsm", states, events)
	if err != nil {
		log.Fatal(err)
	}
	err = f.Start()
	if err != nil {
		log.Fatal(err)
	}
	return f
}

func TestNext(t *testing.T) {
	f := getValidFSM()
	err := f.Next(e0)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if f.currentState != e0.ToState {
		t.Fatalf("Expected %s, got %s", e0.ToState.Name, f.currentState.Name)
	}

	err = f.Next(e1)
	if err != ErrInvalidEvent {
		t.Fatalf("Expected ErrInvalidEvent,Got %s", err)
	}

	// TODO more state transitions here
}

// More tests
//func getMultipleInitialStatesFSM() {
//	//ErrMultipleInitialStates
//}

//func getNoFinalStateFSM() {
//	//ErrNoFinalState
//}
//ErrDuplicateState
//ErrDuplicateEvent
//ErrHandlingFSMEvent
//Next
