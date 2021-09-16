package token_metadata

/// prefix used for PDAs to avoid certain collision attacks (https://en.wikipedia.org/wiki/Collision_attack#Chosen-prefix_collision_attack)
const PREFIX = "metadata"

/// Used in seeds to make Edition model pda address
const EDITION = "edition"
const RESERVATION = "reservation"
const MAX_NAME_LENGTH = 32
const MAX_SYMBOL_LENGTH = 10
const MAX_URI_LENGTH = 200
const MAX_METADATA_LEN = 1 + 32 + 32 + MAX_DATA_SIZE + 1 + 1 + 9 + 172
const MAX_DATA_SIZE = 4 +
	MAX_NAME_LENGTH +
	4 +
	MAX_SYMBOL_LENGTH +
	4 +
	MAX_URI_LENGTH +
	2 +
	1 +
	4 +
	MAX_CREATOR_LIMIT*MAX_CREATOR_LEN
const MAX_EDITION_LEN = 1 + 32 + 8 + 200
const MAX_MASTER_EDITION_LEN = 1 + 9 + 8 + 264
const MAX_CREATOR_LIMIT = 5
const MAX_CREATOR_LEN = 32 + 1 + 1
const MAX_RESERVATIONS = 200
const MAX_RESERVATION_LIST_V1_SIZE = 1 + 32 + 8 + 8 + MAX_RESERVATIONS*34 + 100
const MAX_RESERVATION_LIST_SIZE = 1 + 32 + 8 + 8 + MAX_RESERVATIONS*48 + 8 + 8 + 84
const MAX_EDITION_MARKER_SIZE = 32
const EDITION_MARKER_BIT_SIZE = 248
