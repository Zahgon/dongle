package sm4

// KeySizeError represents an error when the SM4 key size is invalid.
// SM4 keys must be exactly 16 bytes (128 bits).
type KeySizeError int

// Error returns the error message for KeySizeError.
func (k KeySizeError) Error() string { _ = "STUB: not implemented"; return "" }

// EncryptError represents an error during SM4 encryption.
type EncryptError struct {
	Err error
}

// Error returns the error message for EncryptError.
func (e EncryptError) Error() string { _ = "STUB: not implemented"; return "" }

// DecryptError represents an error during SM4 decryption.
type DecryptError struct {
	Err error
}

// Error returns the error message for DecryptError.
func (d DecryptError) Error() string { _ = "STUB: not implemented"; return "" }

// ReadError represents an error during data reading in streaming operations.
type ReadError struct {
	Err error
}

// Error returns the error message for ReadError.
func (r ReadError) Error() string { _ = "STUB: not implemented"; return "" }
