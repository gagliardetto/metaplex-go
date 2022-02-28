package main

import (
	"context"
	"fmt"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	nftcandymachinev2 "github.com/gagliardetto/metaplex-go/clients/nft-candy-machine-v2"
	token_metadata "github.com/gagliardetto/metaplex-go/clients/token-metadata"
	"github.com/gagliardetto/solana-go"
	atok "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	sendAndConfirmTransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
)

var candyMachineV2ProgramID = solana.MustPublicKeyFromBase58("cndyAnrLdpjq1Ssp1z8xxDsB8dxe7u4HL5Nxi2K5WXZ")

func main() {
	user, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
	if err != nil {
		panic(err)
	}

	candyMachineAddress := solana.MustPublicKeyFromBase58("TODO")

	sig, err := mint(
		candyMachineAddress,
		user,
		rpc.MainNetBeta,
	)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Success: %s", sig)
}

func mint(
	candyMachineAddress solana.PublicKey,
	userKeyPair solana.PrivateKey,
	cluster rpc.Cluster,
) (solana.Signature, error) {
	nftcandymachinev2.SetProgramID(candyMachineV2ProgramID)

	mint := solana.NewWallet()

	client := rpc.New(cluster.RPC)
	userTokenAccountAddress, err := getTokenWallet(userKeyPair.PublicKey(), mint.PublicKey())
	if err != nil {
		return solana.Signature{}, err
	}

	candyMachineRaw, err := client.GetAccountInfo(context.TODO(), candyMachineAddress)
	if err != nil {
		return solana.Signature{}, err
	}

	signers := []solana.PrivateKey{mint.PrivateKey, userKeyPair}

	min, err := client.GetMinimumBalanceForRentExemption(context.TODO(), uint64(1024*9), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, err
	}

	dec := bin.NewBorshDecoder(candyMachineRaw.Value.Data.GetBinary())
	var cm nftcandymachinev2.CandyMachine
	err = dec.Decode(&cm)
	if err != nil {
		return solana.Signature{}, err
	}

	var instructions []solana.Instruction
	instructions = append(instructions,
		system.NewCreateAccountInstructionBuilder().
			SetOwner(token.ProgramID).
			SetNewAccount(mint.PublicKey()).
			SetSpace(candyMachineRaw.Value.Lamports).
			SetLamports(min).
			Build(),

		token.NewInitializeMint2InstructionBuilder().
			SetMintAccount(mint.PublicKey()).
			SetDecimals(0).
			SetMintAuthority(userKeyPair.PublicKey()).
			SetFreezeAuthority(userKeyPair.PublicKey()).
			Build(),

		atok.NewCreateInstructionBuilder().
			SetPayer(userKeyPair.PublicKey()).
			SetWallet(userKeyPair.PublicKey()).
			SetMint(mint.PublicKey()).
			Build(),

		token.NewMintToInstructionBuilder().
			SetMintAccount(mint.PublicKey()).
			SetDestinationAccount(userTokenAccountAddress).
			SetAuthorityAccount(userKeyPair.PublicKey()).
			SetAmount(1).
			Build(),
	)

	metadataAddress, err := getMetadata(mint.PublicKey())
	if err != nil {
		return solana.Signature{}, err
	}
	masterEdition, err := getMasterEdition(mint.PublicKey())
	if err != nil {
		return solana.Signature{}, err
	}
	candyMachineCreator, creatorBump, err := getCandyMachineCreator(candyMachineAddress)
	if err != nil {
		return solana.Signature{}, err
	}

	instructions = append(instructions,
		nftcandymachinev2.NewMintNftInstructionBuilder().
			SetCreatorBump(creatorBump).
			SetCandyMachineAccount(candyMachineAddress).
			SetCandyMachineCreatorAccount(candyMachineCreator).
			SetPayerAccount(userKeyPair.PublicKey()).
			SetWalletAccount(cm.Wallet).
			SetMintAccount(mint.PublicKey()).
			SetMetadataAccount(metadataAddress).
			SetMasterEditionAccount(masterEdition).
			SetMintAuthorityAccount(userKeyPair.PublicKey()).
			SetUpdateAuthorityAccount(userKeyPair.PublicKey()).
			SetTokenMetadataProgramAccount(token_metadata.ProgramID).
			SetTokenProgramAccount(token.ProgramID).
			SetSystemProgramAccount(system.ProgramID).
			SetRentAccount(solana.SysVarRentPubkey).
			SetClockAccount(solana.SysVarClockPubkey).
			SetRecentBlockhashesAccount(solana.SysVarRecentBlockHashesPubkey).
			SetInstructionSysvarAccountAccount(solana.SysVarInstructionsPubkey).
			Build(),
	)

	return sendTransaction(
		cluster,
		client,
		userKeyPair,
		instructions,
		signers,
	)
}

func sendTransaction(
	cluster rpc.Cluster,
	client *rpc.Client,
	wallet solana.PrivateKey,
	instructions []solana.Instruction,
	signers []solana.PrivateKey,
) (solana.Signature, error) {

	recent, err := client.GetRecentBlockhash(context.TODO(), rpc.CommitmentFinalized)
	if err != nil {
		return solana.Signature{}, err
	}

	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(wallet.PublicKey()),
	)

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}
		return nil
	})
	if err != nil {
		return solana.Signature{}, err
	}
	spew.Dump(tx)

	wsClient, err := ws.Connect(context.Background(), cluster.WS)
	if err != nil {
		return solana.Signature{}, err
	}

	return sendAndConfirmTransaction.SendAndConfirmTransaction(
		context.TODO(),
		client,
		wsClient,
		tx,
	)
}

func getTokenWallet(wallet solana.PublicKey, mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			wallet.Bytes(),
			solana.TokenProgramID.Bytes(),
			mint.Bytes(),
		},
		solana.SPLAssociatedTokenAccountProgramID,
	)
	return addr, err
}

func getCandyMachineCreator(candyMachineAddress solana.PublicKey) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress(
		[][]byte{
			[]byte("candy_machine"),
			candyMachineAddress.Bytes(),
		},
		candyMachineV2ProgramID,
	)
}

func getMetadata(mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			token_metadata.ProgramID.Bytes(),
			mint.Bytes(),
		},
		token_metadata.ProgramID,
	)
	return addr, err
}

func getMasterEdition(mint solana.PublicKey) (solana.PublicKey, error) {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			token_metadata.ProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		token_metadata.ProgramID,
	)
	return addr, err
}
