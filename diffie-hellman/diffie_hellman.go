package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	pMinusThree := bigZero().Add(p, big.NewInt(-3))
	randomNumber := bigZero().Rand(rand.New(rand.NewSource(time.Now().UnixNano())), pMinusThree)
	return bigZero().Add(randomNumber, big.NewInt(2))
}

func PublicKey(a, p *big.Int, g int64) *big.Int {
	bigG := big.NewInt(g)
	return bigZero().Exp(bigG, a, p)
}

func SecretKey(a, B, p *big.Int) *big.Int {
	return bigZero().Exp(B, a, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	privateKey := PrivateKey(p)
	return privateKey, PublicKey(privateKey, p, g)
}

func bigZero() *big.Int {
	return big.NewInt(0)
}