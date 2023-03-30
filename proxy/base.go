package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (r *RealSubject) Do() string {
	return "Real Subject is doing something."
}

type Proxy struct {
	realSubject *RealSubject
}

func (p *Proxy) Do() string {
	if p.realSubject == nil {
		p.realSubject = &RealSubject{}
	}
	result := "Proxy is doing something, then...\n"
	result += p.realSubject.Do()
	return result
}
