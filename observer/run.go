package observer

// Run
// 在示例中，我们定义了一个 Subject 结构体和一个 Observer 接口。具体观察者类（例如 ConcreteObserverA 和 ConcreteObserverB）实现了 Update() 方法，并注册到被观察者对象中。
// 当被观察者状态发生改变时，它将调用自己的 Notify() 方法来通知所有注册过的观察者对象。每个观察者对象都实现了 Update() 方法，并提供了自己的更新方法。
func Run() {
	subject := &Subject{}

	observerA := &ConcreteObserverA{}
	subject.Attach(observerA)

	observerB := &ConcreteObserverB{}
	subject.Attach(observerB)

	subject.Notify()
}
