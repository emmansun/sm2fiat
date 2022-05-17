package sm2c

import (
	"crypto/elliptic"
	"math/big"
	"sync"
)

var initonce sync.Once

func bigFromDecimal(s string) *big.Int {
	b, ok := new(big.Int).SetString(s, 10)
	if !ok {
		panic("crypto/elliptic: internal error: invalid encoding")
	}
	return b
}

func bigFromHex(s string) *big.Int {
	b, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("sm2/elliptic: internal error: invalid encoding")
	}
	return b
}

func initAll() {
	initSM2P256()
}

func SM2P256() elliptic.Curve {
	initonce.Do(initAll)
	return sm2p256
}
