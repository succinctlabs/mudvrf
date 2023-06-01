import {FulfillRandomWordsEvent, RequestRandomWordsEvent, VRF} from "./types/VRFCoordinator";
import {VRFCoordinator__factory} from "./types/factories/VRFCoordinator__Factory";
import {MockVRFCoordinator} from "./types/MockVRFCoordinator";
import { IVRFCoordinatorSystem__factory } from "./types/factories/IVRFCoordinatorSystem__factory"
import {MockVRFCoordinator__factory} from "./types/factories/MockVRFCoordinator__factory";
import { ethers, utils } from 'ethers';
import { AbiCoder } from "ethers/lib/utils";
import { BigNumber } from "ethers";
import fs from "fs/promises";

async function main() {
    console.log("Listening for randomness requests...");

    // Anvil Chain ID
    const chainId = 31337;

    // TODO: Do we want to parse this?
    const provider = new ethers.providers.StaticJsonRpcProvider("http://127.0.0.1:8545", chainId);

    // TODO: Do we want to parse this?
    const wallet = new ethers.Wallet("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", provider);

    const signer = wallet.connect(provider);
    // Get the contract address from PostDeploy

    // Read address from vrf.json in contracts directory
    // TODO: Parse this in type-safe way
    const vrfJSON = JSON.parse(await fs.readFile("../example-contracts/vrf.json", 'utf8'));
    const vrfContractAddress = vrfJSON.vrfCoordinatorAddress;

    const vrfCoordinator = MockVRFCoordinator__factory.connect(vrfContractAddress, signer);

    const handleFulfilledRandomness = async (event: FulfillRandomWordsEvent) => {
        console.log("Randomness fulfilled");
        console.log("requestId: ", BigNumber.from(event.args.requestId).toHexString());
        console.log("Fulfilled on block: ", event.blockNumber);
    };

    const handleRandomnessRequest = async (event: RequestRandomWordsEvent) => {
        // Log all of the parameters
        console.log("Randomness request received");

        // Unused in mock
        const proof: VRF.ProofStruct = {
            pk: [event.args.requestId, event.args.requestId],
            gamma: [event.args.requestId, event.args.requestId],
            c: event.args.requestId,
            s: event.args.requestId,
            seed: event.args.requestId,
            uWitness: event.args.sender,
            cGammaWitness: [event.args.requestId, event.args.requestId],
            sHashWitness: [event.args.requestId, event.args.requestId],
            zInv: event.args.requestId
        }

        const request: VRF.RequestStruct = {
            sender: event.args.sender,
            nonce: BigNumber.from(event.args.nonce),
            oracleId: event.args.oracleId,
            nbWords: BigNumber.from(event.args.nbWords),
            requestConfirmations: 0,
            callbackGasLimit: BigNumber.from(event.args.callbackGasLimit),
            callbackAddress: event.args.callbackAddress,
            callbackSelector: event.args.callbackSelector,
            blockNumber: BigNumber.from(event.blockNumber)
        }

        console.log("nbWords: ", BigNumber.from(event.args.nbWords).toNumber());

        async function submitTransaction() {
            try {

                const tx = await vrfCoordinator.fulfillRandomWords(proof, request, {
                    gasLimit: 3000000,
                    gasPrice: ethers.utils.parseUnits('200', 'gwei')
                });
                // console.log('Transaction submitted:', tx.hash);
                console.log("Submitted fulfillRandomWords transaction");
                // const receipt = await tx.wait();
          
                // console.log('Transaction mined:', receipt.transactionHash);
          
              // Handle the transaction success
              // ...
            } catch (error) {
              // Handle revert or exception
              console.error('Transaction failed:', error);
              return;
            }
        }
          
        submitTransaction();

    };

    let latestBlockChecked = await provider.getBlockNumber();

    const fetchPastEvents = async () => {
        const requestRandomWordsFilter = vrfCoordinator.filters.RequestRandomWords();
        const fulfillRandomWordsFilter = vrfCoordinator.filters.FulfillRandomWords();
        const latestBlock = await provider.getBlockNumber();
        const rrwEvents = await vrfCoordinator.queryFilter(requestRandomWordsFilter, latestBlockChecked, latestBlock);
        const frwEvents = await vrfCoordinator.queryFilter(fulfillRandomWordsFilter, latestBlockChecked, latestBlock);
        rrwEvents.forEach(handleRandomnessRequest);
        frwEvents.forEach(handleFulfilledRandomness);
        latestBlockChecked = latestBlock + 1; // to avoid duplicate event fetching
    };
    
    fetchPastEvents(); // Fetch past events once at startup
    setInterval(fetchPastEvents, 1000); // Fetch new events every second

    // vrfCoordinator.on("RequestRandomWords(bytes32,address,uint256,bytes32,uint32,uint16,uint32,address,bytes4)", handleRandomnessRequest);
    // vrfCoordinator.on("FulfillRandomWords(bytes32)", handleFulfilledRandomness);

    // Handle exit signals
    process.on('SIGINT', () => {
        // vrfCoordinator.removeAllListeners('RequestRandomWords');
        // vrfCoordinator.removeAllListeners('FulfillRandomWords');
        console.log('Server shutting down...');
        process.exit();
    });
}
main();
