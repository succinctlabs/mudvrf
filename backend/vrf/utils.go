package vrf

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/umbracle/ethgo/abi"
)

func Keccak256Pair(a *big.Int, b *big.Int) []byte {
	type Value struct {
		A *big.Int
		B *big.Int
	}
	typ := abi.MustNewType("tuple(uint256 a, uint256 b)")
	value := &Value{
		a,
		b,
	}
	encoded, err := typ.Encode(value)
	if err != nil {
		panic(err)
	}
	return crypto.Keccak256(encoded)
}

func Keccak256AddressAndU256(a common.Address, b *big.Int) []byte {
	type Value struct {
		A common.Address
		B *big.Int
	}
	typ := abi.MustNewType("tuple(address a, uint256 b)")
	value := &Value{
		a,
		b,
	}
	encoded, err := typ.Encode(value)
	if err != nil {
		panic(err)
	}
	return crypto.Keccak256(encoded)
}
