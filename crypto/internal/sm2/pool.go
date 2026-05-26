package sm2

import (
	"math/big"
	"sync"
)

// bigIntPool is a sync.Pool for reusing big.Int objects to reduce allocations.
var bigIntPool = sync.Pool{
	New: func() interface{} {
		return new(big.Int)
	},
}

// getBigInt gets a big.Int from the pool.
func getBigInt() *big.Int { _ = "STUB: not implemented"; return nil }

// putBigInt returns a big.Int to the pool after zeroing it.
func putBigInt(x *big.Int) { _ = "STUB: not implemented"; return }

// putBigInts returns multiple big.Ints to the pool.
func putBigInts(xs ...*big.Int) { _ = "STUB: not implemented"; return }
