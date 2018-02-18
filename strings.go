package gopractice

import "sort"

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func ContainsString(s, substr string) bool {
	if len(s) < len(substr) {
		return false
	}

	s_a := []rune(s)
	substr_a := []rune(substr)

	for i := 0; i <= len(s_a)-len(substr_a); i++ {
		j := 0
		for ; j < len(substr_a) && s_a[i+j] == substr_a[j]; j++ {
		}

		if j == len(substr_a) {
			return true
		}
	}

	return false
}

func IsAnagram(in1 string, in2 string) bool {
	if len(in1) != len(in2) {
		return false
	}

	return SortString(in1) == SortString(in2)
}
