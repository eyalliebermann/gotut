package stringutil

import "testing"

// func TestFailue(t *testing.T) {
// 	t.Errorf("go test excuted this one and successfully failed")
// }

func TestEmptyRemainsTheSame(t *testing.T) {
	in := ""
	want := ""
	got := ToUpper(in)
	if want != got {
		t.Errorf("Empty string should result in empty string. but got %v", got)
	}
}

func TestListOfExcutionValues(t *testing.T) {

	for _, c := range []struct {
		in, want string
	}{
		{"", ""},
		{"a", "A"},
		{"B", "B"},
		{"ä", "Ä"},
		{"א", "א"},
		{"0", "0"},
		{"abc", "ABC"},
		{"ABCD", "ABCD"},
		{"AbcDE", "ABCDE"},
		{"qWERty123!", "QWERTY123!"},
	} {
		got := ToUpper(c.in)
		if got != c.want {
			t.Errorf("In=%v, want=%v, got=%v", c.in, c.want, got)
		}
	}
}
