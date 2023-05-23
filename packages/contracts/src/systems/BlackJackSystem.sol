// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {RaffleCounter, RaffleTable, RequestIdToRaffleId, RequestIdToBlackJackUser, BlackJack} from "../codegen/Tables.sol";
import {IVRFSystem} from "../codegen/world/IVRFSystem.sol";
import {IBlackJackSystem} from "../codegen/world/IBlackJackSystem.sol";

contract BlackJackSystem is System {
    function startGame() public {
        bytes32 oracleId = bytes32(0);
        uint16 requestConfirmations = 0;
        uint32 callbackGasLimit = 10000000;
        uint32 nbWords = 52;
        bytes4 callbackSelector = IBlackJackSystem.handleStartGame.selector;
        bytes32 requestId = IVRFSystem(_world()).requestRandomWords(
            oracleId,
            requestConfirmations,
            callbackGasLimit,
            nbWords,
            callbackSelector
        );    
        RequestIdToBlackJackUser.set(requestId, _msgSender());
    }

    function handleStartGame(bytes32 requestId, uint256[] memory randomWords) public {
        address user = RequestIdToBlackJackUser.get(requestId);

        uint256[] memory userCards = new uint256[](2);
        userCards[0] = randomWords[0] % 52;
        userCards[1] = randomWords[1] % 52;

        uint256[] memory dealerCards = new uint256[](1);
        dealerCards[0] = randomWords[2] % 52;

        uint256 userWins = BlackJack.getUserWins(user);
        uint256 userLosses = BlackJack.getUserLosses(user);
        BlackJack.set(user, userCards, dealerCards, userWins, userLosses);
    }

    function dealUser() public {
        bytes32 oracleId = bytes32(0);
        uint16 requestConfirmations = 0;
        uint32 callbackGasLimit = 10000000;
        uint32 nbWords = 52;
        bytes4 callbackSelector = IBlackJackSystem.handleDealUser.selector;
        bytes32 requestId = IVRFSystem(_world()).requestRandomWords(
            oracleId,
            requestConfirmations,
            callbackGasLimit,
            nbWords,
            callbackSelector
        );    
        RequestIdToBlackJackUser.set(requestId, _msgSender());
    }

    function handleDealUser(bytes32 requestId, uint256[] memory randomWords) public {
        address user = RequestIdToBlackJackUser.get(requestId);
        uint256[] memory userCards = BlackJack.getUserCards(user);
        uint256[] memory newUserCards = new uint256[](userCards.length + 1);
        newUserCards[userCards.length] = randomWords[0] % 52;

        uint256 sum = 0;
        for (uint256 i = 0; i < newUserCards.length; i++) {
            sum += newUserCards[i];
        }
        if (sum > 21) {
            uint256 userLosses = BlackJack.getUserLosses(user);
            BlackJack.setUserLosses(user, userLosses + 1);
        }
        BlackJack.setUserCards(user, newUserCards);
    }
}
