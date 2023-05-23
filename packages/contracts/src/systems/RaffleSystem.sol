// SPDX-License-Identifier: MIT
pragma solidity >=0.8.0;

import {System} from "@latticexyz/world/src/System.sol";
import {RaffleCounter, RaffleTable, RequestIdToRaffleId} from "../codegen/Tables.sol";
import {IVRFSystem} from "../codegen/world/IVRFSystem.sol";
import {IRaffleSystem} from "../codegen/world/IRaffleSystem.sol";

contract RaffleSystem is System {
    // Starts a new raffle.
    function startNewRaffle() public returns (uint32) {
        uint32 raffleId = RaffleCounter.get();
        RaffleCounter.set(raffleId + 1);
        RaffleTable.set(raffleId, new address[](0));
        return raffleId;
    }

    // Buy a ticket for a given raffle.
    function buyTicket(uint32 raffleId) public {
        address[] memory tickets = RaffleTable.get(raffleId);
        address[] memory newTickets = new address[](tickets.length + 1);
        for (uint32 i = 0; i < tickets.length; i++) {
            newTickets[i] = tickets[i];
        }
        address randomish = address(uint160(uint256(tickets.length)));
        newTickets[tickets.length] = randomish;
        RaffleTable.set(raffleId, newTickets);
    }

    // End a raffle and ask for randomness.
    function endRaffle(uint32 raffleId) public {
        bytes32 requestId = IVRFSystem(_world()).requestRandomWords(
            bytes32(hex"c1ffd3cfee2d9e5cd67643f8f39fd6e51aad88f6f4ce6ab8827279cfffb92266"),
            0,
            10000000,
            1,
            IRaffleSystem.endRaffleCallback.selector
        );
        RequestIdToRaffleId.set(requestId, raffleId);
    }

    // Off-chain actor will call indirectly into endRaffleCallback with a response for the random number.
    function endRaffleCallback(bytes32 requestId, uint256[] memory randomWords) public returns (uint256) {
        uint32 raffleId = RequestIdToRaffleId.get(requestId);
        address[] memory tickets = RaffleTable.get(raffleId);
        return randomWords[0] % tickets.length;
    }
}
