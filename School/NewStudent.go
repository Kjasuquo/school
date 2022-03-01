package School

var Prospect []string

type Applicants struct {
	Name   string
	Scores [3]int
	Age    int
}

//Admission of an Applicant to student
func (nStud *Applicants) NewStudent(passMark int, minAge int) string {
	var result string = nStud.Name + " is Not Admitted"
	var x int
	for i := 0; i < len(nStud.Scores); i++ {
		x += nStud.Scores[i]
	}
	if x/3 >= passMark && nStud.Age <= minAge {
		result = nStud.Name + " is Admitted"
		Prospect = append(Prospect, nStud.Name)
	}
	return result
}
