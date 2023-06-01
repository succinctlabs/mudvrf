import { awaitStreamValue } from "@latticexyz/utils";
import { SetupNetworkResult } from "./setupNetwork";

export type SystemCalls = ReturnType<typeof createSystemCalls>;

export function createSystemCalls(
  { worldSend, txReduced$ }: SetupNetworkResult
) {
  const startGame = async () => {
    const tx = await worldSend("startGame", []);
    await tx.wait();
    await awaitStreamValue(txReduced$, (txHash) => txHash === tx.hash);
  }

  const dealUser = async () => {
    const tx = await worldSend("dealUser", []);
    await tx.wait();
    await awaitStreamValue(txReduced$, (txHash) => txHash === tx.hash);
  }

  const standUser = async () => {
    const tx = await worldSend("standUser", []);
    await tx.wait();
    await awaitStreamValue(txReduced$, (txHash) => txHash === tx.hash);
  }

  return {
    startGame,
    dealUser,
    standUser,
  };
}
