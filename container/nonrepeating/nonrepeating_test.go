package main

import "testing"

func TestNonrepeating(t *testing.T) {
	tests := []struct {
		s      string
		answer int
	}{
		// Normal cases
		{"abcabcbb", 3},
		{"pwwkew", 3},
		// Edge cases
		{"", 0},
		{"b", 1},
		{"bbbbbbbbbbb", 1},
		{"abcabcabcd", 4},
		// Chinese
		{"红鲤鱼与绿鲤鱼与驴", 5},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		actual := LengthOfNonRepeatingSubStr(tt.s)
		if actual != tt.answer {
			t.Errorf("got %d for input %s; exoected %d", actual, tt.s, tt.answer)
		}
	}
}

func BenchmarkLengthOfNonRepeatingSubStr(b *testing.B) {
	s := "红鲤鱼与绿鲤鱼与驴"
	for i := 0; i < 13; i++ {
		s += s
	}
	b.Logf("%d", len(s))
	b.ResetTimer()

	answer := 6
	actual := LengthOfNonRepeatingSubStr(s)
	for i := 0; i < b.N; i++ {
		if actual != answer {
			b.Errorf("got %d for input %s; exoected %d", actual, s, answer)
		}
	}

}
