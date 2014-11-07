package enum

import "testing"

var tests = []struct {
	etos func(uint) string
	e    uint
	s    string
}{
	{TypeString, GL_UNSIGNED_INT, "GL_UNSIGNED_INT"},
	{TypeString, GL_UNSIGNED_SHORT_4_4_4_4, "GL_UNSIGNED_SHORT_4_4_4_4"},
	{TypeString, GL_RGBA, "Invalid type"},

	{FormatString, GL_RGBA, "GL_RGBA"},
	{FormatString, GL_UNSIGNED_INT, "Invalid format"},
}

func TestEnum(t *testing.T) {
	for _, test := range tests {
		s := test.etos(test.e)
		if s != test.s {
			t.Errorf("Expected (%s), got (%s)", test.s, s)
		}
	}
}
