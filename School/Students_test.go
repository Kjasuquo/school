package School

import "testing"

func TestStudent_CourseEnlisted(t *testing.T) {
	var tests = []struct {
		input0 Student
		input1 string
		input2 []string

		want string
	}{
		{Student{Name: "Alex", Grades: 0, CoursesOffered: []string{}, Fees: ""}, "Alex", []string{"Chemistry", "Java", "English"}, "Successful Course Registration"},
	}

	for _, test := range tests {
		got := test.input0.CourseEnlisted(test.input1, test.input2)
		if got != test.want {
			t.Errorf("Expected %v but got %v", test.want, got)
		}
	}
}
