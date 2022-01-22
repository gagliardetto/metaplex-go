# metaplex-go

A suite of Go clients for the 5 metaplex contracts.

This is an `alpha` version.

For usage examples, you can get inspired by their Rust/Typescript counterparts (noted in the below list).

## Requirements

`metaplex-go` requires `go1.16` or later.

## Installation

```bash
cd my-project
go get github.com/gagliardetto/metaplex-go
```

If you encounter `missing go.sum entry for module providing package ...` errors, it means that you need to run `go mod tidy`.

---

## auction ([Go client](/clients/auction)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/auction))

- README: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/auction/program
- Example usage (typescript): https://github.com/metaplex-foundation/metaplex/blob/89716fbe6821e5155971bb085aac39472a13774f/js/packages/common/src/actions/auction.ts
- Source: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/auction/program/src

## token-metadata ([Go client](/clients/token-metadata)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/token-metadata))

- README: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/token-metadata/program/README.md
- Example usage (rust): https://github.com/metaplex-foundation/metaplex-program-library/tree/master/token-metadata/program/tests
- Source: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/token-metadata/program/src

## token-vault ([Go client](/clients/token-vault)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/token-vault))

- README: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/token-vault/program/README.md
- Example usage (rust): https://github.com/metaplex-foundation/metaplex-program-library/blob/master/token-vault/test/src/main.rs
- Source: https://github.com/metaplex-foundation/metaplex-program-library/blob/master/token-vault/program/src

## metaplex ([Go client](/clients/metaplex)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/metaplex))

- README: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/metaplex/program/README.md
- Example usage (rust): https://github.com/metaplex-foundation/metaplex-program-library/blob/master/metaplex/program/tests/
- Source: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/metaplex/program/src

## nft-candy-machine (V1) is DEPRECATED; use V2 instead.

## nft-candy-machine V2 ([Go client](/clients/nft-candy-machine-v2)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/nft-candy-machine-v2))

- Source: https://github.com/metaplex-foundation/metaplex-program-library/tree/master/nft-candy-machine/program
- Example usage (typescript): https://github.com/metaplex-foundation/metaplex/tree/master/js/packages/candy-machine-ui/src
- Usage (go): [non-working draft](/examples/candy-v2)

Clients are build around this version of metaplex programs: https://github.com/metaplex-foundation/metaplex-program-library/tree/821e5aac0780fe45525dae72b9ad6f8dc2ba8e5b

TODO:

## auction-house
## gumdrop
## membership-token
## nft-packs
## token-entagler
