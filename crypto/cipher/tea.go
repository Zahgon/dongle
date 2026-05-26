package cipher

// TeaCipher defines a TeaCipher struct.
type TeaCipher struct {
	blockCipher
	Rounds int
}

// NewTeaCipher returns a new TeaCipher instance.
func NewTeaCipher(block BlockMode) *TeaCipher { _ = "STUB: not implemented"; return nil }

// SetRounds sets the number of rounds for the cipher.
func (c *TeaCipher) SetRounds(rounds int) { _ = "STUB: not implemented"; return }
