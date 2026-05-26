package mock

// ErrorHasher is a mock implementation of hash.Hash that can return errors on Write operations.
// This is useful for testing error handling in code that uses hash.Hash interfaces.
type ErrorHasher struct {
	writeErr error // Error to return from Write method
}

// NewErrorHasher creates a new ErrorHasher that will return the specified error
// when Write() is called. This is useful for testing hash write error scenarios.
func NewErrorHasher(writeErr error) *ErrorHasher { _ = "STUB: not implemented"; return nil }

// Write implements the hash.Hash interface and returns the configured error.
// This simulates a hash write failure for testing purposes.
func (h *ErrorHasher) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Sum implements the hash.Hash interface and returns a mock hash value.
// This always succeeds and returns a unique mock hash for testing.
func (h *ErrorHasher) Sum(b []byte) []byte {
	_ = "STUB: not implemented"
	// Return unique hash based on writeErr to satisfy HMAC requirements
	return nil
}

// Use error message hash to make it unique

// Reset implements the hash.Hash interface but does nothing in this mock.
func (h *ErrorHasher) Reset() {
	_ = "STUB: not implemented"
	// This is a no-op method, but we add a simple operation
	// to ensure proper coverage tracking
	return
}

// Size implements the hash.Hash interface and returns a mock hash size.
func (h *ErrorHasher) Size() int {
	_ = "STUB: not implemented"

	// BlockSize implements the hash.Hash interface and returns a mock block size.
	return 0
}

func (h *ErrorHasher) BlockSize() int { _ = "STUB: not implemented"; return 0 }
