import {VRF} from "./types/ethers-contracts/VRFCoordinator";
import {VRFCoordinator__factory} from "./types/ethers-contracts/factories/VRFCoordinator__Factory";
import {MockVRFCoordinator} from "./types/ethers-contracts/MockVRFCoordinator";
import {MockVRFCoordinator__factory} from "./types/ethers-contracts/factories/MockVRFCoordinator__Factory";
import { ethers, utils } from 'ethers';
import { AbiCoder } from "ethers/lib/utils";
import { BigNumber } from "ethers";
import fs from "fs/promises";

async function main() {
    console.log("Listening for randomness requests...");

    // TODO: Do we want to parse this?
    const provider = new ethers.providers.StaticJsonRpcProvider("http://127.0.0.1:8545", 31337);

    // TODO: Do we want to parse this?
    const wallet = new ethers.Wallet("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", provider);

    const signer = wallet.connect(provider);
    // Get the contract address from PostDeploy

    // Read address from vrf.json in contracts directory
    // TODO: Parse this in type-safe way
    const vrfJson = JSON.parse(await fs.readFile("../contracts/vrf.json", 'utf8'));
    const contractAddress = vrfJson.vrfCoordinatorAddress;
    const vrfCoordinator = MockVRFCoordinator__factory.connect(contractAddress, signer);

    const handleFulfilledRandomness = async (requestId: string, event: ethers.Event) => {
        console.log("Randomness fulfilled");
        console.log("requestId: ", BigNumber.from(requestId).toHexString());
        console.log("Fulfilled on block: ", event.blockNumber);
    };

    const handleRandomnessRequest = async (requestId: string, sender: string, nonce: string, oracleId: string, nbWords: string, requestConfirmations: string, callbackGasLimit: string, callbackSelector: string, event: ethers.Event) => {
        // Log all of the parameters
        console.log("Randomness request received");

        // Unused in mock
        const proof: VRF.ProofStruct = {
            pk: [requestId, requestId],
            gamma: [requestId, requestId],
            c: requestId,
            s: requestId,
            seed: requestId,
            uWitness: sender,
            cGammaWitness: [requestId, requestId],
            sHashWitness: [requestId, requestId],
            zInv: requestId
        }

        const request: VRF.RequestStruct = {
            sender: sender,
            nonce: BigNumber.from(nonce),
            oracleId: oracleId,
            nbWords: BigNumber.from(nbWords),
            requestConfirmations: 0,
            callbackGasLimit: BigNumber.from(callbackGasLimit),
            callbackSelector: callbackSelector,
            blockNumber: BigNumber.from(event.blockNumber)
        }

        console.log("nbWords: ", BigNumber.from(nbWords).toNumber());

        async function submitTransaction() {
            try {

                const tx = await vrfCoordinator.fulfillRandomWords(proof, request, {
                    gasLimit: 3000000,
                    gasPrice: ethers.utils.parseUnits('200', 'gwei')
                });
                console.log('Transaction submitted:', tx.hash);
                const receipt = await tx.wait();
          
                console.log('Transaction mined:', receipt.transactionHash);
          
              // Handle the transaction success
              // ...
            } catch (error) {
              // Handle revert or exception
              console.error('Transaction failed:', error);
              return;
            }
        }
          
        submitTransaction();

        console.log("Submitted fulfillRandomWords transaction");
    };

    vrfCoordinator.on("RequestRandomWords", handleRandomnessRequest);
    vrfCoordinator.on("FulfillRandomWords", handleFulfilledRandomness);

    // Handle exit signals
    process.on('SIGINT', () => {
        vrfCoordinator.removeAllListeners('RequestRandomWords');
        vrfCoordinator.removeAllListeners('FulfillRandomWords');
        console.log('Server shutting down...');
        process.exit();
    });
}
main();
