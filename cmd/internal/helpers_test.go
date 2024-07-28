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

func TestMoreThanOneNumberReg(t *testing.T) {
	type TestCase struct {
		f, got bool
	}
	testTable := []TestCase{
		{moreThanOneNumberReg("1......"), false},
		{moreThanOneNumberReg(".2....."), false},
		{moreThanOneNumberReg("..3...."), false},
		{moreThanOneNumberReg("...4..."), false},
		{moreThanOneNumberReg("....5.."), false},
		{moreThanOneNumberReg(".....6."), false},
		{moreThanOneNumberReg("......7"), false},
		{moreThanOneNumberReg("12....."), true},
		{moreThanOneNumberReg("1.3...."), true},
		{moreThanOneNumberReg("1..4..."), true},
		{moreThanOneNumberReg("1.34..."), true},
		{moreThanOneNumberReg("...456."), true},
		{moreThanOneNumberReg("1234567"), true},
		{moreThanOneNumberReg(".2.4.6."), true},
		{moreThanOneNumberReg("1.3.5.7"), true},
	}

	for _, test := range testTable {
		if test.f != test.got {
			t.Errorf("Output %v not equal to expected %v", test.f, test.got)
		}
	}
}
