package enum

import "testing"

var tests = []struct {
	etos func(uint32) string
	e    uint32
	s    string
}{
	{TypeString, GL_NONE, "GL_NONE"},
	{TypeString, GL_UNSIGNED_INT, "GL_UNSIGNED_INT"},
	{TypeString, GL_UNSIGNED_SHORT_4_4_4_4, "GL_UNSIGNED_SHORT_4_4_4_4"},
	{TypeString, GL_HALF_FLOAT, "GL_HALF_FLOAT"},
	{TypeString, GL_RGBA, "Invalid type(0x1908)"},

	{FormatString, GL_NONE, "GL_NONE"},
	{FormatString, GL_RGBA, "GL_RGBA"},
	{FormatString, GL_UNSIGNED_INT, "Invalid format(0x1405)"},

	{FormatString, GL_ETC1_RGB8_OES, "GL_ETC1_RGB8_OES"},
}

func TestEnum(t *testing.T) {
	for _, test := range tests {
		s := test.etos(test.e)
		if s != test.s {
			t.Errorf("Expected (%s), got (%s)", test.s, s)
		}
	}
}
