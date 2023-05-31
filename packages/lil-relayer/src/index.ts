import {VRF} from "./types/ethers-contracts/VRFCoordinator";
import {VRFCoordinator__factory} from "./types/ethers-contracts/factories/VRFCoordinator__Factory";
import {MockVRFCoordinator} from "./types/ethers-contracts/MockVRFCoordinator";
import {MockVRFCoordinator__factory} from "./types/ethers-contracts/factories/MockVRFCoordinator__Factory";
import { ethers, utils } from 'ethers';
import { AbiCoder } from "ethers/lib/utils";
import { BigNumber } from "ethers";

async function main() {
    console.log("Listening for randomness requests...");

    const provider = new ethers.providers.StaticJsonRpcProvider("http://localhost:8545", 1337);

    const wallet = new ethers.Wallet("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", provider);

    const signer = wallet.connect(provider);
    // Get the contract address from PostDeploy

    const contractAddress = "0x84eA74d481Ee0A5332c457a4d796187F6Ba67fEB";
    const vrfCoordinator = VRFCoordinator__factory.connect(contractAddress, provider);

    const handleFulfilledRandomness = async (nonce: string, requestId: string, words: string[]) => {
        console.log("Randomness fulfilled");
        console.log("nonce: ", BigNumber.from(nonce).toBigInt());
        console.log("requestId: ", BigNumber.from(requestId).toHexString());
        console.log("words: ", (words).map((word) => BigNumber.from(word).toHexString()));
    };

    // console.log("Here: ", await vrfCoordinator.MAXIMUM_NB_WORDS());

    const handleRandomnessRequest = async (nonce: string, requestId: string, sender: string, blockNumber: string, seed: string, callbackGasLimit: string, nbWords: string, callbackSelector: string) => {
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

        console.log("Callback Selector: ", callbackSelector);

        console.log("nbWords: ", BigNumber.from(nbWords).toNumber());

        async function submitTransaction() {
            try {
            const txData = vrfCoordinator.connect(signer).fulfillRandomWords(proof, request);
            const txOptions = { gasLimit: 300000, gasPrice: ethers.utils.parseUnits('20', 'gwei') };
              const transaction = await signer.sendTransaction({ ...txData, ...txOptions });
              console.log('Transaction submitted:', transaction.hash);
          
              // Wait for the transaction to be mined
              const receipt = await transaction.wait();
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

    try {
        console.log("Num words", await vrfCoordinator.MAXIMUM_NB_WORDS());
    } catch(e) {
        console.log(e);
    }

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
