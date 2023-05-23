import { useComponentValue, useRows } from "@latticexyz/react";
import { useMUD } from "./MUDContext";
import { useState } from "react";

export const App = () => {
  const {
    components: { RaffleCounter },
    systemCalls: { startNewRaffle, buyTicket, endRaffle },
    network: { storeCache, singletonEntity },
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

  return (
    <>
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
