package School

var Students = make(map[string]Student)

type Student struct {
	Name           string
	Grades         int
	CoursesOffered []string
	Fees           string
}

// CourseEnlisted is function of courses offered by a student
func (c *Student) CourseEnlisted(name string, courses []string) string {
	var status = "Not A Student"
	if name == c.Name {
		Students[c.Name] = Student{Name: c.Name, CoursesOffered: courses}
		status = "Successful Course Registration"
	}
	return status
}
