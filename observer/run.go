package observer

func Run() {
	subject := &Subject{}

	observerA := &ConcreteObserverA{}
	subject.Attach(observerA)

	observerB := &ConcreteObserverB{}
	subject.Attach(observerB)

	subject.Notify()
}
