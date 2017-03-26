//Package fsm provides a Finite State Machine implementation in Golang
package fsm

import (
	"errors"
	"log"
	"reflect"
)

var (
	// ErrInvalidEvent is returned when the event cannot be applied to the
	// current fsm state
	ErrInvalidEvent = errors.New("Invalid event received by the FSM")

	// ErrDuplicateState is returned during initialization when the newly
	// defined state already exists in the fsm
	ErrDuplicateState = errors.New("Found a duplicate state in the FSM")

	// ErrDuplicateEvent is returned during initialization when the newly
	// defined event already exists in the fsm
	ErrDuplicateEvent = errors.New("Found a duplicate event in the FSM")

	// ErrMultipleInitialStates is returned during initialization when there exists
	// two initial states for the FSM
	ErrMultipleInitialStates = errors.New("Multiple initial states found in the FSM")

	// ErrNoFinalState is returned during initialization when there are no final
	// state/s defined for the FSM
	ErrNoFinalState = errors.New("No final state found in the FSM")

	// ErrFSMInit is returned when the fsm cannot be moved to the initial state;
	// when no initial state is found
	ErrFSMInit = errors.New("Cannot Start() the FSM")

	// ErrHandlingFSMEvent is returned when the transition function defined in
	// an event does not return the desired final state
	ErrHandlingFSMEvent = errors.New("Unknown error in transition function")
)

// State is a struct which defines a particular state of the fsm.
// Every State object has a unique Name, and is tagged with informationabout
// whether it is the initial or the final state
type State struct {
	Name      string
	isInitial bool
	isFinal   bool
}

// IsInitialState returns whether the current state is the starting/initial state
func (s *State) IsInitialState() bool { return s.isInitial }

// IsFinalState returns whether the current state is the final/terminated state
func (s *State) IsFinalState() bool { return s.isFinal }

// isValidState verifies that a state cannot be in both initial and final state
func (s *State) isValidState() bool {
	return !(s.isFinal == true && s.isInitial == true)
}

// transitionFunction returns the new state of the fsm
type Event struct {
	Name               string
	FromState          *State
	ToState            *State
	transitionFunction func() *State
}

// FSM is a struct which defines a new finite state machine. It contains a list
// of states and the corresponding events.
type FSM struct {
	Name         string
	states       []*State
	events       []*Event
	currentState *State
}

// New create an fsm object with the specified states and events.
// It expects each event and state passed to be unique, and that there is only one initial state.
// It does not initialize/start the fsm, you have to call Start() for that.
func New(fsmName string, definedStates []*State,
	definedEvents []*Event) (*FSM, error) {
	fsm := new(FSM)
	fsm.Name = fsmName
	fsm.currentState = nil

	// add all the defined states
	log.Printf("Adding state: %s\n", definedStates[0].Name)
	fsm.states = []*State{definedStates[0]}
	for _, state := range definedStates[1:] {
		if fsm.isNewState(state) && state.isValidState() {
			log.Printf("Adding state: %s\n", state.Name)
			fsm.states = append(fsm.states, state)
		} else {
			return nil, ErrDuplicateState
		}
	}

	// verify if there is only one initial state
	var tempState *State = nil
	for _, state := range fsm.states {
		if state.IsInitialState() {
			if tempState == nil {
				tempState = state
			} else {
				return nil, ErrMultipleInitialStates
			}
		}
	}

	// verify there is at least one final state
	tempState = nil
	for _, state := range fsm.states {
		if state.IsFinalState() {
			tempState = state
			break
		}
	}
	if tempState == nil {
		return nil, ErrNoFinalState
	}

	// add all the defined events
	log.Printf("Adding event: %s\n", definedEvents[0].Name)
	fsm.events = []*Event{definedEvents[0]}
	for _, event := range definedEvents[1:] {
		// BUG(krish7919): verify every event must have a transition function?
		if fsm.isNewEvent(event) {
			log.Printf("Adding event: %s\n", event.Name)
			fsm.events = append(fsm.events, event)
		} else {
			return nil, ErrDuplicateEvent
		}
	}

	return fsm, nil
}

func (fsm *FSM) GetStates() []*State { return fsm.states }

func (fsm *FSM) GetEvents() []*Event { return fsm.events }

func (fsm *FSM) GetCurrentState() *State { return fsm.currentState }

func (fsm *FSM) GetName() string { return fsm.Name }

// isNewState determines if the State s passed to the fsm is a unique state
func (f *FSM) isNewState(s *State) bool {
	for _, state := range f.states {
		if reflect.DeepEqual(s, state) {
			return false
		}
	}
	return true
}

// isNewEvent determines if the Event e passed to the fsm is a unique event
func (f *FSM) isNewEvent(e *Event) bool {
	for _, event := range f.events {
		//BUG(krish7919): Can a delta in only transition_fn affect this?
		if reflect.DeepEqual(e, event) {
			return false
		}
	}
	return true
}

// isDefinedEvent determines if the event passed to the fsm is an event that the
// fsm is supposed to know about.
// It checks the current event with the predefined events specified in New()
func (fsm *FSM) isDefinedEvent(e *Event) bool {
	for _, event := range fsm.events {
		if reflect.DeepEqual(event, e) {
			return true
		}
	}
	return false
}

// isValidEvent returns whether the fsm is not in the final state and if the
// current state can handle the provided event
func (fsm *FSM) isValidEvent(e *Event) bool {
	if !fsm.currentState.isFinal && fsm.currentState == e.FromState &&
		fsm.isDefinedEvent(e) {
		return true
	}
	return false
}

// Start moves the current state of the fsm to the initial state.
// We have already verified that there is only one initial state in New()
func (fsm *FSM) Start() error {
	for _, state := range fsm.states {
		if state.IsInitialState() {
			fsm.currentState = state
			return nil
		}
	}
	return ErrFSMInit
}

// Next determines the next state of the fsm based on the event supplied.
// If everything is fine, it moves the current state of the fsm as defined in
// the event.
func (fsm *FSM) Next(e *Event) error {
	if !fsm.isValidEvent(e) {
		return ErrInvalidEvent
	}
	// call the transition function
	var newState *State = nil
	newState = e.transitionFunction()
	if reflect.DeepEqual(newState, e.ToState) {
		fsm.currentState = e.ToState
		return nil
	}
	return ErrHandlingFSMEvent
}
