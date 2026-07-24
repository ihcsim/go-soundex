package soundex

import "testing"

func TestSoundex(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{input: "Robert", expected: "R163"},
		{input: "Rupert", expected: "R163"},
		{input: "Ashcraft", expected: "A261"},
		{input: "Ashcroft", expected: "A261"},
		{input: "Mary", expected: "M600"},
		{input: "Simon", expected: "S550"},
		{input: "Allen", expected: "A450"},
		{input: "Sabestian", expected: "S123"},
		{input: "Laura", expected: "L600"},
		{input: "Tymczak", expected: "T522"},
		{input: "Elizabeth", expected: "E421"},
		{input: "Ivan", expected: "I150"},
		{input: "Jennifer", expected: "J516"},
		{input: "Ash", expected: "A200"},
		{input: "Loretta", expected: "L630"},
		{input: "Burroughs", expected: "B620"},
	}

	for _, test := range tests {
		if actual := Encode(test.input); actual != test.expected {
			t.Errorf("Expected encoding of %q to be %q, but got %q", test.input, test.expected, actual)
		}
	}
}

// TestSoundexFirstLetterRule covers NARA American Soundex rule 3: "This rule
// also applies to the first letter." A second letter sharing the first
// letter's code must be dedup'd against it, e.g. "Pfister" -> "P236" (the 'f'
// shares the 'P' code 1, so it is not emitted as a digit). These cases
// previously produced "P123", "L430", "P116" and "B110" respectively.
func TestSoundexFirstLetterRule(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		// first-letter same-code dedup (NARA rule 3)
		{input: "Pfister", expected: "P236"},
		{input: "Lloyd", expected: "L300"},
		{input: "Pfeiffer", expected: "P160"},
		// "Bvbb": B, v, b, b are all code 1 and adjacent, so only the first
		// letter is retained -> B000 (NARA rule 3).
		{input: "Bvbb", expected: "B000"},
		// controls: the first-letter rule must not alter these
		{input: "Robert", expected: "R163"},
		{input: "Rupert", expected: "R163"},
		{input: "Rubin", expected: "R150"},
		{input: "Ashcraft", expected: "A261"},
		{input: "Ashcroft", expected: "A261"},
		{input: "Tymczak", expected: "T522"},
		{input: "Honeyman", expected: "H555"},
		{input: "Euler", expected: "E460"},
		{input: "Gauss", expected: "G200"},
		{input: "Washington", expected: "W252"},
		{input: "Adams", expected: "A352"},
	}

	for _, test := range tests {
		if actual := Encode(test.input); actual != test.expected {
			t.Errorf("Expected encoding of %q to be %q, but got %q", test.input, test.expected, actual)
		}
	}
}
