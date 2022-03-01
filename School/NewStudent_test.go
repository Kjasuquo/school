package School

import "testing"

//var applicant1 = Applicants{Name: "Joseph", Scores: [3]int{90, 100, 99}, Age: 17}

func TestApplicants_NewStudent(t *testing.T) {
	var tests = []struct {
		input0 Applicants
		input1 int
		input2 int

		want string
	}{
		{Applicants{Name: "Joseph", Scores: [3]int{90, 100, 99}, Age: 17}, 50, 18, "Joseph is Admitted"},
		{Applicants{Name: "Clinton", Scores: [3]int{8, 4, 99}, Age: 17}, 50, 18, "Clinton is Not Admitted"},
	}

	for _, test := range tests {
		got := test.input0.NewStudent(test.input1, test.input2)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}

	}
}
