package server

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jtguibas/mudvrf/bindings"
	"github.com/jtguibas/mudvrf/vrf"
)

type VRFRequestWatcher struct {
	privateKey string
	key        *ecdsa.PrivateKey
	vrfkey     *vrf.VRFKey
	contract   *bindings.VRFCoordinator
	ethclient  *ethclient.Client
}

// abigen --abi ../packages/contracts/out/VRFCoordinator.sol/VRFCoordinator.abi.json --out ../backend/bindings/VRFCoordinator.go --pkg bindings --type VRFCoordinator
func New(privateKey string, contractAddress common.Address, ethclient *ethclient.Client) *VRFRequestWatcher {
	contract, err := bindings.NewVRFCoordinator(contractAddress, ethclient)
	if err != nil {
		panic(err)
	}
	vrfkey, err := vrf.NewSecp256k1VRFKey(privateKey)
	if err != nil {
		panic(err)
	}
	key, err := crypto.HexToECDSA(privateKey)
	return &VRFRequestWatcher{
		privateKey: privateKey,
		key:        key,
		vrfkey:     vrfkey,
		contract:   contract,
		ethclient:  ethclient,
	}
}

func (s *VRFRequestWatcher) Start() {
	oracleId, err := s.vrfkey.OracleId()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nStarting VRFRequestWatcher...")
	fmt.Println("Oracle:", hex.EncodeToString(oracleId))

	address, err := s.contract.MAXIMUMNBWORDS(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)

	fmt.Println("\nListening for RequestRandomWords events...")
	requests := make(chan *bindings.VRFCoordinatorRequestRandomWords)
	s.contract.WatchRequestRandomWords(nil, requests, nil, nil)

	for {
		select {
		case event := <-requests:
			fmt.Printf("event: %+v\n", hex.EncodeToString(event.RequestId[:]))
		}
	}
}

// func (s *VRFRequestWatcher) FulfillRandomness(req []byte) error {
// 	// First 20 bytes is address
// 	address := common.BytesToAddress(req[:20])
// 	// Next 32 bytes is nonce
// 	nonce := new(big.Int).SetBytes(req[20:52])
// 	// Next 32 bytes is blockNumber
// 	blockNumber := new(big.Int).SetBytes(req[52:84])
// 	// Next 4 bytes is gasLimit
// 	gasLimit := new(big.Int).SetBytes(req[84:88])
// 	// Next 4 bytes is numWords
// 	numWords := new(big.Int).SetBytes(req[88:92])
// 	// Last 4 bytes is callback
// 	callback := req[92:96]

// 	fmt.Println("address", address)
// 	fmt.Println("nonce", nonce)
// 	fmt.Println("blockNumber", blockNumber)
// 	fmt.Println("gasLimit", gasLimit)
// 	fmt.Println("numWords", numWords)
// 	fmt.Println("callback", hex.EncodeToString(callback))

// 	seedBytes := vrf.Keccak256AddressAndU256(address, nonce)
// 	seed := new(big.Int).SetBytes(seedBytes)
// 	proof, err := s.vrfkey.GenerateProof(seed)
// 	if err != nil {
// 		return err
// 	}
// 	proofWithWitnesses, err := proof.CalculateWitnesses()
// 	if err != nil {
// 		return err
// 	}
// 	// oracleAddress := s.vrfkey.OracleAddress()
// 	// oracleId, err := s.vrfkey.OracleId()
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	key, err := crypto.HexToECDSA(s.privateKey)
// 	if err != nil {
// 		return fmt.Errorf("failed to parse private key: %v", err)
// 	}
// 	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
// 	if err != nil {
// 		return fmt.Errorf("creating transactor: %w", err)
// 	}

// 	resultProof := bindings.VRFProof{
// 		Pk:            proofWithWitnesses.PublicKey,
// 		Gamma:         proofWithWitnesses.Gamma,
// 		C:             proofWithWitnesses.C,
// 		S:             proofWithWitnesses.S,
// 		Seed:          proofWithWitnesses.Seed,
// 		UWitness:      proofWithWitnesses.UWitness,
// 		CGammaWitness: proofWithWitnesses.CGammaWitness,
// 		SHashWitness:  proofWithWitnesses.SHashWitness,
// 		ZInv:          proofWithWitnesses.ZInv,
// 	}
// 	resultRequest := bindings.VRFRequest{
// 		BlockNumber:      blockNumber.Uint64(),
// 		CallbackGasLimit: uint32(gasLimit.Uint64()),
// 		NbWords:          uint32(numWords.Uint64()),
// 		Sender:           address,
// 		CallbackSelector: [4]byte(callback),
// 	}
// 	tx, err := s.contract.FulfillRandomWords(transactor, resultProof, resultRequest)
// 	if err != nil {
// 		return fmt.Errorf("fulfilling randomness: %w", err)
// 	}
// 	fmt.Printf("tx: %+v\n", tx.Hash().Hex())

// 	return nil
// }
