package cipher

// ChaCha20Cipher defines a ChaCha20Cipher struct.
type ChaCha20Cipher struct {
	baseCipher
	Nonce []byte
}

// NewChaCha20Cipher returns a new ChaCha20Cipher instance.
func NewChaCha20Cipher() (c *ChaCha20Cipher) { _ = "STUB: not implemented"; return nil }

// SetNonce sets the nonce for the cipher.
func (c *ChaCha20Cipher) SetNonce(nonce []byte) { _ = "STUB: not implemented"; return }
