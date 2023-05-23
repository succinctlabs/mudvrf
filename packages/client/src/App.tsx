import { useComponentValue, useRow, useRows } from "@latticexyz/react";
import { useMUD } from "./MUDContext";
import { useEffect, useState } from "react";
import { world } from "./mud/world";

export const App = () => {
  const {
    components: { RaffleCounter },
    systemCalls: { startGame, dealUser, standUser },
    network: { storeCache, singletonEntity, worldContract },
  } = useMUD();

  // const raffleId = useComponentValue(RaffleCounter, singletonEntity);

  // const [inputRaffleId, setInputRaffleId] = useState(0);

  // const onChangeHandler = (event: any) => {
  //   setInputRaffleId(event.target.value);
  // };

  const [address, setAddress] = useState("");

  useEffect(() => {
    async function set() {
      setAddress(await worldContract.signer.getAddress());
    }
    set();
  }, [worldContract.signer]);

  // const blackjackGames = useRows(storeCache, { table: "BlackJack" });
  // const [game, setGame] = useState<typeof blackjackGames[number] | null>(null);
  // useEffect(() => {
  //   if (blackjackGames.length > 0) {
  //     for (const game of blackjackGames) {
  //       if (game.key.userAddress === address) {
  //         setGame(game);
  //         return;
  //       }
  //     }
  //   }
  //   setGame(null);
  // }, [blackjackGames]);

  // const entity = world.registerEntity({ id: address });

  const game = useRow(storeCache, {
    table: "BlackJack",
    key: { userAddress: address as any },
  });

  console.log(address, game);

  // if (blackjackGames.length > 0) {
  //   console.log(blackjackGames[0].value.userCards);
  // }
  // console.log("all games", blackjackGames);

  const onProofChangeHandler = (event: any) => {
    console.log(JSON.parse(event.target.value));
  };

  const onRequestChangeHandler = (event: any) => {
    console.log(JSON.parse(event.target.value));
  };

  const [msgSender, setMsgSender] = useState("");
  const [nonce, setNonce] = useState(BigInt(0));
  const [blockNumber, setBlockNumber] = useState(BigInt(0));
  const [callbackGasLimit, setCallbackGasLimit] = useState(BigInt(0));
  const [nbWords, setNbWords] = useState(BigInt(0));
  const [callbackSelector, setCallbackSelector] = useState("");

  worldContract.on("StoreSetRecord", (table, key, data) => {
    if (
      table !==
      "0x00000000000000000000000000000000565246526571756573745461626c6556"
    ) {
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

  const CARDS = [
    "A",
    "2",
    "3",
    "4",
    "5",
    "6",
    "7",
    "8",
    "9",
    "10",
    "J",
    "Q",
    "K",
  ];
  const SUITS = ["♠", "♣", "♥", "♦"];

  function getCard(card: bigint) {
    return CARDS[Number(card % 13n)] + SUITS[Number((card / 13n) % 4n)];
  }

  function getHandTotal(cards: bigint[]) {
    let total = 0;
    let aces = false;
    for (const card of cards) {
      const cardValue = Number(card % 13n);
      if (cardValue === 0) {
        aces = true;
        total += 1;
      } else if (cardValue >= 9) {
        total += 10;
      } else {
        total += cardValue + 1;
      }
    }
    return [total, aces] as const;
  }

  function renderHandTotal(cards: bigint[]) {
    const [total, aces] = getHandTotal(cards);
    if (aces && total <= 11) {
      return total + 10;
    } else {
      return total;
    }
  }

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
      {/* <button onClick={startNewRaffle}>Start New Raffle</button> */}
      <br></br>
      <button
        onClick={() => startGame()}
        // disabled={game != null}
      >
        Start game
      </button>
      <h2>CURRENT GAME:</h2>
      {game === undefined ||
      game === null ||
      game.value.userCards === undefined ? (
        <div>NO GAME</div>
      ) : (
        <div key={game.key.userAddress}>
          <div>
            Dealer cards:{" "}
            {game.value.dealerCards.map((card) => getCard(card)).join(", ")}
          </div>
          <div>
            Your cards:{" "}
            {game.value.userCards.map((card) => getCard(card)).join(", ")} (
            {renderHandTotal(game.value.userCards)})
          </div>
          <br></br>
          <button onClick={dealUser}>Hit</button>
          <button onClick={standUser} style={{ marginLeft: "6px" }}>
            Stand
          </button>
        </div>
      )}
      <br></br>
      <br></br>
      {/* <input
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
      ))} */}
      {/* VRF Proof JSON
      <input type="text" name="VRF Proof JSON" onChange={onProofChangeHandler} />
      <br></br>
      VRF Request JSON
      <input type="text" name="VRF Request JSON" onChange={onRequestChangeHandler} />
      <br></br>
      <br></br>
      <button>VRF Coordinator Callback</button> */}
    </>
  );
};
