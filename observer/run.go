package observer

func Run() {
	teacher := new(Teacher)

	student1 := new(Student)
	student1.SetSubject(teacher)
	student2 := new(Student)
	student2.SetSubject(teacher)
	student3 := new(Student)
	student3.SetSubject(teacher)

	teacher.SetSign("体育老师请假了，这节课我来上")
}
