package crypto

import (
	"github.com/dromara/dongle/crypto/keypair"
)

// ByEd25519 signs by ed25519.
func (s Signer) ByEd25519(kp *keypair.Ed25519KeyPair) Signer {
	_ = "STUB: not implemented"
	return *new(Signer)
}

// Streaming signing mode

// Standard signing mode

// ByEd25519 verifies by ed25519.
func (v Verifier) ByEd25519(kp *keypair.Ed25519KeyPair) Verifier {
	_ = "STUB: not implemented"
	return *new(Verifier)
}

// Streaming verification mode

// Create a stream verifier

// Write data to the stream verifier

// Close the verifier to perform verification

// Standard verification mode
