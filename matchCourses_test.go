package skiena

import "testing"

func TestMatchCourses(t *testing.T) {
	n := MatchCourses([]string{
		"2ABCD",
		"1YZ",
		"3ABCDEF",
	})
	if n != "ABYCDE" {
		t.Errorf("Funny courses number: %s", n)
	}
}

func TestMatchCoursesNoMatch(t *testing.T) {
	n := MatchCourses([]string{
		"1ABC",
		"1ABC",
		"1ABC",
		"1ABC",
	})
	if n != "" {
		t.Errorf("Funny courses: %s", n)
	}
}
