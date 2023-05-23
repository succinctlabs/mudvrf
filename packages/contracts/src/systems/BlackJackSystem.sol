// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {RaffleCounter, RaffleTable, RequestIdToRaffleId, RequestIdToBlackJackUser, BlackJack} from "../codegen/Tables.sol";
import {IVRFSystem} from "../codegen/world/IVRFSystem.sol";
import {IBlackJackSystem} from "../codegen/world/IBlackJackSystem.sol";

contract BlackJackSystem is System {
    bytes32 public constant ORACLE_ID = bytes32(hex"c1ffd3cfee2d9e5cd67643f8f39fd6e51aad88f6f4ce6ab8827279cfffb92266");
    uint16 public constant REQUEST_CONFIRMATIONS = 0;
    uint32 public constant CALLBACK_GAS_LIMIT = 10000000;
    uint32 public constant NB_WORDS = 52;

    function requestRandomness(bytes4 selector) internal returns (bytes32) {
        bytes32 requestId = IVRFSystem(_world()).requestRandomWords(
            ORACLE_ID,
            REQUEST_CONFIRMATIONS,
            CALLBACK_GAS_LIMIT,
            NB_WORDS,
            selector
        );
        RequestIdToBlackJackUser.set(requestId, _msgSender());
        return requestId;
    }

    function calculateValueOfCards(uint256[] memory cards) internal pure returns (uint256[] memory) {
        uint256[] memory sums = new uint256[](2);
        bool seenAce = false;

        for (uint256 i = 0; i < cards.length; i++) {
            uint256 value = 0;
            if (cards[i] % 13 > 9) {
                value = 10;
            } else {
                value = 1 + cards[i] % 13;
            }
            sums[0] += value;

            if (cards[i] % 13 == 0) {
                seenAce = true;
            }
        }

        if (seenAce) {
            sums[1] = sums[0] + 10;
        } else {
            sums[1] = sums[0];
        }

        return sums;
    }


    function startGame() public {
        requestRandomness(IBlackJackSystem.handleStartGame.selector);
    }

    function handleStartGame(bytes32 requestId, uint256[] memory randomWords) public {
        address user = RequestIdToBlackJackUser.get(requestId);

        uint256[] memory userCards = new uint256[](2);
        userCards[0] = randomWords[0] % 52;
        userCards[1] = randomWords[1] % 52;

        uint256[] memory valueOfCards = calculateValueOfCards(userCards);
        if (valueOfCards[0] == 21 || valueOfCards[1] == 21) {
            uint256 userWins = BlackJack.getUserWins(user);
            BlackJack.setUserWins(user, userWins + 1);
            BlackJack.setGameEnded(user, true);
            return;
        }

        uint256[] memory dealerCards = new uint256[](1);
        dealerCards[0] = randomWords[2] % 52;

        uint256 userWins = BlackJack.getUserWins(user);
        uint256 userLosses = BlackJack.getUserLosses(user);
        BlackJack.set(user, userWins, userLosses, userCards, dealerCards, false);
    }

    function dealUser() public {
        requestRandomness(IBlackJackSystem.handleStartGame.selector);
    }

    function handleDealUser(bytes32 requestId, uint256[] memory randomWords) public {
        address user = RequestIdToBlackJackUser.get(requestId);
        uint256[] memory userCards = BlackJack.getUserCards(user);
        uint256[] memory newUserCards = new uint256[](userCards.length + 1);
        newUserCards[userCards.length] = randomWords[0] % 52;
        BlackJack.setUserCards(user, newUserCards);

        uint256[] memory valueOfCards = calculateValueOfCards(newUserCards);
        if (valueOfCards[0] > 21 && valueOfCards[1] > 21) {
            uint256 userLosses = BlackJack.getUserLosses(user);
            BlackJack.setUserLosses(user, userLosses + 1);
        } else if (valueOfCards[0] == 21 || valueOfCards[1] == 21) {
            uint256 userWins = BlackJack.getUserWins(user);
            BlackJack.setUserWins(user, userWins + 1);
        }

        BlackJack.setGameEnded(user, true);
    }

    function standUser() public {
        requestRandomness(IBlackJackSystem.handleStartGame.selector);
    }

    function handleStandUser(bytes32 requestId, uint256[] memory randomWords) public {
        address user = RequestIdToBlackJackUser.get(requestId); 
        uint256[] memory userCards = BlackJack.getUserCards(user);  
        uint256[] memory dealerCards = BlackJack.getDealerCards(user);

        uint256[] memory newDealerCards = new uint256[](dealerCards.length + 1);
        newDealerCards[dealerCards.length] = randomWords[0] % 52;
        BlackJack.setDealerCards(user, newDealerCards);

        uint256[] memory valueOfUserCards = calculateValueOfCards(userCards);
        uint256[] memory valueOfDealerCards = calculateValueOfCards(newDealerCards);

        // Tie at 21.
        if ((valueOfUserCards[0] == 21 || valueOfUserCards[1] == 21) && (valueOfDealerCards[0] == 21 || valueOfDealerCards[1] == 21)) {
            BlackJack.setGameEnded(user, true);
            return;
        }

        // Tie at greater than 21.
        if ((valueOfUserCards[0] > 21 && valueOfUserCards[1] > 21) && (valueOfDealerCards[0] > 21 && valueOfDealerCards[1] > 21)) {
            BlackJack.setGameEnded(user, true);
            return;
        }

        // User goes over 21.
        if (valueOfUserCards[0] > 21 && valueOfUserCards[1] > 21) {
            uint256 userLosses = BlackJack.getUserLosses(user);
            BlackJack.setUserLosses(user, userLosses + 1);
            BlackJack.setGameEnded(user, true);
            return;
        }

        // Opponent goes over 21.
        if (valueOfDealerCards[0] > 21 && valueOfDealerCards[1] > 21) {
            uint256 userWins = BlackJack.getUserWins(user);
            BlackJack.setUserWins(user, userWins + 1);
            BlackJack.setGameEnded(user, true);
            return;
        }

        // Find users best card under 21.
        uint256 userBestCard = 0;
        if (valueOfUserCards[0] >= valueOfUserCards[1] && valueOfUserCards[0] <= 21) {
            userBestCard = valueOfUserCards[0];
        } else if (valueOfUserCards[1] >= valueOfUserCards[0] && valueOfUserCards[1] <= 21) {
            userBestCard = valueOfUserCards[1];
        }

        // Find dealers best card under 21.
        uint256 dealerBestCard = 0;
        if (valueOfDealerCards[0] >= valueOfDealerCards[1] && valueOfDealerCards[0] <= 21) {
            dealerBestCard = valueOfDealerCards[0];
        } else if (valueOfDealerCards[1] >= valueOfDealerCards[0] && valueOfDealerCards[1] <= 21) {
            dealerBestCard = valueOfDealerCards[1];
        }   

        // Evaluate result.
        if (userBestCard > dealerBestCard) {
            uint256 userWins = BlackJack.getUserWins(user);
            BlackJack.setUserWins(user, userWins + 1);
        } else if (userBestCard < dealerBestCard) {
            uint256 userLosses = BlackJack.getUserLosses(user);
            BlackJack.setUserLosses(user, userLosses + 1);
        }
        BlackJack.setGameEnded(user, true);
    }
}
