package observer

import "fmt"

type Observer interface {
	Update()
}

type Subject struct {
	observers []Observer
}

func (s *Subject) Attach(o Observer) {
	s.observers = append(s.observers, o)
}

func (s *Subject) Detach(o Observer) {
	for i, observer := range s.observers {
		if observer == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *Subject) Notify() {
	for _, observer := range s.observers {
		observer.Update()
	}
}

type ConcreteObserverA struct{}

func (o *ConcreteObserverA) Update() {
	fmt.Println("ConcreteObserverA has been notified")
}

type ConcreteObserverB struct{}

func (o *ConcreteObserverB) Update() {
	fmt.Println("ConcreteObserverB has been notified")
}
