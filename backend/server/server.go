package server

import (
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jtguibas/mudvrf/bindings"
	"github.com/jtguibas/mudvrf/vrf"
)

type VRFServer struct {
	privateKey string
	key        *ecdsa.PrivateKey
	vrfkey     *vrf.VRFKey
	contract   *bindings.IWorld
	ethclient  *ethclient.Client
}

// bytes32 constant _tableId = bytes32(abi.encodePacked(bytes16(""), bytes16("VRFRequestTableV")));
// var TABLE_ID = common.Hex2Bytes("00000000000000000000000000000000565246526571756573745461626c6556")
var TABLE_ID = common.Hex2Bytes("0000000000000000000000000000000056524652657175657374730000000000")

// abigen --abi ../../solidity/contracts/VRFConsumer.abi --pkg bindings --type VRFSystem --out ../bindings/VRFSystem.go --bin ..
func New(privateKey string, contractAddress common.Address, ethclient *ethclient.Client) *VRFServer {
	contract, err := bindings.NewIWorld(contractAddress, ethclient)
	if err != nil {
		panic(err)
	}
	vrfkey, err := vrf.NewSecp256k1VRFKey(privateKey)
	if err != nil {
		panic(err)
	}
	key, err := crypto.HexToECDSA(privateKey)
	return &VRFServer{
		privateKey: privateKey,
		key:        key,
		vrfkey:     vrfkey,
		contract:   contract,
		ethclient:  ethclient,
	}
}

type VRFRequest struct {
}

func (s *VRFServer) Start() {
	oracleId, err := s.vrfkey.OracleId()
	if err != nil {
		panic(err)
	}
	fmt.Println("starting with oracle id", hex.EncodeToString(oracleId))
	requests := make(chan *bindings.IWorldStoreSetRecord)
	// signals := make(chan os.Signal, 1)
	// s.contract.WatchRandomnessRequest(nil, requests, nil, nil)
	s.contract.WatchStoreSetRecord(nil, requests)

	for {
		select {
		case event := <-requests:
			fmt.Println("got event")
			fmt.Printf("event: %+v\n", hex.EncodeToString(event.Data))
			// tableKey := crypto.Keccak256([]byte("VRFRequestTableV2"))
			// fmt.Printf("tableKey: %+v\n", hex.EncodeToString(tableKey))
			// fmt.Printf("event.Table: %+v\n", hex.EncodeToString(event.Table[:]))
			// fmt.Printf("event.Table: %+v\n", hex.EncodeToString(TABLE_ID))
			if reflect.DeepEqual(event.Table[:], TABLE_ID) {
				err := s.FulfillRandomness(event.Data)
				if err != nil {
					fmt.Println(err)
				}
			}
			// case <-signals:
			// 	break
		}
	}
}

func (s *VRFServer) FulfillRandomness(req []byte) error {
	// First 20 bytes is address
	address := common.BytesToAddress(req[:20])
	// Next 32 bytes is nonce
	nonce := new(big.Int).SetBytes(req[20:52])
	// Next 32 bytes is blockNumber
	blockNumber := new(big.Int).SetBytes(req[52:84])
	// Next 4 bytes is gasLimit
	gasLimit := new(big.Int).SetBytes(req[84:88])
	// Next 4 bytes is numWords
	numWords := new(big.Int).SetBytes(req[88:92])
	// Last 4 bytes is callback
	callback := req[92:96]

	fmt.Println("address", address)
	fmt.Println("nonce", nonce)
	fmt.Println("blockNumber", blockNumber)
	fmt.Println("gasLimit", gasLimit)
	fmt.Println("numWords", numWords)
	fmt.Println("callback", hex.EncodeToString(callback))

	seedBytes := vrf.Keccak256AddressAndU256(address, nonce)
	seed := new(big.Int).SetBytes(seedBytes)
	proof, err := s.vrfkey.GenerateProof(seed)
	if err != nil {
		return err
	}
	proofWithWitnesses, err := proof.CalculateWitnesses()
	if err != nil {
		return err
	}
	// oracleAddress := s.vrfkey.OracleAddress()
	// oracleId, err := s.vrfkey.OracleId()
	// if err != nil {
	// 	return err
	// }

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
		BlockNumber:      blockNumber.Uint64(),
		CallbackGasLimit: uint32(gasLimit.Uint64()),
		NbWords:          uint32(numWords.Uint64()),
		Sender:           address,
		CallbackSelector: [4]byte(callback),
	}
	tx, err := s.contract.FulfillRandomWords(transactor, resultProof, resultRequest)
	if err != nil {
		return fmt.Errorf("fulfilling randomness: %w", err)
	}
	fmt.Printf("tx: %+v\n", tx.Hash().Hex())

	return nil
}
