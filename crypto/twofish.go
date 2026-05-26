package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByTwofish encrypts by twofish.
func (e Encrypter) ByTwofish(c *cipher.TwofishCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByTwofish decrypts by twofish.
func (d Decrypter) ByTwofish(c *cipher.TwofishCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
