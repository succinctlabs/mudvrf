import { useComponentValue, useRow, useRows } from "@latticexyz/react";
import { useMUD } from "./MUDContext";
import { useEffect, useState } from "react";
import { world } from "./mud/world";
import './App.css'

export const App = () => {
  const {
    systemCalls: { startGame, dealUser, standUser },
    network: { storeCache, worldContract },
  } = useMUD();

  const [address, setAddress] = useState("");

  useEffect(() => {
    async function set() {
      setAddress(await worldContract.signer.getAddress());
    }
    set();
  }, [worldContract.signer]);

  const game = useRow(storeCache, {
    table: "BlackJack",
    key: { userAddress: address as any },
  });

  console.log("game", address, game);

  // worldContract.on("StoreSetRecord", (table, key, data) => {
  //   if (
  //     table !==
  //     "0x00000000000000000000000000000000565246526571756573745461626c6556"
  //   ) {
  //     console.log("wrong table", table);
  //     return;
  //   }

  //   const bytes: string = data.slice(2);

  //   let offset = 0;
  //   const msgSender = "0x" + bytes.slice(offset, offset + 40);
  //   offset += 40;

  //   const nonce = BigInt("0x" + bytes.slice(offset, offset + 64));
  //   offset += 64;

  //   const blockNumber = BigInt("0x" + bytes.slice(offset, offset + 64));
  //   offset += 64;

  //   const callbackGasLimit = BigInt("0x" + bytes.slice(offset, offset + 8));
  //   offset += 8;

  //   const nbWords = BigInt("0x" + bytes.slice(offset, offset + 8));
  //   offset += 8;

  //   const callbackSelector = "0x" + bytes.slice(offset, offset + 8);
  //   offset += 8;

  //   // console.log(msgSender, nonce, blockNumber, callbackGasLimit, nbWords, callbackSelector);
  //   setMsgSender(msgSender);
  //   setNonce(nonce);
  //   setBlockNumber(blockNumber);
  //   setCallbackGasLimit(callbackGasLimit);
  //   setNbWords(nbWords);
  //   setCallbackSelector(callbackSelector);
  // });

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
    <div style={{backgroundColor: ""}}>
      <h1>MUDVRF🎲</h1>
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
            {game.value.dealerCards.map((card: any) => getCard(card)).join(", ")}
            ({renderHandTotal(game.value.dealerCards)})
          </div>
          <div>
            Your cards:{" "}
            {game.value.userCards.map((card: any) => getCard(card)).join(", ")} (
            {renderHandTotal(game.value.userCards)})
          </div>
          <div>
            Game Status: {game.value.userWon ? "Win" : "Lose"}
          </div>
          <div>
            Game Ended: {game.value.gameEnded ? "Yes" : "No"}
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
    </div>
    </>
  );
};
