package server

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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
	s.contract.WatchRequestRandomWords(nil, requests)

	for {
		select {
		case event := <-requests:
			fmt.Printf("Event: %+v\n", hex.EncodeToString(event.RequestId[:]))
			err := s.FulfillRandomWords(event)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	}
}

func (s *VRFRequestWatcher) FulfillRandomWords(event *bindings.VRFCoordinatorRequestRandomWords) error {
	fmt.Println("> RequestId:", "0x"+hex.EncodeToString(event.RequestId[:]))
	fmt.Println("> Sender:", event.Sender.Hex())
	fmt.Println("> Nonce:", event.Nonce)
	fmt.Println("> OracleId:", "0x"+hex.EncodeToString(event.OracleId[:]))
	fmt.Println("> NbWords:", event.NbWords)
	fmt.Println("> CallbackGasLimit:", event.CallbackGasLimit)
	fmt.Println("> CallbackSelector:", "0x"+hex.EncodeToString(event.CallbackSelector[:]))
	fmt.Println("> BlockNumber:", event.Raw.BlockNumber)

	seedBytes := vrf.Keccak256AddressAndU256(event.Sender, event.Nonce)
	seed := new(big.Int).SetBytes(seedBytes)

	proof, err := s.vrfkey.GenerateProof(seed)
	if err != nil {
		return err
	}

	proofWithWitnesses, err := proof.CalculateWitnesses()
	if err != nil {
		return err
	}

	key, err := crypto.HexToECDSA(s.privateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
	if err != nil {
		return fmt.Errorf("creating transactor: %w", err)
	}

	resultProof := bindings.VRFProof{
		Pk:            proofWithWitnesses.PublicKey,
		Gamma:         proofWithWitnesses.Gamma,
		C:             proofWithWitnesses.C,
		S:             proofWithWitnesses.S,
		Seed:          proofWithWitnesses.Seed,
		UWitness:      proofWithWitnesses.UWitness,
		CGammaWitness: proofWithWitnesses.CGammaWitness,
		SHashWitness:  proofWithWitnesses.SHashWitness,
		ZInv:          proofWithWitnesses.ZInv,
	}

	resultRequest := bindings.VRFRequest{
		Sender:               event.Sender,
		Nonce:                event.Nonce,
		OracleId:             event.OracleId,
		RequestConfirmations: event.RequestConfirmations,
		CallbackGasLimit:     event.CallbackGasLimit,
		NbWords:              event.NbWords,
		CallbackSelector:     [4]byte(event.CallbackSelector),
	}

	tx, err := s.contract.FulfillRandomWords(transactor, resultProof, resultRequest)
	if err != nil {
		return fmt.Errorf("fulfilling randomness: %w", err)
	}
	fmt.Printf("tx: %+v\n", tx.Hash().Hex())

	return nil
}
