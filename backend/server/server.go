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
	privateKey     string
	key            *ecdsa.PrivateKey
	vrfkey         *vrf.VRFKey
	vrfCoordinator *bindings.VRFCoordinator
	blockHashStore *bindings.BlockHashStore
	ethclient      *ethclient.Client
}

// abigen --abi ../packages/contracts/out/VRFCoordinator.sol/VRFCoordinator.abi.json --out ../backend/bindings/VRFCoordinator.go --pkg bindings --type VRFCoordinator
func New(privateKey string, vrfCoordinatorAddress common.Address, blockHashStoreAddress common.Address, ethclient *ethclient.Client) *VRFRequestWatcher {
	vrfCoordinator, err := bindings.NewVRFCoordinator(vrfCoordinatorAddress, ethclient)
	if err != nil {
		panic(err)
	}
	blockHashStore, err := bindings.NewBlockHashStore(blockHashStoreAddress, ethclient)
	if err != nil {
		panic(err)
	}
	vrfkey, err := vrf.NewSecp256k1VRFKey(privateKey)
	if err != nil {
		panic(err)
	}
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		panic(err)
	}
	return &VRFRequestWatcher{
		privateKey:     privateKey,
		key:            key,
		vrfkey:         vrfkey,
		vrfCoordinator: vrfCoordinator,
		blockHashStore: blockHashStore,
		ethclient:      ethclient,
	}
}

func (s *VRFRequestWatcher) Start() {
	oracleId, err := s.vrfkey.OracleId()
	if err != nil {
		panic(err)
	}

	fmt.Println("\nStarting VRFRequestWatcher...")
	fmt.Println("Oracle:", hex.EncodeToString(oracleId))

	address, err := s.vrfCoordinator.MAXIMUMNBWORDS(nil)
	if err != nil {
		panic(err)
	}
	fmt.Println(address)

	fmt.Println("\nListening for RequestRandomWords events...")
	requests := make(chan *bindings.VRFCoordinatorRequestRandomWords)
	s.vrfCoordinator.WatchRequestRandomWords(nil, requests)

	for {
		select {
		case event := <-requests:
			fmt.Printf("Event: %+v\n", hex.EncodeToString(event.RequestId[:]))
			fmt.Println("> RequestId:", "0x"+hex.EncodeToString(event.RequestId[:]))
			fmt.Println("> Sender:", event.Sender.Hex())
			fmt.Println("> Nonce:", event.Nonce)
			fmt.Println("> OracleId:", "0x"+hex.EncodeToString(event.OracleId[:]))
			fmt.Println("> NbWords:", event.NbWords)
			fmt.Println("> CallbackGasLimit:", event.CallbackGasLimit)
			fmt.Println("> CallbackSelector:", "0x"+hex.EncodeToString(event.CallbackSelector[:]))
			fmt.Println("> BlockNumber:", event.Raw.BlockNumber)
			fmt.Println("Storing block hash...")
			err := s.StoreBlockHash(event)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("Fulfilling random words...")
			err = s.FulfillRandomWords(event)
			if err != nil {
				fmt.Println("Error:", err)
			}
			fmt.Println("")
		}
	}
}

func (s *VRFRequestWatcher) StoreBlockHash(event *bindings.VRFCoordinatorRequestRandomWords) error {
	key, err := crypto.HexToECDSA(s.privateKey)
	if err != nil {
		return fmt.Errorf("failed to parse private key: %v", err)
	}

	transactor, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(31337))
	if err != nil {
		return fmt.Errorf("creating transactor: %w", err)
	}

	blockNumber := big.NewInt(int64(event.Raw.BlockNumber))
	tx, err := s.blockHashStore.StoreBlockHashesViaOpCode(transactor, []*big.Int{blockNumber})
	if err != nil {
		return fmt.Errorf("storing block hash: %w", err)
	}
	fmt.Printf("tx: %+v\n", tx.Hash().Hex())

	return nil
}

func (s *VRFRequestWatcher) FulfillRandomWords(event *bindings.VRFCoordinatorRequestRandomWords) error {
	seedBytes := vrf.Keccak256AddressAndU256(event.Sender, event.Nonce)
	seed := new(big.Int).SetBytes(seedBytes)
	actualSeedBytes := vrf.Keccak256Pair(seed, event.Raw.BlockHash.Big())
	actualSeed := new(big.Int).SetBytes(actualSeedBytes)

	proof, err := s.vrfkey.GenerateProof(actualSeed)
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
		NbWords:              event.NbWords,
		RequestConfirmations: event.RequestConfirmations,
		CallbackGasLimit:     event.CallbackGasLimit,
		CallbackSelector:     [4]byte(event.CallbackSelector),
		BlockNumber:          event.Raw.BlockNumber,
	}

	tx, err := s.vrfCoordinator.FulfillRandomWords(transactor, resultProof, resultRequest)
	if err != nil {
		return fmt.Errorf("fulfilling randomness: %w", err)
	}
	fmt.Printf("tx: %+v\n", tx.Hash().Hex())

	return nil
}
