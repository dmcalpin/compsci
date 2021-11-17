package behavioral

import (
	"fmt"
)

// The observer pattern allows multiple
// objects to observer changes in a
// different object.

// Subject is an object whos data, when
// chaged, notifies it's Observers
type Subject interface {
	AddObserver(Observer)
	GetState() interface{}
	SetState(interface{})
	notifyObservers()
}

// Observer maintains a reference to the
// subject, and will have it's Update
// function called when the Subject changes
type Observer interface {
	Update()
	GetSubject() Subject
	SetSubject(Subject)
}

type NumberSubject struct {
	num       interface{}
	observers []Observer
}

func (s *NumberSubject) AddObserver(o Observer) {
	o.SetSubject(s)
	s.observers = append(s.observers, o)
}

func (s *NumberSubject) GetState() interface{} {
	return s.num
}

func (s *NumberSubject) SetState(num interface{}) {
	s.num = num
	s.notifyObservers()
}

func (s *NumberSubject) notifyObservers() {
	for _, o := range s.observers {
		o.Update()
	}
}

type PrintObserver struct {
	updatedVal string
	subject    Subject
}

func (o *PrintObserver) GetSubject() Subject {
	return o.subject
}

func (o *PrintObserver) SetSubject(s Subject) {
	o.subject = s
}

func (o *PrintObserver) Update() {
	o.updatedVal = fmt.Sprintln("Print Observer: ", o.GetSubject().GetState())
}

type RepeatingPrintObserver struct {
	PrintObserver
}

func (o *RepeatingPrintObserver) Update() {
	state := o.GetSubject().GetState()

	o.updatedVal = fmt.Sprintln("Repeating Print Observer: ", state, state, state)
}
