package tea

// KeySizeError represents an error when the TEA key size is invalid.
// TEA keys must be exactly 16 bytes (128 bits) long.
// This error occurs when the provided key does not meet this size requirement.
type KeySizeError int

// Error returns a formatted error message describing the invalid key size.
// The message includes the actual key size and the required size for debugging.
func (k KeySizeError) Error() string { _ = "STUB: not implemented"; return "" }

// EncryptError represents an error when TEA encryption fails.
// This error occurs when the underlying TEA encryption operation fails.
// The error includes the underlying error for detailed debugging.
type EncryptError struct {
	Err error // The underlying error that caused the failure
}

// Error returns a formatted error message describing the encryption failure.
// The message includes the underlying error for debugging.
func (e EncryptError) Error() string { _ = "STUB: not implemented"; return "" }

// DecryptError represents an error when TEA decryption fails.
// This error occurs when the underlying TEA decryption operation fails.
// The error includes the underlying error for detailed debugging.
type DecryptError struct {
	Err error // The underlying error that caused the failure
}

// Error returns a formatted error message describing the decryption failure.
// The message includes the underlying error for debugging.
func (e DecryptError) Error() string { _ = "STUB: not implemented"; return "" }

// WriteError represents an error when writing encrypted data fails.
// This error occurs when writing encrypted data to the underlying writer fails.
// The error includes the underlying error for detailed debugging.
type WriteError struct {
	Err error // The underlying error that caused the failure
}

// Error returns a formatted error message describing the write failure.
// The message includes the underlying error for debugging.
func (e WriteError) Error() string { _ = "STUB: not implemented"; return "" }

// ReadError represents an error when reading encrypted data fails.
// This error occurs when reading encrypted data from the underlying reader fails.
// The error includes the underlying error for detailed debugging.
type ReadError struct {
	Err error // The underlying error that caused the failure
}

// Error returns a formatted error message describing the read failure.
// The message includes the underlying error for debugging.
func (e ReadError) Error() string { _ = "STUB: not implemented"; return "" }

// InvalidDataSizeError represents an error when the data size is invalid for TEA operations.
// TEA requires data to be a multiple of 8 bytes (64 bits).
type InvalidDataSizeError struct {
	Size int // The actual data size that caused the error
}

// Error returns a formatted error message describing the invalid data size.
// The message includes the actual size and the required size for debugging.
func (e InvalidDataSizeError) Error() string { _ = "STUB: not implemented"; return "" }

// UnsupportedBlockModeError represents an error when an unsupported block mode is used.
type UnsupportedBlockModeError struct {
	Mode string // The unsupported mode name
}

// Error returns a formatted error message describing the unsupported mode.
// The message includes the mode name and explains why it's not supported.
func (e UnsupportedBlockModeError) Error() string { _ = "STUB: not implemented"; return "" }
