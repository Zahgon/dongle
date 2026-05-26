package mock

import (
	"io"
	"io/fs"
	"os"
	"time"
)

// File is a mock implementation of the fs.File interface for testing purposes.
// It provides an in-memory file representation that can be used to simulate
// file operations without touching the actual file system.
type File struct {
	data   []byte // File content stored in memory
	pos    int64  // Current read/write position
	closed bool   // Whether the file has been closed
	name   string // File name for identification
}

// NewFile creates a new mock file with the specified data and name.
// This function is commonly used in tests to create file-like objects
// that can be passed to functions expecting file interfaces.
func NewFile(data []byte, name string) *File { _ = "STUB: not implemented"; return nil }

// Read implements the io.Reader interface for mock file operations.
// It reads data from the current position and advances the position accordingly.
// Returns os.ErrClosed if the file has been closed, or io.EOF when reaching the end.
func (f *File) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements the io.Closer interface for mock file operations.
// Marks the file as closed, preventing further read operations.
func (f *File) Close() error { _ = "STUB: not implemented"; return nil }

// Stat returns file information for the mock file.
// Creates a fileInfo object with the file's name and size based on data length.
func (f *File) Stat() (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}

// ReadDir implements the fs.ReadDirFile interface.
// Since this mock represents a regular file, not a directory, it always returns an error.
func (f *File) ReadDir(count int) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Seek implements the io.Seeker interface for mock file operations.
// Allows positioning the file pointer at different locations within the file.
// Supports seeking from start, current position, or end of file.
func (f *File) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Write implements the io.Writer interface for mock file operations.
// It writes data to the current position and advances the position accordingly.
// If writing beyond the current data length, the file is extended.
// Returns os.ErrClosed if the file has been closed.
func (f *File) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// If writing beyond current data length, extend the data slice

// Write data at current position

// If we wrote at position 0 and the new data is shorter than the original,
// truncate the data to avoid keeping old content

// Bytes returns the current file content (for testing)
func (f *File) Bytes() []byte {
	_ = "STUB: not implemented"

	// Reset resets the file position to the beginning
	return nil
}

func (f *File) Reset() {
	_ = "STUB: not implemented"

	// Truncate truncates the file to the specified size
	return
}

func (f *File) Truncate(n int) { _ = "STUB: not implemented"; return }

// ErrorFile is a mock file implementation that always returns errors.
// This is useful for testing error handling paths in code that operates on files.
type ErrorFile struct {
	err error // The error to return for all operations
}

// NewErrorFile creates a new error file that will return the specified error
// for all file operations. This is commonly used to test error scenarios.
func NewErrorFile(err error) *ErrorFile { _ = "STUB: not implemented"; return nil }

// Read always returns the configured error, simulating a file read failure.
func (e *ErrorFile) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"

	// Close always returns the configured error, simulating a file close failure.
	return 0, nil
}

func (e *ErrorFile) Close() error {
	_ = "STUB: not implemented"

	// Stat always returns the configured error, simulating a file stat failure.
	return nil
}

func (e *ErrorFile) Stat() (os.FileInfo, error) {
	_ = "STUB: not implemented"

	// ReadDir always returns the configured error, simulating a directory read failure.
	return *new(os.FileInfo), nil
}

func (e *ErrorFile) ReadDir(count int) ([]fs.DirEntry, error) {
	_ = "STUB: not implemented"

	// Seek always returns the configured error, simulating a file seek failure.
	return nil, nil
}

func (e *ErrorFile) Seek(offset int64, whence int) (int64, error) {
	_ = "STUB: not implemented"

	// Write always returns the configured error, simulating a file write failure.
	return 0, nil
}

func (e *ErrorFile) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// fileInfo implements the os.FileInfo interface for mock file information.
	// Provides basic file metadata for mock files used in testing.
	return 0, nil
}

type fileInfo struct {
	name string // File name
	size int64  // File size in bytes
}

// Name returns the file name.
func (f *fileInfo) Name() string {
	_ = "STUB: not implemented"

	// Size returns the file size in bytes.
	return ""
}

func (f *fileInfo) Size() int64 {
	_ = "STUB: not implemented"

	// Mode returns a read-only file mode (0444) for mock files.
	return 0
}

func (f *fileInfo) Mode() os.FileMode {
	_ = "STUB: not implemented"

	// ModTime returns a zero time value for mock files.
	return *new(os.FileMode)
}

func (f *fileInfo) ModTime() time.Time {
	_ = "STUB: not implemented"

	// IsDir returns false since mock files represent regular files, not directories.
	return *new(time.Time)
}

func (f *fileInfo) IsDir() bool {
	_ = "STUB: not implemented"

	// Sys returns nil for mock files as they don't have underlying system-specific data.
	return false
}

func (f *fileInfo) Sys() interface{} {
	_ = "STUB: not implemented"

	// WriteCloser is a mock implementation of io.WriteCloser for testing purposes.
	// It wraps an io.Writer and adds close functionality with state tracking.
	return nil
}

type WriteCloser struct {
	w      io.Writer // Underlying writer to delegate writes to
	closed bool      // Whether the WriteCloser has been closed
}

// NewWriteCloser creates a new mock WriteCloser that wraps the provided io.Writer.
// This is useful for testing code that requires both write and close operations.
func NewWriteCloser(w io.Writer) *WriteCloser { _ = "STUB: not implemented"; return nil }

// Write implements the io.Writer interface by delegating to the underlying writer.
// Returns os.ErrClosed if the WriteCloser has been closed.
func (w *WriteCloser) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Close implements the io.Closer interface for the mock WriteCloser.
// Marks the WriteCloser as closed and prevents further write operations.
func (w *WriteCloser) Close() error { _ = "STUB: not implemented"; return nil }

// ErrorWriteCloser is a mock io.WriteCloser that always returns errors.
// Useful for testing error handling in code that writes to files or other writers.
type ErrorWriteCloser struct {
	err error // The error to return for all operations
}

// NewErrorWriteCloser creates a new error WriteCloser that will return
// the specified error for all write and close operations.
func NewErrorWriteCloser(err error) *ErrorWriteCloser { _ = "STUB: not implemented"; return nil }

// Write always returns the configured error, simulating a write failure.
func (e *ErrorWriteCloser) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"

	// Close always returns the configured error, simulating a close failure.
	return 0, nil
}

func (e *ErrorWriteCloser) Close() error {
	_ = "STUB: not implemented"

	// CloseErrorWriteCloser is a mock io.WriteCloser where only the Close() method returns an error.
	// This is useful for testing scenarios where writes succeed but closing fails.
	return nil
}

type CloseErrorWriteCloser struct {
	w   io.Writer // Underlying writer for successful write operations
	err error     // Error to return when Close() is called
}

// NewCloseErrorWriteCloser creates a new WriteCloser that writes successfully
// but returns an error when Close() is called. This simulates partial failure scenarios.
func NewCloseErrorWriteCloser(w io.Writer, err error) *CloseErrorWriteCloser {
	_ = "STUB: not implemented"
	return nil
}

// Write implements the io.Writer interface by delegating to the underlying writer.
// This method always succeeds, allowing testing of close error scenarios.
func (c *CloseErrorWriteCloser) Write(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0,

		// Close always returns the configured error, simulating a close failure
		// while allowing writes to succeed.
		nil
}

func (c *CloseErrorWriteCloser) Close() error {
	_ = "STUB: not implemented"

	// ErrorReadWriteCloser is a mock that implements io.Reader, io.Writer, and io.Closer interfaces,
	// always returning the specified error for all operations. This is useful for testing
	// scenarios where all I/O operations fail, such as network failures or corrupted streams.
	return nil
}

type ErrorReadWriteCloser struct {
	Err error // The error to return for all read, write, and close operations
}

// NewErrorReadWriteCloser creates a new ErrorReadWriteCloser that will return
// the specified error for all read, write, and close operations. This mock is
// particularly useful for testing error handling in streaming operations where
// all I/O methods need to fail consistently.
func NewErrorReadWriteCloser(err error) *ErrorReadWriteCloser {
	_ = "STUB: not implemented"
	return nil
}

// Read always returns the configured error, simulating a read failure.
// This method implements the io.Reader interface for consistent error testing.
func (e *ErrorReadWriteCloser) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"

	// Write always returns the configured error, simulating a write failure.
	// This method implements the io.Writer interface for consistent error testing.
	return 0, nil
}

func (e *ErrorReadWriteCloser) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"

	// Close always returns the configured error, simulating a close failure.
	// This method implements the io.Closer interface for consistent error testing.
	return 0, nil
}

func (e *ErrorReadWriteCloser) Close() error {
	_ = "STUB: not implemented"

	// ErrorWriteAfterN is a mock io.Writer that succeeds for the first N writes
	// and then returns an error for all subsequent writes. This is useful for testing
	// scenarios where a writer works initially but fails after a certain number of operations,
	// such as disk full errors or connection drops.
	return nil
}

type ErrorWriteAfterN struct {
	N          int   // Number of successful writes before returning error
	Err        error // The error to return after N successful writes
	writeCount int   // Internal counter for tracking number of writes
	totalBytes int   // Total bytes written successfully (for testing/debugging)
}

// NewErrorWriteAfterN creates a new ErrorWriteAfterN that will allow N successful
// writes before returning the specified error. This is commonly used to test partial
// write scenarios and error recovery in streaming operations.
func NewErrorWriteAfterN(n int, err error) *ErrorWriteAfterN { _ = "STUB: not implemented"; return nil }

// Write implements the io.Writer interface. It succeeds for the first N calls
// and returns the configured error for all subsequent calls.
func (e *ErrorWriteAfterN) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteCount returns the number of write operations attempted (for testing).
func (e *ErrorWriteAfterN) WriteCount() int { _ = "STUB: not implemented"; return 0 }

// TotalBytes returns the total number of bytes successfully written (for testing).
func (e *ErrorWriteAfterN) TotalBytes() int { _ = "STUB: not implemented"; return 0 }

// Reset resets the write counter and total bytes, allowing the mock to be reused.
func (e *ErrorWriteAfterN) Reset() { _ = "STUB: not implemented"; return }

// CloseErrorReadCloser is a mock io.ReadCloser where only the Close() method returns an error.
// This is useful for testing scenarios where reads succeed but closing fails.
type CloseErrorReadCloser struct {
	r   io.Reader // Underlying reader for successful read operations
	err error     // Error to return when Close() is called
}

// NewCloseErrorReadCloser creates a new ReadCloser that reads successfully
// but returns an error when Close() is called. This simulates partial failure scenarios.
func NewCloseErrorReadCloser(r io.Reader, err error) *CloseErrorReadCloser {
	_ = "STUB: not implemented"
	return nil
}

// Read implements the io.Reader interface by delegating to the underlying reader.
// This method always succeeds, allowing testing of close error scenarios.
func (c *CloseErrorReadCloser) Read(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// Close always returns the configured error, simulating a close failure
		// while allowing reads to succeed.
		nil
}

func (c *CloseErrorReadCloser) Close() error { _ = "STUB: not implemented"; return nil }
