package internal

import "testing"

func TestSSIMtoDate(t *testing.T) {
	type TestCase struct {
		f, got string
	}

	testTable := []TestCase{
		{SSIMtoDate("13JUL24"), "2024-07-13"},
		{SSIMtoDate("6MAR24"), "2024-03-06"},
		{SSIMtoDate("1JAN25"), "2025-01-01"},
	}

	for _, test := range testTable {
		if test.f != test.got {
			t.Errorf("Output %q not equal to expected %q", test.f, test.got)
		}
	}

}
