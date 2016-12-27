package suddendeath

import (
	"bytes"
	"strings"

	"golang.org/x/text/width"
)

const (
	headerLeft  = "＿人"
	headerRight = "人＿"
	header      = "人"
	bodyLeft    = "＞　"
	bodyRight   = "　＜"
	footerLeft  = "￣Y"
	footerRight = "^Y￣"
	footer      = "^Y"
)

// Generate generates a formated text.
func Generate(text string) string {
	// remove line breaks
	text = strings.Replace(text, "\n", "", -1)

	width := textWidth(text)

	var buf bytes.Buffer

	// header
	buf.WriteString(headerLeft)
	for i := 0; i < (width+1)/2; i++ {
		buf.WriteString(header)
	}
	buf.WriteString(headerRight)
	buf.WriteString("\n")

	// body
	buf.WriteString(bodyLeft)
	buf.WriteString(text)
	if width%2 != 0 {
		buf.WriteString(" ")
	}
	buf.WriteString(bodyRight)
	buf.WriteString("\n")

	// footer
	buf.WriteString(footerLeft)
	for i := 0; i < (width+1)/2; i++ {
		buf.WriteString(footer)
	}
	buf.WriteString(footerRight)

	return buf.String()
}

func textWidth(text string) int {
	w := 0

	for _, r := range text {
		p := width.LookupRune(r)
		switch p.Kind() {
		case width.EastAsianFullwidth, width.EastAsianWide, width.EastAsianAmbiguous:
			w += 2
		default:
			w += 1
		}
	}

	return w
}
