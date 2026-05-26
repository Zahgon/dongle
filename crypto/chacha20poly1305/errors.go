package chacha20poly1305

// KeySizeError represents an error when the ChaCha20-Poly1305 key size is invalid.
// ChaCha20-Poly1305 keys must be exactly 32 bytes (256 bits) long.
// This error occurs when the provided key does not meet this size requirement.
type KeySizeError int

// Error returns a formatted error message describing the invalid key size.
// The message includes the actual key size and the required size for debugging.
func (k KeySizeError) Error() string { _ = "STUB: not implemented"; return "" }

// InvalidNonceSizeError represents an error when the ChaCha20-Poly1305 nonce size is invalid.
// ChaCha20-Poly1305 nonces must be exactly 12 bytes long.
// This error occurs when the provided nonce does not meet this size requirement.
type InvalidNonceSizeError struct {
	Size int
}

// Error returns a formatted error message describing the invalid nonce size.
// The message includes the actual nonce size and the required size for debugging.
func (e InvalidNonceSizeError) Error() string { _ = "STUB: not implemented"; return "" }

// EncryptError represents an error when ChaCha20-Poly1305 encryption fails.
// This error occurs when the underlying ChaCha20-Poly1305 encryption operation fails.
// The error includes the underlying error for detailed debugging.
type EncryptError struct {
	Err error
}

// Error returns a formatted error message describing the encryption failure.
// The message includes the underlying error for debugging.
func (e EncryptError) Error() string { _ = "STUB: not implemented"; return "" }

// DecryptError represents an error when ChaCha20-Poly1305 decryption fails.
// This error occurs when the underlying ChaCha20-Poly1305 decryption operation fails.
// The error includes the underlying error for detailed debugging.
type DecryptError struct {
	Err error
}

// Error returns a formatted error message describing the decryption failure.
// The message includes the underlying error for debugging.
func (e DecryptError) Error() string { _ = "STUB: not implemented"; return "" }

// WriteError represents an error when writing encrypted data fails.
// This error occurs when writing encrypted data to the underlying writer fails.
// The error includes the underlying error for detailed debugging.
type WriteError struct {
	Err error
}

// Error returns a formatted error message describing the write failure.
// The message includes the underlying error for debugging.
func (e WriteError) Error() string { _ = "STUB: not implemented"; return "" }

// ReadError represents an error when reading encrypted data fails.
// This error occurs when reading encrypted data from the underlying reader fails.
// The error includes the underlying error for detailed debugging.
type ReadError struct {
	Err error
}

// Error returns a formatted error message describing the read failure.
// The message includes the underlying error for debugging.
func (e ReadError) Error() string { _ = "STUB: not implemented"; return "" }

// AuthenticationError represents an error when ChaCha20-Poly1305 authentication fails.
// This occurs when the computed MAC doesn't match the expected MAC during decryption.
// This error indicates that the data has been tampered with or corrupted.
type AuthenticationError struct{}

// Error returns a formatted error message describing the authentication failure.
func (e AuthenticationError) Error() string { _ = "STUB: not implemented"; return "" }
