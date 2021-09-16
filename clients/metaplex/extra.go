package metaplex

/// prefix used for PDAs to avoid certain collision attacks (https://en.wikipedia.org/wiki/Collision_attack#Chosen-prefix_collision_attack)
const PREFIX = "metaplex"
const TOTALS = "totals"
const BASE_TRACKER_SIZE = 1 + 1 + 1 + 4
const MAX_AUCTION_MANAGER_V2_SIZE = 1 + 32 + 32 + 32 + 32 + 32 + 1 + 1 + 8 + 200
const MAX_STORE_SIZE = 2 + 32 + 32 + 32 + 32 + 100
const MAX_WHITELISTED_CREATOR_SIZE = 2 + 32 + 10
const MAX_PAYOUT_TICKET_SIZE = 1 + 32 + 8
const MAX_BID_REDEMPTION_TICKET_SIZE = 3
const MAX_AUTHORITY_LOOKUP_SIZE = 33
const MAX_PRIZE_TRACKING_TICKET_SIZE = 1 + 32 + 8 + 8 + 8 + 50
const BASE_SAFETY_CONFIG_SIZE = 1 + 32 + 8 + 1 + 1 + 1 + 4 + 1 + 1 + 1 + 9 + 1 + 8 + 20

const MAX_WINNERS = 200
const MAX_WINNER_SIZE = 6 * MAX_WINNERS
const MAX_AUCTION_MANAGER_V1_SIZE = 1 +
    32 +
    32 +
    32 +
    32 +
    32 +
    1 +
    1 +
    8 +
    8 +
    MAX_WINNER_SIZE +
    1 +
    8 +
    1 +
    1 +
    32 +
    8 +
    8 +
    1 +
    1 +
    1 +
    1 +
    9 +
    1 +
    AUCTION_MANAGER_PADDING

const AUCTION_MANAGER_PADDING = 149
const MAX_VALIDATION_TICKET_SIZE = 1 + 32 + 10
const MAX_WINNING_CONFIG_STATE_ITEM_SIZE = 2

const ORDER_POSITION = 33
const AUCTION_MANAGER_POSITION = 1
const WINNING_CONFIG_POSITION = 41
const AMOUNT_POSITION = 42
const LENGTH_POSITION = 43
const AMOUNT_RANGE_SIZE_POSITION = 44
const AMOUNT_RANGE_FIRST_EL_POSITION = 48
