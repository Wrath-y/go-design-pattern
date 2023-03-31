package decorator

type Decorator struct {
	component Component
}

func (d *Decorator) SetComponent(component Component) {
	d.component = component
}

func (d *Decorator) Operation() string {
	if d.component != nil {
		return d.component.Operation()
	}
	return ""
}
