/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */
import type {
  BaseContract,
  BigNumber,
  BigNumberish,
  BytesLike,
  CallOverrides,
  ContractTransaction,
  Overrides,
  PopulatedTransaction,
  Signer,
  utils,
} from "ethers";
import type {
  FunctionFragment,
  Result,
  EventFragment,
} from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "./common";

export declare namespace VRF {
  export type ProofStruct = {
    pk: [PromiseOrValue<BigNumberish>, PromiseOrValue<BigNumberish>];
    gamma: [PromiseOrValue<BigNumberish>, PromiseOrValue<BigNumberish>];
    c: PromiseOrValue<BigNumberish>;
    s: PromiseOrValue<BigNumberish>;
    seed: PromiseOrValue<BigNumberish>;
    uWitness: PromiseOrValue<string>;
    cGammaWitness: [PromiseOrValue<BigNumberish>, PromiseOrValue<BigNumberish>];
    sHashWitness: [PromiseOrValue<BigNumberish>, PromiseOrValue<BigNumberish>];
    zInv: PromiseOrValue<BigNumberish>;
  };

  export type ProofStructOutput = [
    [BigNumber, BigNumber],
    [BigNumber, BigNumber],
    BigNumber,
    BigNumber,
    BigNumber,
    string,
    [BigNumber, BigNumber],
    [BigNumber, BigNumber],
    BigNumber
  ] & {
    pk: [BigNumber, BigNumber];
    gamma: [BigNumber, BigNumber];
    c: BigNumber;
    s: BigNumber;
    seed: BigNumber;
    uWitness: string;
    cGammaWitness: [BigNumber, BigNumber];
    sHashWitness: [BigNumber, BigNumber];
    zInv: BigNumber;
  };

  export type RequestStruct = {
    sender: PromiseOrValue<string>;
    nonce: PromiseOrValue<BigNumberish>;
    blockNumber: PromiseOrValue<BigNumberish>;
    oracleId: PromiseOrValue<BytesLike>;
    requestConfirmations: PromiseOrValue<BigNumberish>;
    callbackGasLimit: PromiseOrValue<BigNumberish>;
    nbWords: PromiseOrValue<BigNumberish>;
    callbackSelector: PromiseOrValue<BytesLike>;
  };

  export type RequestStructOutput = [
    string,
    BigNumber,
    BigNumber,
    string,
    number,
    number,
    number,
    string
  ] & {
    sender: string;
    nonce: BigNumber;
    blockNumber: BigNumber;
    oracleId: string;
    requestConfirmations: number;
    callbackGasLimit: number;
    nbWords: number;
    callbackSelector: string;
  };
}

export interface MockVRFCoordinatorInterface extends utils.Interface {
  functions: {
    "fulfillRandomWords((uint256[2],uint256[2],uint256,uint256,uint256,address,uint256[2],uint256[2],uint256),(address,uint256,uint64,bytes32,uint16,uint32,uint32,bytes4))": FunctionFragment;
    "nonce()": FunctionFragment;
    "oracles(bytes32)": FunctionFragment;
    "requestRandomWords(bytes32,uint16,uint32,uint32,bytes4)": FunctionFragment;
    "requests(bytes32)": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "fulfillRandomWords"
      | "nonce"
      | "oracles"
      | "requestRandomWords"
      | "requests"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "fulfillRandomWords",
    values: [VRF.ProofStruct, VRF.RequestStruct]
  ): string;
  encodeFunctionData(functionFragment: "nonce", values?: undefined): string;
  encodeFunctionData(
    functionFragment: "oracles",
    values: [PromiseOrValue<BytesLike>]
  ): string;
  encodeFunctionData(
    functionFragment: "requestRandomWords",
    values: [
      PromiseOrValue<BytesLike>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BigNumberish>,
      PromiseOrValue<BytesLike>
    ]
  ): string;
  encodeFunctionData(
    functionFragment: "requests",
    values: [PromiseOrValue<BytesLike>]
  ): string;

  decodeFunctionResult(
    functionFragment: "fulfillRandomWords",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "nonce", data: BytesLike): Result;
  decodeFunctionResult(functionFragment: "oracles", data: BytesLike): Result;
  decodeFunctionResult(
    functionFragment: "requestRandomWords",
    data: BytesLike
  ): Result;
  decodeFunctionResult(functionFragment: "requests", data: BytesLike): Result;

  events: {
    "FulfilledRandomness(uint256,bytes32,uint256[])": EventFragment;
    "RandomnessRequest(uint256,bytes32,address,uint256,bytes32,uint32,uint64,bytes4)": EventFragment;
  };

  getEvent(nameOrSignatureOrTopic: "FulfilledRandomness"): EventFragment;
  getEvent(nameOrSignatureOrTopic: "RandomnessRequest"): EventFragment;
}

export interface FulfilledRandomnessEventObject {
  nonce: BigNumber;
  requestId: string;
  words: BigNumber[];
}
export type FulfilledRandomnessEvent = TypedEvent<
  [BigNumber, string, BigNumber[]],
  FulfilledRandomnessEventObject
>;

export type FulfilledRandomnessEventFilter =
  TypedEventFilter<FulfilledRandomnessEvent>;

export interface RandomnessRequestEventObject {
  nonce: BigNumber;
  requestId: string;
  sender: string;
  blockNumber: BigNumber;
  seed: string;
  callbackGasLimit: number;
  nbWords: BigNumber;
  callbackSelector: string;
}
export type RandomnessRequestEvent = TypedEvent<
  [BigNumber, string, string, BigNumber, string, number, BigNumber, string],
  RandomnessRequestEventObject
>;

export type RandomnessRequestEventFilter =
  TypedEventFilter<RandomnessRequestEvent>;

export interface MockVRFCoordinator extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: MockVRFCoordinatorInterface;

  queryFilter<TEvent extends TypedEvent>(
    event: TypedEventFilter<TEvent>,
    fromBlockOrBlockhash?: string | number | undefined,
    toBlock?: string | number | undefined
  ): Promise<Array<TEvent>>;

  listeners<TEvent extends TypedEvent>(
    eventFilter?: TypedEventFilter<TEvent>
  ): Array<TypedListener<TEvent>>;
  listeners(eventName?: string): Array<Listener>;
  removeAllListeners<TEvent extends TypedEvent>(
    eventFilter: TypedEventFilter<TEvent>
  ): this;
  removeAllListeners(eventName?: string): this;
  off: OnEvent<this>;
  on: OnEvent<this>;
  once: OnEvent<this>;
  removeListener: OnEvent<this>;

  functions: {
    fulfillRandomWords(
      _proof: VRF.ProofStruct,
      _request: VRF.RequestStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    nonce(overrides?: CallOverrides): Promise<[BigNumber]>;

    oracles(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    requestRandomWords(
      _oracleId: PromiseOrValue<BytesLike>,
      _requestConfirmations: PromiseOrValue<BigNumberish>,
      _callbackGasLimit: PromiseOrValue<BigNumberish>,
      _nbWords: PromiseOrValue<BigNumberish>,
      _callbackSelector: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    requests(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<[string]>;
  };

  fulfillRandomWords(
    _proof: VRF.ProofStruct,
    _request: VRF.RequestStruct,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  nonce(overrides?: CallOverrides): Promise<BigNumber>;

  oracles(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  requestRandomWords(
    _oracleId: PromiseOrValue<BytesLike>,
    _requestConfirmations: PromiseOrValue<BigNumberish>,
    _callbackGasLimit: PromiseOrValue<BigNumberish>,
    _nbWords: PromiseOrValue<BigNumberish>,
    _callbackSelector: PromiseOrValue<BytesLike>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  requests(
    arg0: PromiseOrValue<BytesLike>,
    overrides?: CallOverrides
  ): Promise<string>;

  callStatic: {
    fulfillRandomWords(
      _proof: VRF.ProofStruct,
      _request: VRF.RequestStruct,
      overrides?: CallOverrides
    ): Promise<void>;

    nonce(overrides?: CallOverrides): Promise<BigNumber>;

    oracles(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    requestRandomWords(
      _oracleId: PromiseOrValue<BytesLike>,
      _requestConfirmations: PromiseOrValue<BigNumberish>,
      _callbackGasLimit: PromiseOrValue<BigNumberish>,
      _nbWords: PromiseOrValue<BigNumberish>,
      _callbackSelector: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;

    requests(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<string>;
  };

  filters: {
    "FulfilledRandomness(uint256,bytes32,uint256[])"(
      nonce?: PromiseOrValue<BigNumberish> | null,
      requestId?: PromiseOrValue<BytesLike> | null,
      words?: null
    ): FulfilledRandomnessEventFilter;
    FulfilledRandomness(
      nonce?: PromiseOrValue<BigNumberish> | null,
      requestId?: PromiseOrValue<BytesLike> | null,
      words?: null
    ): FulfilledRandomnessEventFilter;

    "RandomnessRequest(uint256,bytes32,address,uint256,bytes32,uint32,uint64,bytes4)"(
      nonce?: PromiseOrValue<BigNumberish> | null,
      requestId?: PromiseOrValue<BytesLike> | null,
      sender?: null,
      blockNumber?: null,
      seed?: null,
      callbackGasLimit?: null,
      nbWords?: null,
      callbackSelector?: null
    ): RandomnessRequestEventFilter;
    RandomnessRequest(
      nonce?: PromiseOrValue<BigNumberish> | null,
      requestId?: PromiseOrValue<BytesLike> | null,
      sender?: null,
      blockNumber?: null,
      seed?: null,
      callbackGasLimit?: null,
      nbWords?: null,
      callbackSelector?: null
    ): RandomnessRequestEventFilter;
  };

  estimateGas: {
    fulfillRandomWords(
      _proof: VRF.ProofStruct,
      _request: VRF.RequestStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    nonce(overrides?: CallOverrides): Promise<BigNumber>;

    oracles(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    requestRandomWords(
      _oracleId: PromiseOrValue<BytesLike>,
      _requestConfirmations: PromiseOrValue<BigNumberish>,
      _callbackGasLimit: PromiseOrValue<BigNumberish>,
      _nbWords: PromiseOrValue<BigNumberish>,
      _callbackSelector: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    requests(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    fulfillRandomWords(
      _proof: VRF.ProofStruct,
      _request: VRF.RequestStruct,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    nonce(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    oracles(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    requestRandomWords(
      _oracleId: PromiseOrValue<BytesLike>,
      _requestConfirmations: PromiseOrValue<BigNumberish>,
      _callbackGasLimit: PromiseOrValue<BigNumberish>,
      _nbWords: PromiseOrValue<BigNumberish>,
      _callbackSelector: PromiseOrValue<BytesLike>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    requests(
      arg0: PromiseOrValue<BytesLike>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;
  };
}
