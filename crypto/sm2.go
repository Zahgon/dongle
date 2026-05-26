package crypto

import (
	"github.com/dromara/dongle/crypto/keypair"
)

// BySm2 encrypts by SM2.
func (e Encrypter) BySm2(kp *keypair.Sm2KeyPair) Encrypter {
	_ = "STUB: not implemented"
	return *new(Encrypter)
}

// Streaming encryption mode

// Standard encryption mode

// BySm2 decrypts by SM2.
func (d Decrypter) BySm2(kp *keypair.Sm2KeyPair) Decrypter {
	_ = "STUB: not implemented"
	return *new(Decrypter)
}

// Streaming decryption mode

// Standard decryption mode

// BySm2 signs by SM2.
func (s Signer) BySm2(kp *keypair.Sm2KeyPair) Signer {
	_ = "STUB: not implemented"
	return *new(Signer)
}

// Streaming signing mode

// Standard signing mode

// BySm2 verifies by SM2.
func (v Verifier) BySm2(kp *keypair.Sm2KeyPair) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

// Streaming verification mode

// Write the data to be verified

// Close the verifier to perform verification

// Standard verification mode
