package School

import "testing"

func TestPrincipal_AdmitStudent(t *testing.T) {
	var tests = []struct {
		input0 Principal
		input1 string
		input2 []string
		want   string
	}{
		{Principal{}, "Oreva", []string{"Alex", "Oreva", "Chuks"}, "Oreva Is Admitted"},
	}
	for _, test := range tests {
		got := test.input0.AdmitStudent(test.input1, test.input2)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}

func TestPrincipal_ExpelStudent(t *testing.T) {
	m := make(map[string]Student)
	m["Alex"] = Student{"Alex", 115, []string{"English", "Math", "Java"}, "Completed"}
	m["Blex"] = Student{"Blex", 137, []string{"PHE", "Java", "English"}, "Not Completed"}
	m["Dlex"] = Student{"Dlex", 70, []string{"Math", "Java", "English"}, ""}
	var tests = []struct {
		input0 Principal
		input1 int
		input2 map[string]Student

		want string
	}{
		{Principal{}, 30, m, "Dlex Expelled"},
	}
	for _, test := range tests {
		got := test.input0.ExpelStudent(test.input1, test.input2)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
