/* Autogenerated file. Do not edit manually. */
/* tslint:disable */
/* eslint-disable */

import { Contract, Signer, utils } from "ethers";
import type { Provider } from "@ethersproject/providers";
import type { SliceLib, SliceLibInterface } from "../SliceLib";

const _abi = [
  {
    inputs: [
      {
        internalType: "bytes",
        name: "data",
        type: "bytes",
      },
      {
        internalType: "uint256",
        name: "start",
        type: "uint256",
      },
      {
        internalType: "uint256",
        name: "end",
        type: "uint256",
      },
    ],
    name: "Slice_OutOfBounds",
    type: "error",
  },
] as const;

export class SliceLib__factory {
  static readonly abi = _abi;
  static createInterface(): SliceLibInterface {
    return new utils.Interface(_abi) as SliceLibInterface;
  }
  static connect(
    address: string,
    signerOrProvider: Signer | Provider
  ): SliceLib {
    return new Contract(address, _abi, signerOrProvider) as SliceLib;
  }
}
