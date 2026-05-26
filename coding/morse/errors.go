package morse

// InvalidInputError represents an error when the morse input is invalid.
// This error is now rarely used since most characters are supported.
type InvalidInputError struct {
	Char string // The invalid character that was found
}

// Error returns a formatted error message describing the invalid input.
func (e InvalidInputError) Error() string { _ = "STUB: not implemented"; return "" }

// InvalidCharacterError represents an error when an invalid morse character is found
// during decoding. This error occurs when a morse code sequence is not recognized.
type InvalidCharacterError struct {
	Char string // The invalid morse character that was found
}

// Error returns a formatted error message describing the invalid character.
func (e InvalidCharacterError) Error() string { _ = "STUB: not implemented"; return "" }
