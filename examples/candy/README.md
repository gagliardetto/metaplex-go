# demo on localnet

- Install dependencies: https://project-serum.github.io/anchor/getting-started/installation.html

- Airdrop yourself some sol.

- In a dedicated terminal tab, start a solana test node:

```bash
solana-test-validator
```

- Build and deploy the programs:

```bash
// Clone metaplex repo:
git clone https://github.com/metaplex-foundation/metaplex.git
cd metaplex/rust/token-metadata

// Build and deploy the `token-metadata` program:
cargo build-bpf program/src
solana program deploy /home/laptop/go/src/github.com/metaplex-foundation/metaplex/rust/target/deploy/spl_token_metadata.so
// and then save the `Program ID` that gets printed after the successful deployment.

// Build and deploy the `nft-candy-machine` program to the test node: 
cd metaplex/rust
anchor deploy
// and then save the `Program Id` that gets printed after the successful deployment.

```

# What to do if I get an error:

- Go to https://github.com/project-serum/anchor/blob/84a2b8200cc3c7cb51d7127918e6cbbd836f0e99/ts/src/error.ts#L48 and look for the error code number.
