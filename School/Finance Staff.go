package School

type Staff struct {
	StudentName string
	FeesPaid    int
}

func (s *Staff) SchoolFees(schoolFees int) string {
	var fees string

	if schoolFees > s.FeesPaid {
		fees = "Fees Not Complete"
	} else {
		fees = "Completed"
	}

	Students[s.StudentName] = Student{Name: s.StudentName, Fees: fees}
	return fees
}
