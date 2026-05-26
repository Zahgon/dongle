package sm2

type EncryptError struct {
	Err error
}

func (e EncryptError) Error() string { _ = "STUB: not implemented"; return "" }

type DecryptError struct {
	Err error
}

func (e DecryptError) Error() string { _ = "STUB: not implemented"; return "" }

type ReadError struct{ Err error }

func (e ReadError) Error() string { _ = "STUB: not implemented"; return "" }

type SignError struct {
	Err error
}

func (e SignError) Error() string { _ = "STUB: not implemented"; return "" }

type VerifyError struct {
	Err error
}

func (e VerifyError) Error() string { _ = "STUB: not implemented"; return "" }
