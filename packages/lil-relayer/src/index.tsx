import {VRF} from "../types/ethers-contracts/VRFCoordinator";
import {MockVRFCoordinator, RandomnessRequestEventObject} from "../types/ethers-contracts/MockVRFCoordinator";
import {MockVRFCoordinator__factory} from "../types/ethers-contracts/factories/MockVRFCoordinator__Factory";
import { ethers, utils } from 'ethers';
import { AbiCoder } from "ethers/lib/utils";
import { BigNumber } from "ethers";


async function main() {
    console.log("Listening for randomness requests...");

    const provider = new ethers.providers.JsonRpcProvider("http://localhost:8545");

    const wallet = new ethers.Wallet("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", provider);

    const signer = wallet.connect(provider);
    // Get the contract address from PostDeploy
    const contractAddress = "0xc3e53F4d16Ae77Db1c982e75a937B9f60FE63690";
    const vrfCoordinator = MockVRFCoordinator__factory.connect(contractAddress, provider);

    const handleFulfilledRandomness = async (nonce: string, requestId: string, words: string[]) => {
        console.log("Randomness fulfilled");
        console.log("requestId: ", BigNumber.from(nonce).toNumber());
        console.log("randomness: ", BigNumber.from(requestId).toNumber());
        console.log("words: ", words);
    };

    const handleRandomnessRequest = async (nonce: string, requestId: string, sender: string, blockNumber: string, seed: string, callbackGasLimit: string, nbWords: string, callbackSelector: string) => {
        // Log all of the parameters
        console.log("Randomness request received");

        // Call the fulfillRandomWords function

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

        const oracleId = utils.keccak256(utils.defaultAbiCoder.encode(["address", "uint256"], [sender, nonce]));

        const request: VRF.RequestStruct = {
            sender: sender,
            nonce: nonce,
            blockNumber: blockNumber,
            oracleId: oracleId,
            requestConfirmations: 0,
            callbackGasLimit: callbackGasLimit,
            nbWords: nbWords,
            callbackSelector: callbackSelector,
        }

        console.log("nbWords: ", BigNumber.from(nbWords).toNumber());

        const txResponse = await vrfCoordinator.connect(signer).fulfillRandomWords(proof, request);
        const receipt = await txResponse.wait();
        console.log(receipt);



        console.log("Submitted fulfillRandomWords transaction");
    };
    vrfCoordinator.on("RandomnessRequest", handleRandomnessRequest);
    vrfCoordinator.on("FulfilledRandomness", handleFulfilledRandomness);

    // Handle exit signals
    process.on('SIGINT', () => {
        vrfCoordinator.removeAllListeners('RandomnessRequested');
        console.log('Server shutting down...');
        process.exit();
    });
}
main();
