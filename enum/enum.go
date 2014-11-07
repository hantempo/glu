package enum

import "fmt"

const (
	GL_BYTE                    = 0x00001400
	GL_UNSIGNED_BYTE           = 0x00001401
	GL_SHORT                   = 0x00001402
	GL_UNSIGNED_SHORT          = 0x00001403
	GL_INT                     = 0x00001404
	GL_UNSIGNED_INT            = 0x00001405
	GL_FLOAT                   = 0x00001406
	GL_UNSIGNED_BYTE_3_3_2     = 0x00008032
	GL_UNSIGNED_SHORT_4_4_4_4  = 0x00008033
	GL_UNSIGNED_SHORT_5_5_5_1  = 0x00008034
	GL_UNSIGNED_INT_8_8_8_8    = 0x00008035
	GL_UNSIGNED_INT_10_10_10_2 = 0x00008036
	GL_RED                     = 0x00001903
	GL_GREEN                   = 0x00001904
	GL_BLUE                    = 0x00001905
	GL_ALPHA                   = 0x00001906
	GL_RGB                     = 0x00001907
	GL_RGBA                    = 0x00001908
	GL_LUMINANCE               = 0x00001909
	GL_LUMINANCE_ALPHA         = 0x0000190A
)

var typeStrings = map[uint32]string{
	GL_BYTE:                    "GL_BYTE",
	GL_UNSIGNED_BYTE:           "GL_UNSIGNED_BYTE",
	GL_SHORT:                   "GL_SHORT",
	GL_UNSIGNED_SHORT:          "GL_UNSIGNED_SHORT",
	GL_INT:                     "GL_INT",
	GL_UNSIGNED_INT:            "GL_UNSIGNED_INT",
	GL_FLOAT:                   "GL_FLOAT",
	GL_UNSIGNED_BYTE_3_3_2:     "GL_UNSIGNED_BYTE_3_3_2",
	GL_UNSIGNED_SHORT_4_4_4_4:  "GL_UNSIGNED_SHORT_4_4_4_4",
	GL_UNSIGNED_SHORT_5_5_5_1:  "GL_UNSIGNED_SHORT_5_5_5_1",
	GL_UNSIGNED_INT_8_8_8_8:    "GL_UNSIGNED_INT",
	GL_UNSIGNED_INT_10_10_10_2: "GL_UNSIGNED_INT_10_10_10_2",
}

func TypeString(e uint32) string {
	if s, ok := typeStrings[e]; ok {
		return s
	} else {
		return fmt.Sprintf("Invalid type(0x%X)", e)
	}
}

var formatStrings = map[uint32]string{
	GL_RED:             "GL_RED",
	GL_GREEN:           "GL_GREEN",
	GL_BLUE:            "GL_BLUE",
	GL_ALPHA:           "GL_ALPHA",
	GL_RGB:             "GL_RGB",
	GL_RGBA:            "GL_RGBA",
	GL_LUMINANCE:       "GL_LUMINANCE",
	GL_LUMINANCE_ALPHA: "GL_LUMINANCE_ALPHA",
}

func FormatString(e uint32) string {
	if s, ok := formatStrings[e]; ok {
		return s
	} else {
		return fmt.Sprintf("Invalid format(0x%X)", e)
	}
}
