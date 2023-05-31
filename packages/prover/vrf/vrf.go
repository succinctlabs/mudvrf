package vrf

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/smartcontractkit/chainlink/v2/core/services/keystore/keys/vrfkey"
	"github.com/smartcontractkit/chainlink/v2/core/services/signatures/secp256k1"
	"github.com/smartcontractkit/chainlink/v2/core/services/vrf/proof"
)

type VRFPublicKeyUncompressed struct {
	X *big.Int
	Y *big.Int
}

type VRFKey struct {
	vrfkey vrfkey.KeyV2
}

type VRFProof struct {
	proof vrfkey.Proof
}

type VRFProofWithWitnesses struct {
	PublicKey     [2]*big.Int
	Gamma         [2]*big.Int
	C             *big.Int
	S             *big.Int
	Seed          *big.Int
	UWitness      common.Address
	CGammaWitness [2]*big.Int
	SHashWitness  [2]*big.Int
	ZInv          *big.Int
}

func NewSecp256k1VRFKey(privateKey string) (*VRFKey, error) {
	ecdsaKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("failed to convert private key to ECDSA key")
	}
	secretKeyBytes := ecdsaKey.D.Bytes()
	vrfkey := VRFKey{
		vrfkey: vrfkey.Raw(secretKeyBytes).Key(),
	}
	return &vrfkey, nil
}

func (k *VRFKey) GenerateProof(seed *big.Int) (VRFProof, error) {
	proof, err := k.vrfkey.GenerateProof(seed)
	return VRFProof{proof}, err
}

func (k *VRFKey) OracleAddress() common.Address {
	return k.vrfkey.PublicKey.Address()
}

func (k *VRFKey) OracleId() ([]byte, error) {
	proof, err := k.GenerateProof(big.NewInt(1))
	if err != nil {
		return nil, err
	}

	x, y := secp256k1.Coordinates(proof.proof.PublicKey)
	oracleId := Keccak256Pair(x, y)

	return oracleId, nil
}

func (p *VRFProof) VerifyProof() (bool, error) {
	return p.proof.VerifyVRFProof()
}

func (p *VRFProof) CalculateWitnesses() (*VRFProofWithWitnesses, error) {
	solidityProof, err := proof.SolidityPrecalculations(&p.proof)
	if err != nil {
		return nil, err
	}

	pubkeyX, pubkeyY := secp256k1.Coordinates(solidityProof.P.PublicKey)
	gammaX, gammaY := secp256k1.Coordinates(solidityProof.P.Gamma)
	c := solidityProof.P.C
	s := solidityProof.P.S
	seed := solidityProof.P.Seed
	uWitness := solidityProof.UWitness
	cGammaWitnessX, cGammaWitnessY := secp256k1.Coordinates(solidityProof.CGammaWitness)
	sHashWitnessX, sHashWitnessY := secp256k1.Coordinates(solidityProof.SHashWitness)
	zInv := solidityProof.ZInv

	return &VRFProofWithWitnesses{
		PublicKey:     [2]*big.Int{pubkeyX, pubkeyY},
		Gamma:         [2]*big.Int{gammaX, gammaY},
		C:             c,
		S:             s,
		Seed:          seed,
		UWitness:      uWitness,
		CGammaWitness: [2]*big.Int{cGammaWitnessX, cGammaWitnessY},
		SHashWitness:  [2]*big.Int{sHashWitnessX, sHashWitnessY},
		ZInv:          zInv,
	}, nil
}
