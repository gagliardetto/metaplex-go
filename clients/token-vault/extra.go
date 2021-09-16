package token_vault

/// prefix used for PDAs to avoid certain collision attacks (https://en.wikipedia.org/wiki/Collision_attack#Chosen-prefix_collision_attack)
const PREFIX = "vault"

const MAX_SAFETY_DEPOSIT_SIZE = 1 + 32 + 32 + 32 + 1
const MAX_VAULT_SIZE = 1 + 32 + 32 + 32 + 32 + 1 + 32 + 1 + 32 + 1 + 1 + 8
const MAX_EXTERNAL_ACCOUNT_SIZE = 1 + 8 + 32 + 1
