package slugify

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
	"github.com/digitalxero/slugify/transliterations"

)

var SKIP = []*unicode.RangeTable{
	unicode.Mark,
	unicode.Sk,
	unicode.Lm,
}

var SAFE = []*unicode.RangeTable{
	unicode.Letter,
	unicode.Number,
}

var DASH = []*unicode.RangeTable{
	unicode.Pd,
}

var SPACE = []*unicode.RangeTable{
	unicode.Space,
}

var OK = "-_"
var TO_DASH = "/\\—–.~!@#$%^&*(){}[]+=?><;:`"
var extra_dashes = regexp.MustCompile("[-]{2,}")

// Slugify a string. The result will only contain lowercase letters,
// digits and dashes. It will not begin or end with a dash, and it
// will not contain runs of multiple dashes.
//
// It is NOT forced into being ASCII, but may contain any Unicode
// characters, with the above restrictions.
func Slugify(text string) string {
	buf := make([]rune, 0, len(text))
	text = SanatizeText(text)
	for _, r := range norm.NFKD.String(text) {
		s := strconv.QuoteRune(r)
		switch {
		case unicode.IsOneOf(SAFE, r):
			buf = append(buf, unicode.ToLower(r))
		case strings.ContainsAny(s, OK):
			buf = append(buf, r)
		case unicode.IsOneOf(SPACE, r):
			buf = append(buf, '-')
		case unicode.IsOneOf(DASH, r):
			buf = append(buf, '-')
		case strings.ContainsAny(s, TO_DASH):
			buf = append(buf, '-')
		}
	}
	return cleanup(string(buf))
}

// IDify a string. The result will only contain ASCII letters,
// digits and dashes. It will not begin or end with a dash, and it
// will not contain runs of multiple dashes.
//
// It is forced into being ASCII, but may contain any Unicode
// characters, with the above restrictions.
func IDify(text string) string {
	buf := make([]rune, 0, len(text))
	text = SanatizeText(text)
	for _, r := range norm.NFKD.String(text) {
		s := strconv.QuoteRune(r)
		switch {
		case unicode.IsOneOf(SAFE, r):
			buf = append(buf, r)
		case strings.ContainsAny(s, OK):
			buf = append(buf, r)
		case unicode.IsOneOf(SPACE, r):
			buf = append(buf, '-')
		case unicode.IsOneOf(DASH, r):
			buf = append(buf, '-')
		case strings.ContainsAny(s, TO_DASH):
			buf = append(buf, '-')
		}
	}
	return cleanup(string(buf))
}

func SanatizeText(text string) string {
	b := bytes.NewBufferString("")
	for _, c := range text {
		b.WriteString(transliterations.Transliterate(c))
	}
	return b.String()
}

func cleanup(text string) string {
	text = strings.Trim(text, "-")
	text = extra_dashes.ReplaceAllString(text, "-")
	return text
}
