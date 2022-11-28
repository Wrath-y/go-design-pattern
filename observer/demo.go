package observer

import "strconv"

type Subject interface {
	SetSign(_ string)
	GetSign() string
	AddObserver(_ Observer)
}

type Observer interface {
	GetSign(i int)
	SetSubject(_ Subject)
}

type Teacher struct {
	contentOfTheCourse string
	Students           []Observer
}
type Student struct {
	Teacher Subject
}

func (t *Teacher) SetSign(contentOfTheCourse string) {
	t.contentOfTheCourse = contentOfTheCourse
	println("老师:" + contentOfTheCourse)
	println("老师:我刚刚说了什么你们复述一遍")
	t.NotifyStudents()
}

func (t *Teacher) GetSign() string {
	return t.contentOfTheCourse
}

func (t *Teacher) AddObserver(s Observer) {
	if t.Students == nil {
		t.Students = make([]Observer, 0)
	}
	t.Students = append(t.Students, s)
}

func (t *Teacher) NotifyStudents() {
	for i, stu := range t.Students {
		stu.GetSign(i)
	}
}

func (s *Student) GetSign(i int) {
	println("学生" + strconv.Itoa(i) + ":" + s.Teacher.GetSign())
}

func (s *Student) SetSubject(t Subject) {
	s.Teacher = t
	s.Teacher.AddObserver(s)
}
