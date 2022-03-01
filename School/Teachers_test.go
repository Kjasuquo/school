package School

import "testing"

func TestTeacher_Marks(t *testing.T) {
	m := make(map[string]Student)
	m["Alex"] = Student{"Alex", 0, []string{"English", "Math", "Java"}, "Completed"}
	m["Blex"] = Student{"Blex", 0, []string{"PHE", "Java", "English"}, "Not Completed"}
	m["Dlex"] = Student{"Dlex", 0, []string{"Math", "Java", "English"}, ""}

	var tests = []struct {
		input0 Teacher
		input1 string
		input2 int
		input3 int
		input4 map[string]Student

		want string
	}{
		{Teacher{"Gbenga", "Java"}, "Blex", 80, 50, m, "Passed"},
	}
	for _, test := range tests {
		got := test.input0.Marks(test.input1, test.input2, test.input3, test.input4)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
