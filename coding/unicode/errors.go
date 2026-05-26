package unicode

// DecodeFailedError represents an error when unicode decoding fails.
// This error occurs when invalid unicode escape sequences are encountered
// during decoding operations.
type DecodeFailedError struct {
	Input string // The invalid input that caused the error
}

// Error returns a formatted error message describing the decode failure.
func (e DecodeFailedError) Error() string { _ = "STUB: not implemented"; return "" }

// InvalidUnicodeError represents an error when invalid unicode data is encountered.
// This error occurs when malformed unicode escape sequences are found.
type InvalidUnicodeError struct {
	Char string // The invalid unicode character that was found
}

// Error returns a formatted error message describing the invalid unicode.
func (e InvalidUnicodeError) Error() string { _ = "STUB: not implemented"; return "" }

// EncodeFailedError represents an error when unicode encoding fails.
// This error is rarely used since strconv.QuoteToASCII rarely fails.
type EncodeFailedError struct {
	Input string // The input that failed to encode
}

// Error returns a formatted error message describing the encode failure.
func (e EncodeFailedError) Error() string { _ = "STUB: not implemented"; return "" }
