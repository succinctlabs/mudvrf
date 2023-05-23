import { useComponentValue, useRows } from "@latticexyz/react";
import { useMUD } from "./MUDContext";
import { useState } from "react";

export const App = () => {
  const {
    components: { RaffleCounter },
    systemCalls: { startNewRaffle, buyTicket, endRaffle },
    network: { storeCache, singletonEntity, worldContract },
  } = useMUD();

  const raffleId = useComponentValue(RaffleCounter, singletonEntity);
  const raffles = useRows(storeCache, { table: "RaffleTable" });

  const [inputRaffleId, setInputRaffleId] = useState(0);

  const onChangeHandler = (event: any) => {
    setInputRaffleId(event.target.value);
  };

  const onProofChangeHandler = (event: any) => {
    console.log(JSON.parse(event.target.value));
  }

  const onRequestChangeHandler = (event: any) => {
    console.log(JSON.parse(event.target.value));
  }

  const [msgSender, setMsgSender] = useState("");
  const [nonce, setNonce] = useState(BigInt(0));
  const [blockNumber, setBlockNumber] = useState(BigInt(0));
  const [callbackGasLimit, setCallbackGasLimit] = useState(BigInt(0));
  const [nbWords, setNbWords] = useState(BigInt(0));
  const [callbackSelector, setCallbackSelector] = useState("");

  worldContract.on("StoreSetRecord", (table, key, data) => {
    if (table !== '0x00000000000000000000000000000000565246526571756573745461626c6556') {
      console.log("wrong table", table);
      return;
    }

    const bytes: string = data.slice(2);

    let offset = 0;
    const msgSender = "0x" + bytes.slice(offset, offset + 40);
    offset += 40;

    const nonce = BigInt("0x" + bytes.slice(offset, offset + 64));
    offset += 64;

    const blockNumber = BigInt("0x" + bytes.slice(offset, offset + 64));
    offset += 64;

    const callbackGasLimit = BigInt("0x" + bytes.slice(offset, offset + 8));
    offset += 8;

    const nbWords = BigInt("0x" + bytes.slice(offset, offset + 8));
    offset += 8;

    const callbackSelector = "0x" + bytes.slice(offset, offset + 8);
    offset += 8;

    // console.log(msgSender, nonce, blockNumber, callbackGasLimit, nbWords, callbackSelector);
    setMsgSender(msgSender);
    setNonce(nonce);
    setBlockNumber(blockNumber);
    setCallbackGasLimit(callbackGasLimit);
    setNbWords(nbWords);
    setCallbackSelector(callbackSelector);
  });

  return (
    <>
      Latest VRF Request
      <br></br>
      <p>msgSender: {msgSender}</p>
      <p>nonce: {nonce.toString()}</p>
      <p>blockNumber: {blockNumber.toString()}</p>
      <p>callbackGasLimit: {callbackGasLimit.toString()}</p>
      <p>nbWords: {nbWords.toString()}</p>
      <p>callbackSelector: {callbackSelector.toString()}</p>
      <button onClick={startNewRaffle}>Start New Raffle</button>
      <br></br>
      <input
        type="text"
        name="name"
        onChange={onChangeHandler}
        value={inputRaffleId}
      />
      <button onClick={() => {buyTicket(inputRaffleId)}}>Buy Raffle Ticket</button>
      <button onClick={() => {endRaffle(inputRaffleId)}}>End Raffle</button>
      <p>Raffle Id: {raffleId?.value}</p>
      {raffles.map((raffle) => (
        <p>
          ID: {raffle.key.raffleId} Tickets: [{raffle.value.tickets.join(",")}]
        </p>
      ))}
      VRF Proof JSON
      <input type="text" name="VRF Proof JSON" onChange={onProofChangeHandler} />
      <br></br>
      VRF Request JSON
      <input type="text" name="VRF Request JSON" onChange={onRequestChangeHandler} />
      <br></br>
      <br></br>
      <button>VRF Coordinator Callback</button>
    </>
  );
};
