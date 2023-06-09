import {FulfillRandomWordsEvent, RequestRandomWordsEvent, VRF} from "./types/VRFCoordinator";
import {MockVRFCoordinator__factory} from "./types/factories/MockVRFCoordinator__factory";
import { ethers, utils } from 'ethers';
import { BigNumber } from "ethers";
import fs from "fs/promises";
import yargs from "yargs";
import { hideBin } from 'yargs/helpers';

async function main() {
    const argv = await yargs(hideBin(process.argv))
        .option('vrfJsonPath', {
        type: 'string',
        demandOption: true,
        describe: 'Path to the VRF JSON file'
        })
        .option('rpcUrl', {
        type: 'string',
        demandOption: true,
        describe: 'RPC URL'
        })
        .option('privateKey', {
        type: 'string',
        demandOption: true,
        describe: 'Private key'
        })
        .argv;

    const vrfJSONPath = argv.vrfJsonPath;
    const rpcURL = argv.rpcUrl;
    const privateKey = argv.privateKey;

    const chainId = 31337;

    const provider = new ethers.providers.WebSocketProvider(rpcURL, chainId);

    const wallet = new ethers.Wallet('0x'+privateKey, provider);

    const signer = wallet.connect(provider);

    // Read address from vrf.json in contracts directory
    const vrfJSON = JSON.parse(await fs.readFile(vrfJSONPath, 'utf8'));
    const vrfContractAddress = vrfJSON.vrfCoordinatorAddress;

    console.log("VRFCoordinator Address:", vrfContractAddress)

    const vrfCoordinator = MockVRFCoordinator__factory.connect(vrfContractAddress, signer);

    console.log("\nStarting VRFRequestWatcher...");
    console.log("\nListening for RequestRandomWords events...")
    const handleFulfilledRandomness = async (event: FulfillRandomWordsEvent) => {
        console.log("tx: ", event.transactionHash);
    };

    const handleRandomnessRequest = async (event: RequestRandomWordsEvent) => {
        console.log("\nEvent: ", event.args.requestId);
        console.log("> Sender:", event.args.sender);
        console.log("> Nonce:", BigNumber.from(event.args.nonce).toNumber());
        console.log("> OracleId:", event.args.oracleId);
        console.log("> NbWords:", BigNumber.from(event.args.nbWords).toNumber());
        console.log("> CallbackGasLimit:", BigNumber.from(event.args.callbackGasLimit).toNumber());
        console.log("> CallbackAddress:", event.args.callbackAddress);
        console.log("> CallbackSelector:", event.args.callbackSelector);
        console.log("> BlockNumber:", BigNumber.from(event.blockNumber).toNumber())

        // Proof is unused in mock
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

        async function submitTransaction() {
            try {
                console.log("Fulfilling random words...");
                const tx = await vrfCoordinator.fulfillRandomWords(proof, request, {
                    gasLimit: 3000000,
                    gasPrice: ethers.utils.parseUnits('200', 'gwei')
                });
          
              // Handle the transaction success
              // ...
            } catch (error) {
              // Handle revert or exception
              console.error('Error:', error);
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
    
    fetchPastEvents();
    setInterval(fetchPastEvents, 1000); // Fetch new events every second

    // Handle exit signals
    process.on('SIGINT', () => {
        console.log('Server shutting down...');
        process.exit();
    });
}
main();
