package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByTea encrypts by tea.
func (e Encrypter) ByTea(c *cipher.TeaCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByTea decrypts by tea.
func (d Decrypter) ByTea(c *cipher.TeaCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
