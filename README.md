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

- README: https://github.com/metaplex-foundation/metaplex/tree/master/rust/auction/program
- Example usage (rust): https://github.com/metaplex-foundation/metaplex/blob/master/rust/auction/cli/src/main.rs
- Source: https://github.com/metaplex-foundation/metaplex/tree/master/rust/auction/program/src


## token-metadata ([Go client](/clients/token-metadata)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/token-metadata))

- README: https://github.com/metaplex-foundation/metaplex/blob/master/rust/token-metadata/program/README.md
- Example usage (rust): https://github.com/metaplex-foundation/metaplex/blob/master/rust/token-metadata/test/src/main.rs
- Source: https://github.com/metaplex-foundation/metaplex/tree/master/rust/token-metadata/program/src

## token-vault ([Go client](/clients/token-vault)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/token-vault))

- README: https://github.com/metaplex-foundation/metaplex/blob/master/rust/token-vault/program/README.md
- Example usage (rust): https://github.com/metaplex-foundation/metaplex/blob/master/rust/token-vault/test/src/main.rs
- Source: https://github.com/metaplex-foundation/metaplex/tree/master/rust/token-vault/program/src


## metaplex ([Go client](/clients/metaplex)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/metaplex))

- README: https://github.com/metaplex-foundation/metaplex/blob/master/rust/metaplex/program/README.md
- Example usage (rust): https://github.com/metaplex-foundation/metaplex/tree/master/rust/metaplex/test/src
- Source: https://github.com/metaplex-foundation/metaplex/tree/master/rust/metaplex/program/src

## nft-candy-machine ([Go client](/clients/nft-candy-machine)) ([Go docs](https://pkg.go.dev/github.com/gagliardetto/metaplex-go/clients/nft-candy-machine))

- Source: https://github.com/metaplex-foundation/metaplex/tree/master/rust/nft-candy-machine/src
- Example usage (typescript): https://github.com/metaplex-foundation/metaplex/blob/master/rust/test/nft-candy-machine.ts
- Usage (go): [draft](/examples/candy)

Clients are build around this version of metaplex programs: https://github.com/metaplex-foundation/metaplex/tree/e9841d4bb121fbea784ff60c83ddd3bb1a26d220
