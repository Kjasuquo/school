package School

type Principal struct {
}

//For admitting prospective students
func (admit *Principal) AdmitStudent(Name string, P []string) string {
	for _, studentName := range P {
		if studentName == Name {
			Students[studentName] = Student{Name: studentName}
		}
	}
	return Name + " Is Admitted"
}

//Expelling from the Principal
func (expel *Principal) ExpelStudent(cutoff int, M map[string]Student) string {
	var x = "Nobody Is Expelled"
	for index, value := range M {
		if value.Grades/3 < cutoff {
			x = value.Name + " Expelled"
			delete(M, index)
		}
	}
	return x
}
