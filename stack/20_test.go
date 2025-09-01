package stack

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{s: "))", want: false},
		{s: "()", want: true},
		{s: "()[]{}", want: true},
		{s: "(]", want: false},
		{s: ")", want: false},
	}

	for _, test := range tests {
		if result := isValid(test.s); result != test.want {
			t.Errorf("isValid(%q) = %v; want %v", test.s, result, test.want)
		}
	}
}
