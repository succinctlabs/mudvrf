"use strict";
var __assign = (this && this.__assign) || function () {
    __assign = Object.assign || function(t) {
        for (var s, i = 1, n = arguments.length; i < n; i++) {
            s = arguments[i];
            for (var p in s) if (Object.prototype.hasOwnProperty.call(s, p))
                t[p] = s[p];
        }
        return t;
    };
    return __assign.apply(this, arguments);
};
var __awaiter = (this && this.__awaiter) || function (thisArg, _arguments, P, generator) {
    function adopt(value) { return value instanceof P ? value : new P(function (resolve) { resolve(value); }); }
    return new (P || (P = Promise))(function (resolve, reject) {
        function fulfilled(value) { try { step(generator.next(value)); } catch (e) { reject(e); } }
        function rejected(value) { try { step(generator["throw"](value)); } catch (e) { reject(e); } }
        function step(result) { result.done ? resolve(result.value) : adopt(result.value).then(fulfilled, rejected); }
        step((generator = generator.apply(thisArg, _arguments || [])).next());
    });
};
var __generator = (this && this.__generator) || function (thisArg, body) {
    var _ = { label: 0, sent: function() { if (t[0] & 1) throw t[1]; return t[1]; }, trys: [], ops: [] }, f, y, t, g;
    return g = { next: verb(0), "throw": verb(1), "return": verb(2) }, typeof Symbol === "function" && (g[Symbol.iterator] = function() { return this; }), g;
    function verb(n) { return function (v) { return step([n, v]); }; }
    function step(op) {
        if (f) throw new TypeError("Generator is already executing.");
        while (g && (g = 0, op[0] && (_ = 0)), _) try {
            if (f = 1, y && (t = op[0] & 2 ? y["return"] : op[0] ? y["throw"] || ((t = y["return"]) && t.call(y), 0) : y.next) && !(t = t.call(y, op[1])).done) return t;
            if (y = 0, t) op = [op[0] & 2, t.value];
            switch (op[0]) {
                case 0: case 1: t = op; break;
                case 4: _.label++; return { value: op[1], done: false };
                case 5: _.label++; y = op[1]; op = [0]; continue;
                case 7: op = _.ops.pop(); _.trys.pop(); continue;
                default:
                    if (!(t = _.trys, t = t.length > 0 && t[t.length - 1]) && (op[0] === 6 || op[0] === 2)) { _ = 0; continue; }
                    if (op[0] === 3 && (!t || (op[1] > t[0] && op[1] < t[3]))) { _.label = op[1]; break; }
                    if (op[0] === 6 && _.label < t[1]) { _.label = t[1]; t = op; break; }
                    if (t && _.label < t[2]) { _.label = t[2]; _.ops.push(op); break; }
                    if (t[2]) _.ops.pop();
                    _.trys.pop(); continue;
            }
            op = body.call(thisArg, _);
        } catch (e) { op = [6, e]; y = 0; } finally { f = t = 0; }
        if (op[0] & 5) throw op[1]; return { value: op[0] ? op[1] : void 0, done: true };
    }
};
exports.__esModule = true;
var MockVRFCoordinator__Factory_1 = require("./types/ethers-contracts/factories/MockVRFCoordinator__Factory");
var ethers_1 = require("ethers");
var ethers_2 = require("ethers");
function main() {
    return __awaiter(this, void 0, void 0, function () {
        var provider, wallet, signer, contractAddress, vrfCoordinator, handleFulfilledRandomness, handleRandomnessRequest;
        var _this = this;
        return __generator(this, function (_a) {
            console.log("Listening for randomness requests...");
            provider = new ethers_1.ethers.providers.JsonRpcProvider("http://localhost:8545");
            wallet = new ethers_1.ethers.Wallet("0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80", provider);
            signer = wallet.connect(provider);
            contractAddress = "0xc3e53F4d16Ae77Db1c982e75a937B9f60FE63690";
            vrfCoordinator = MockVRFCoordinator__Factory_1.MockVRFCoordinator__factory.connect(contractAddress, provider);
            handleFulfilledRandomness = function (nonce, requestId, words) { return __awaiter(_this, void 0, void 0, function () {
                return __generator(this, function (_a) {
                    console.log("Randomness fulfilled");
                    console.log("nonce: ", ethers_2.BigNumber.from(nonce).toBigInt());
                    console.log("requestId: ", ethers_2.BigNumber.from(requestId).toHexString());
                    console.log("words: ", (words).map(function (word) { return ethers_2.BigNumber.from(word).toHexString(); }));
                    return [2 /*return*/];
                });
            }); };
            handleRandomnessRequest = function (nonce, requestId, sender, blockNumber, seed, callbackGasLimit, nbWords, callbackSelector) { return __awaiter(_this, void 0, void 0, function () {
                function submitTransaction() {
                    return __awaiter(this, void 0, void 0, function () {
                        var txData, txOptions, transaction, receipt, error_1;
                        return __generator(this, function (_a) {
                            switch (_a.label) {
                                case 0:
                                    _a.trys.push([0, 3, , 4]);
                                    txData = vrfCoordinator.connect(signer).fulfillRandomWords(proof, request);
                                    txOptions = { gasLimit: 300000, gasPrice: ethers_1.ethers.utils.parseUnits('20', 'gwei') };
                                    return [4 /*yield*/, signer.sendTransaction(__assign(__assign({}, txData), txOptions))];
                                case 1:
                                    transaction = _a.sent();
                                    console.log('Transaction submitted:', transaction.hash);
                                    return [4 /*yield*/, transaction.wait()];
                                case 2:
                                    receipt = _a.sent();
                                    console.log('Transaction mined:', receipt.transactionHash);
                                    return [3 /*break*/, 4];
                                case 3:
                                    error_1 = _a.sent();
                                    // Handle revert or exception
                                    console.error('Transaction failed:', error_1);
                                    return [2 /*return*/];
                                case 4: return [2 /*return*/];
                            }
                        });
                    });
                }
                var proof, oracleId, request;
                return __generator(this, function (_a) {
                    // Log all of the parameters
                    console.log("Randomness request received");
                    proof = {
                        pk: [requestId, requestId],
                        gamma: [requestId, requestId],
                        c: requestId,
                        s: requestId,
                        seed: requestId,
                        uWitness: sender,
                        cGammaWitness: [requestId, requestId],
                        sHashWitness: [requestId, requestId],
                        zInv: requestId
                    };
                    oracleId = ethers_1.utils.keccak256(ethers_1.utils.defaultAbiCoder.encode(["address", "uint256"], [sender, nonce]));
                    request = {
                        sender: sender,
                        nonce: nonce,
                        blockNumber: blockNumber,
                        oracleId: oracleId,
                        requestConfirmations: 0,
                        callbackGasLimit: callbackGasLimit,
                        nbWords: nbWords,
                        callbackSelector: callbackSelector
                    };
                    console.log("Callback Selector: ", callbackSelector);
                    console.log("nbWords: ", ethers_2.BigNumber.from(nbWords).toNumber());
                    submitTransaction();
                    console.log("Submitted fulfillRandomWords transaction");
                    return [2 /*return*/];
                });
            }); };
            vrfCoordinator.on("RandomnessRequest", handleRandomnessRequest);
            vrfCoordinator.on("FulfilledRandomness", handleFulfilledRandomness);
            // Handle exit signals
            process.on('SIGINT', function () {
                vrfCoordinator.removeAllListeners('RandomnessRequested');
                vrfCoordinator.removeAllListeners('FulfilledRandomness');
                console.log('Server shutting down...');
                process.exit();
            });
            return [2 /*return*/];
        });
    });
}
main();
