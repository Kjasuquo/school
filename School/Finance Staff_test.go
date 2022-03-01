package School

import (
	"testing"
)

func TestStaff_SchoolFees(t *testing.T) {
	var tests = []struct {
		input0 Staff
		input1 int

		want string
	}{
		{Staff{StudentName: "Alex", FeesPaid: 39000}, 40000, "Fees Not Complete"},
		{Staff{StudentName: "Blex", FeesPaid: 40000}, 40000, "Completed"},
	}

	for _, test := range tests {
		got := test.input0.SchoolFees(test.input1)
		if got != test.want {
			t.Errorf("expected %v but got %v", test.want, got)
		}
	}
}
