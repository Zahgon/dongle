package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByAes encrypts by aes.
func (e Encrypter) ByAes(c *cipher.AesCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByAes decrypts by aes.
func (d Decrypter) ByAes(c *cipher.AesCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
