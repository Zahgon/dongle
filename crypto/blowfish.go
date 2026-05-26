package crypto

import (
	"github.com/dromara/dongle/crypto/cipher"
)

// ByBlowfish encrypts by blowfish.
func (e Encrypter) ByBlowfish(c *cipher.BlowfishCipher) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// ByBlowfish decrypts by blowfish.
func (d Decrypter) ByBlowfish(c *cipher.BlowfishCipher) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode
