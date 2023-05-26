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
import type { FunctionFragment, Result } from "@ethersproject/abi";
import type { Listener, Provider } from "@ethersproject/providers";
import type {
  TypedEventFilter,
  TypedEvent,
  TypedListener,
  OnEvent,
  PromiseOrValue,
} from "./common";

export interface StorageProofOracleInterface extends utils.Interface {
  functions: {
    "blockHashStore()": FunctionFragment;
    "blockHashes(uint256)": FunctionFragment;
    "getBlockHash(uint256)": FunctionFragment;
    "importBlockHash(uint256)": FunctionFragment;
    "proveBlockHash(uint256)": FunctionFragment;
    "proveBlockHashesAndStateRoots(bytes[])": FunctionFragment;
  };

  getFunction(
    nameOrSignatureOrTopic:
      | "blockHashStore"
      | "blockHashes"
      | "getBlockHash"
      | "importBlockHash"
      | "proveBlockHash"
      | "proveBlockHashesAndStateRoots"
  ): FunctionFragment;

  encodeFunctionData(
    functionFragment: "blockHashStore",
    values?: undefined
  ): string;
  encodeFunctionData(
    functionFragment: "blockHashes",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "getBlockHash",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "importBlockHash",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "proveBlockHash",
    values: [PromiseOrValue<BigNumberish>]
  ): string;
  encodeFunctionData(
    functionFragment: "proveBlockHashesAndStateRoots",
    values: [PromiseOrValue<BytesLike>[]]
  ): string;

  decodeFunctionResult(
    functionFragment: "blockHashStore",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "blockHashes",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "getBlockHash",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "importBlockHash",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "proveBlockHash",
    data: BytesLike
  ): Result;
  decodeFunctionResult(
    functionFragment: "proveBlockHashesAndStateRoots",
    data: BytesLike
  ): Result;

  events: {};
}

export interface StorageProofOracle extends BaseContract {
  connect(signerOrProvider: Signer | Provider | string): this;
  attach(addressOrName: string): this;
  deployed(): Promise<this>;

  interface: StorageProofOracleInterface;

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
    blockHashStore(overrides?: CallOverrides): Promise<[string]>;

    blockHashes(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    getBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<[string]>;

    importBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    proveBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;

    proveBlockHashesAndStateRoots(
      blockHeaders: PromiseOrValue<BytesLike>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<ContractTransaction>;
  };

  blockHashStore(overrides?: CallOverrides): Promise<string>;

  blockHashes(
    arg0: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<string>;

  getBlockHash(
    blockNumber: PromiseOrValue<BigNumberish>,
    overrides?: CallOverrides
  ): Promise<string>;

  importBlockHash(
    blockNumber: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  proveBlockHash(
    blockNumber: PromiseOrValue<BigNumberish>,
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  proveBlockHashesAndStateRoots(
    blockHeaders: PromiseOrValue<BytesLike>[],
    overrides?: Overrides & { from?: PromiseOrValue<string> }
  ): Promise<ContractTransaction>;

  callStatic: {
    blockHashStore(overrides?: CallOverrides): Promise<string>;

    blockHashes(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<string>;

    getBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<string>;

    importBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    proveBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<void>;

    proveBlockHashesAndStateRoots(
      blockHeaders: PromiseOrValue<BytesLike>[],
      overrides?: CallOverrides
    ): Promise<void>;
  };

  filters: {};

  estimateGas: {
    blockHashStore(overrides?: CallOverrides): Promise<BigNumber>;

    blockHashes(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    getBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<BigNumber>;

    importBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    proveBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;

    proveBlockHashesAndStateRoots(
      blockHeaders: PromiseOrValue<BytesLike>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<BigNumber>;
  };

  populateTransaction: {
    blockHashStore(overrides?: CallOverrides): Promise<PopulatedTransaction>;

    blockHashes(
      arg0: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    getBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: CallOverrides
    ): Promise<PopulatedTransaction>;

    importBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    proveBlockHash(
      blockNumber: PromiseOrValue<BigNumberish>,
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;

    proveBlockHashesAndStateRoots(
      blockHeaders: PromiseOrValue<BytesLike>[],
      overrides?: Overrides & { from?: PromiseOrValue<string> }
    ): Promise<PopulatedTransaction>;
  };
}
