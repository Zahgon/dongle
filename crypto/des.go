package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByDes encrypts by des.
func (e Encrypter) ByDes(c *cipher.DesCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByDes decrypts by des.
func (d Decrypter) ByDes(c *cipher.DesCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
