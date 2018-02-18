package gopractice

import "testing"

func TestSortString(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"", ""},
		{"abcdef", "abcdef"},
		{"cbeadf", "abcdef"},
	}
	for _, c := range cases {
		got := SortString(c.in)
		if got != c.want {
			t.Errorf("SortString(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestContainsString(t *testing.T) {
	cases := []struct {
		in1, in2 string
		want     bool
	}{
		{"this is a string", "ing", true},
		{"this is a string", "this ", true},
		{"this is a string", "is a", true},
		{"", "is a", false},
		{"good good", "bad", false},
		{"bad", "good", false},
		{"", "", true},
	}
	for _, c := range cases {
		got := ContainsString(c.in1, c.in2)
		if got != c.want {
			t.Errorf("ContainsString(%q, %q) == %q, want %t", c.in1, c.in2, got, c.want)
		}
	}
}

func TestIsAnagram(t *testing.T) {
	cases := []struct {
		in1, in2 string
		want     bool
	}{
		{"arroz", "zorra", true},
		{"arroz", "zorras", false},
		{"dude", "crap", false},
		{"", "", true},
	}
	for _, c := range cases {
		got := IsAnagram(c.in1, c.in2)
		if got != c.want {
			t.Errorf("IsAnagram(%q, %q) == %q, want %t", c.in1, c.in2, got, c.want)
		}
	}
}
