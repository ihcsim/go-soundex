package soundex

import (
	"bytes"
	"strings"
)

const codeLen = 4

var codes = map[string]string{
	"a": "",
	"b": "1",
	"c": "2",
	"d": "3",
	"e": "",
	"f": "1",
	"g": "2",
	"h": "",
	"i": "",
	"j": "2",
	"k": "2",
	"l": "4",
	"m": "5",
	"n": "5",
	"o": "",
	"p": "1",
	"q": "2",
	"r": "6",
	"s": "2",
	"t": "3",
	"u": "",
	"v": "1",
	"w": "",
	"x": "2",
	"y": "",
	"z": "2",
}

func Encode(s string) string {
	var encoded bytes.Buffer
	encoded.WriteByte(s[0])

	// NARA American Soundex rule 3: "This rule also applies to the first
	// letter." The first letter is seeded above, but its code was never
	// recorded, so a second letter sharing that code (e.g. the 'f' in
	// "Pfister") was wrongly emitted as a digit. prevCode tracks the code of
	// the first letter's same-code run so such letters are dedup'd against it.
	prevCode := codes[strings.ToLower(string(s[0]))]

	for i := 1; i < len(s); i++ {
		if encoded.Len() == codeLen {
			break
		}

		previous, current := strings.ToLower(string(s[i-1])), strings.ToLower(string(s[i]))

		var next string
		if i+1 < len(s) {
			next = strings.ToLower(string(s[i+1]))
		}

		if (current == "h" || current == "w") && (codes[previous] == codes[next]) {
			i = i + 1
			continue
		}

		if c, ok := codes[current]; ok && len(c) > 0 {
			// NARA rule 3: skip a letter that shares the code of the first
			// letter's same-code run. Vowels end the run (letters separated by
			// a vowel are coded twice); the first emitted digit ends it too.
			if c == prevCode {
				continue
			}
			encoded.WriteByte(c[0])
			prevCode = ""
		} else if current != "h" && current != "w" {
			// a vowel (a, e, i, o, u, y) breaks the first letter's same-code
			// run; h and w are transparent and do not.
			prevCode = ""
		}

		if codes[current] == codes[next] {
			i = i + 1
			continue
		}
	}

	if encoded.Len() < codeLen {
		padding := strings.Repeat("0", codeLen-encoded.Len())
		encoded.WriteString(padding)
	}

	return strings.ToUpper(encoded.String())
}
