package School

type Teacher struct {
	Name         string
	CourseTaught string
}

//Grade given to a teacher of a particular course to a particular student
func (grade *Teacher) Marks(studentName string, studentGrade int, passMark int, M map[string]Student) string {
	var status = "Failed"
	for index, value := range M {
		for i := 0; i < len(value.CoursesOffered); i++ {
			if index == studentName && value.CoursesOffered[i] == grade.CourseTaught {
				if studentGrade >= passMark {
					status = "Passed"
				}
				Students[index] = Student{Name: index, Grades: value.Grades + studentGrade, CoursesOffered: value.CoursesOffered}
			}
		}
	}
	return status
}
