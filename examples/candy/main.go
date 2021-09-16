package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
	bin "github.com/gagliardetto/binary"
	nftCandyMachine "github.com/gagliardetto/metaplex-go/clients/nft-candy-machine"
	"github.com/gagliardetto/solana-go"
	atok "github.com/gagliardetto/solana-go/programs/associated-token-account"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
	sendAndConfirmTransaction "github.com/gagliardetto/solana-go/rpc/sendAndConfirmTransaction"
	"github.com/gagliardetto/solana-go/rpc/ws"
	"github.com/gagliardetto/solana-go/text"
	. "github.com/gagliardetto/utilz"
)

var myWallet solana.PrivateKey

func init() {
	var err error
	myWallet, err = solana.PrivateKeyFromSolanaKeygenFile("/path/to/.config/solana/id.json")
	if err != nil {
		panic(err)
	}
}

var (
	candyMachineProgram solana.PublicKey

	// The pubkey of the `token metadata program`;
	// NOTE: the token-metadata program is not part of the standard
	// programs, which means that if you're testing outside of mainnet-beta,
	// YOU MUST DEPLOY THE TOKEN-METADATA PROGRAM YOURSELF (and then specify the ID here).
	//
	// On mainnet-beta, the token-metadata program is metaqbxxUerdq28cj1RbAWkYQm3ybzjb6a8bt518x1s
	TokenMetadataProgramID solana.PublicKey

// var TokenMetadataProgramID = solana.TokenMetadataProgramID
)

type Cache struct {
	WalletPath string
}

func main() {
	type NetEndpoints struct {
		RPC string
		WS  string
	}
	netMap := map[string]NetEndpoints{
		"mainnet": {
			RPC: rpc.MainNetBeta_RPC,
			WS:  rpc.MainNetBeta_WS,
		},
		"devnet": {
			RPC: rpc.DevNet_RPC,
			WS:  rpc.DevNet_WS,
		},
		"testnet": {
			RPC: rpc.TestNet_RPC,
			WS:  rpc.TestNet_WS,
		},
		"localnet": {
			RPC: rpc.LocalNet_RPC,
			WS:  rpc.LocalNet_WS,
		},
	}

	{
		var candyMachineProgramIDString string
		flag.StringVar(&candyMachineProgramIDString, "candy-machine", "", "Candy machine program ID")

		var tokenMetadataProgramIDString string
		flag.StringVar(&tokenMetadataProgramIDString, "token-metadata", "", "Token metadata program ID")

		var netName string
		flag.StringVar(&netName, "net", "localnet", "Solana Network")

		flag.Parse()

		if candyMachineProgramIDString == "" {
			panic("--candy-machine param not provided")
		}
		if tokenMetadataProgramIDString == "" {
			panic("--token-metadata param not provided")
		}

		candyMachineProgram = solana.MustPublicKeyFromBase58(candyMachineProgramIDString)
		TokenMetadataProgramID = solana.MustPublicKeyFromBase58(tokenMetadataProgramIDString)

		if net, ok := netMap[netName]; ok {
			Ln("selected net:", net)
		} else {
			Ln("net not recognized:", net)
		}
		if netName == "mainnet" {
			TokenMetadataProgramID = solana.TokenMetadataProgramID
		} else {
			if TokenMetadataProgramID.IsZero() {
				// TODO
			}
		}
	}

	// Set the programID in the client package;
	// this will allow the istruction parsers to associate the
	// program ID with the dedicated parser.
	nftCandyMachine.SetProgramID(candyMachineProgram)

	// Create a new RPC client:
	rpcClient := rpc.New(rpc.LocalNet_RPC)

	// Create a new WS client (used for confirming transactions)
	wsClient, err := ws.Connect(context.Background(), rpc.LocalNet_WS)
	if err != nil {
		panic(err)
	}

	this := new(configParams)
	config := solana.NewWallet()
	this.Config = config

	// {
	// 	confAccount, err := getConfigAccount(rpcClient, this.Config.PublicKey())
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	spew.Dump(confAccount)
	// }

	{
		txInstr, err := createConfig(this, false, 10)
		if err != nil {
			panic(err)
		}

		linesInstr, err := addConfigLines(this, 10)
		if err != nil {
			panic(err)
		}

		// this.CandyMachineUuid = string(solana.NewWallet().PublicKey().String())[:6]
		this.CandyMachineUuid = "Ew8Q2K"

		candyMachine, bump, err := getCandyMachine(this.Config.PublicKey(), this.CandyMachineUuid)
		if err != nil {
			panic(err)
		}

		// {
		// 	candyMachineAccount, err := getCandyMachineAccount(rpcClient, candyMachine)
		// 	if err != nil {
		// 		panic(err)
		// 	}
		// 	spew.Dump(candyMachineAccount)
		// }

		initInstr := nftCandyMachine.NewInitializeCandyMachineInstructionBuilder().
			SetBump(bump).
			SetData(nftCandyMachine.CandyMachineData{
				Uuid:           this.CandyMachineUuid,
				Price:          1000000000,
				ItemsAvailable: 10,
			}).
			SetCandyMachineAccount(candyMachine).
			SetWalletAccount(myWallet.PublicKey()).
			SetConfigAccount(this.Config.PublicKey()).
			SetAuthorityAccount(this.Authority.PublicKey()).
			SetPayerAccount(myWallet.PublicKey()).
			SetSystemProgramAccount(solana.SystemProgramID).
			SetRentAccount(solana.SysVarRentPubkey)

		signers := make([]solana.PrivateKey, 0)
		signers = append(signers,
			myWallet,
			this.Authority.PrivateKey,
			this.Config.PrivateKey,
		)
		if true {
			spew.Dump(this)
			space := uint64(configArrayStart + 4 + 10*configLineSize + 4 + 2)
			instructions := []solana.Instruction{
				system.NewCreateAccountInstructionBuilder().
					SetLamports(solana.LAMPORTS_PER_SOL * 10).
					SetSpace(space).
					SetOwner(candyMachineProgram).
					SetFundingAccount(myWallet.PublicKey()).
					SetNewAccount(config.PublicKey()).
					Build(),

				system.NewTransferInstructionBuilder().
					SetLamports(500000).
					SetFundingAccount(myWallet.PublicKey()).
					SetRecipientAccount(this.Authority.PublicKey()).
					Build(),

				txInstr,
			}
			instructions = append(instructions, linesInstr...)
			sendTx(
				rpcClient,
				wsClient,
				"create candy machine account",
				instructions,
				signers,
			)
			sendTx(
				rpcClient,
				wsClient,
				"init candy machine",
				[]solana.Instruction{initInstr.Build()},
				signers,
			)
		}
		{
			// Mint:

			// mint, err := solana.PrivateKeyFromSolanaKeygenFile("/path/to/candy-mint-1.json")
			// if err != nil {
			// 	panic(err)
			// }
			mint := solana.NewWallet()
			signers = append(signers, mint.PrivateKey)
			tok := getTokenWallet(
				this.Authority.PublicKey(),
				mint.PublicKey(),
			)
			metadata := getMetadata(mint.PublicKey())
			masterEdition := getMasterEdition(mint.PublicKey())

			mintInstruction := nftCandyMachine.NewMintNftInstructionBuilder().
				SetConfigAccount(this.Config.PublicKey()).
				SetCandyMachineAccount(candyMachine).
				SetPayerAccount(this.Authority.PublicKey()).
				SetWalletAccount(myWallet.PublicKey()).
				SetMetadataAccount(metadata).
				SetMintAccount(mint.PublicKey()).
				SetMintAuthorityAccount(this.Authority.PublicKey()).
				SetUpdateAuthorityAccount(this.Authority.PublicKey()).
				SetMasterEditionAccount(masterEdition).
				SetTokenMetadataProgramAccount(TokenMetadataProgramID).
				SetTokenProgramAccount(solana.TokenProgramID).
				SetSystemProgramAccount(solana.SystemProgramID).
				SetRentAccount(solana.SysVarRentPubkey).
				SetClockAccount(solana.SysVarClockPubkey)

			_ = mintInstruction

			instructions := []solana.Instruction{
				system.NewTransferInstructionBuilder().
					SetLamports(1000000000 + 10000000). // add minting fees in there
					SetFundingAccount(myWallet.PublicKey()).
					SetRecipientAccount(this.Authority.PublicKey()).
					Build(),

				func() solana.Instruction {
					in, err := system.NewCreateAccountInstructionBuilder().
						SetLamports(solana.LAMPORTS_PER_SOL * 100).
						SetSpace(token.MINT_SIZE).
						SetOwner(solana.TokenProgramID).
						SetFundingAccount(myWallet.PublicKey()).
						SetNewAccount(mint.PublicKey()).
						ValidateAndBuild()
					if err != nil {
						panic(err)
					}
					return in
				}(),

				token.NewInitializeMintInstructionBuilder().
					SetDecimals(0).
					SetMintAuthority(this.Authority.PublicKey()).
					SetFreezeAuthority(this.Authority.PublicKey()).
					SetMintAccount(mint.PublicKey()).
					SetSysVarRentPubkeyAccount(solana.SysVarRentPubkey).
					Build(),

				atok.NewCreateInstructionBuilder().
					SetPayer(myWallet.PublicKey()).
					SetWallet(this.Authority.PublicKey()).
					SetMint(mint.PublicKey()).Build(),
				// createAssociatedTokenAccountInstruction(
				// 	tok,
				// 	myWallet.PublicKey(),
				// 	this.Authority.PublicKey(),
				// 	mint.PublicKey(),
				// ),

				token.NewMintToInstructionBuilder().
					SetAmount(1).
					SetMintAccount(mint.PublicKey()).
					SetDestinationAccount(tok).
					SetAuthorityAccount(this.Authority.PublicKey()).
					Build(),

				// mintInstruction.Build(),
			}
			_ = instructions

			sendTx(
				rpcClient,
				wsClient,
				"init mint",
				instructions,
				signers,
			)
			Ln("ok----")
			if true {
				resp, err := rpcClient.GetAccountInfo(context.Background(), mint.PublicKey())
				if err != nil {
					panic(err)
				}
				candyBoxBytes := resp.Value.Data.GetBinary()

				Ln(resp.Value.Data)
				Ln(FormatByteSlice(candyBoxBytes[:]))
			}
			sendTx(
				rpcClient,
				wsClient,
				"do mint",
				[]solana.Instruction{mintInstruction.Build()},
				signers,
			)
		}
	}
}

// ConfigAndLines holds the contents of a Config account.
type ConfigAndLines struct {
	Config nftCandyMachine.Config
	Lines  []nftCandyMachine.ConfigLine
}

// where the lines start in the config account data
const configArrayStart = 32 + // authority
	4 +
	6 + // uuid + u32 len
	4 +
	10 + // u32 len + symbol
	2 + // seller fee basis points
	1 +
	4 +
	5*34 + // optional + u32 len + actual vec
	8 + //max supply
	1 + //is mutable
	1 + // retain authority
	4 // max number of lines;

const configLineSize = 4 + 32 + 4 + 200

func (acc *ConfigAndLines) UnmarshalWithDecoder(decoder *bin.Decoder) (err error) {
	total := decoder.Remaining()

	// Decode Config:
	err = decoder.Decode(&acc.Config)
	if err != nil {
		return err
	}

	remaining := decoder.Remaining()
	read := total - remaining
	// Fast formard to the lines array:
	err = decoder.SkipBytes(uint(configArrayStart - read))
	if err != nil {
		return err
	}

	acc.Lines = make([]nftCandyMachine.ConfigLine, 0)
	// Decode config lines:
	return decoder.Decode(&acc.Lines)
}

// getConfigAccount loads and parses the contents of a `Config` account.
func getConfigAccount(rpcClient *rpc.Client, configAccountPubkey solana.PublicKey) (*ConfigAndLines, error) {
	resp, err := rpcClient.GetAccountInfo(context.Background(), configAccountPubkey)
	if err != nil {
		return nil, err
	}

	configAccountBytes := resp.Value.Data.GetBinary()
	dec := bin.NewBorshDecoder(configAccountBytes)

	var acc ConfigAndLines

	err = dec.Decode(&acc)
	if err != nil {
		return nil, err
	}
	return &acc, nil
}

// getCandyMachineAccount loads and parses the contents of a `CandyMachine` account.
func getCandyMachineAccount(rpcClient *rpc.Client, candyMachineAccountPubkey solana.PublicKey) (*nftCandyMachine.CandyMachine, error) {
	resp, err := rpcClient.GetAccountInfo(context.Background(), candyMachineAccountPubkey)
	if err != nil {
		return nil, err
	}
	candyBoxBytes := resp.Value.Data.GetBinary()

	dec := bin.NewBorshDecoder(candyBoxBytes)
	var config nftCandyMachine.CandyMachine
	err = dec.Decode(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func mustGetMinimumBalanceForRentExemption(
	rpcClient *rpc.Client,
	dataSize uint64,
) uint64 {
	minBalance, err := rpcClient.GetMinimumBalanceForRentExemption(
		context.Background(),
		dataSize,
		rpc.CommitmentFinalized,
	)
	if err != nil {
		panic(err)
	}
	return minBalance
}

const CANDY_MACHINE = "candy_machine"

func sendTx(
	rpcClient *rpc.Client,
	wsClient *ws.Client,
	doc string,
	instructions []solana.Instruction,
	signers []solana.PrivateKey,
) {
	recent, err := rpcClient.GetRecentBlockhash(context.Background(), rpc.CommitmentFinalized)
	if err != nil {
		panic(err)
	}

	tx, err := solana.NewTransaction(
		instructions,
		recent.Value.Blockhash,
		solana.TransactionPayer(myWallet.PublicKey()),
	)
	if err != nil {
		panic(err)
	}

	_, err = tx.Sign(func(key solana.PublicKey) *solana.PrivateKey {
		for _, candidate := range signers {
			if candidate.PublicKey().Equals(key) {
				return &candidate
			}
		}
		return nil
	})
	if err != nil {
		panic(fmt.Errorf("unable to sign transaction: %w", err))
	}
	tx.EncodeTree(text.NewTreeEncoder(os.Stdout, doc))
	{
		out, err := tx.MarshalBinary()
		if err != nil {
			panic(err)
		}
		Ln(FormatByteSlice(out))
	}

	sig, err := sendAndConfirmTransaction.SendAndConfirmTransaction(
		context.Background(),
		rpcClient,
		wsClient,
		tx,
	)
	if err != nil {
		panic(err)
	}
	spew.Dump(sig)
}

func getTokenWallet(
	wallet solana.PublicKey,
	mint solana.PublicKey,
) solana.PublicKey {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			wallet.Bytes(),
			solana.TokenProgramID.Bytes(),
			mint.Bytes(),
		},
		solana.SPLAssociatedTokenAccountProgramID,
	)
	if err != nil {
		panic(err)
	}
	return addr
}

func getMetadata(
	mint solana.PublicKey,
) solana.PublicKey {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			TokenMetadataProgramID.Bytes(),
			mint.Bytes(),
		},
		TokenMetadataProgramID,
	)
	if err != nil {
		panic(err)
	}
	return addr
}

func getMasterEdition(
	mint solana.PublicKey,
) solana.PublicKey {
	addr, _, err := solana.FindProgramAddress(
		[][]byte{
			[]byte("metadata"),
			TokenMetadataProgramID.Bytes(),
			mint.Bytes(),
			[]byte("edition"),
		},
		TokenMetadataProgramID,
	)
	if err != nil {
		panic(err)
	}
	return addr
}

func getCandyMachine(
	config solana.PublicKey,
	uuid string,
) (solana.PublicKey, uint8, error) {
	return solana.FindProgramAddress(
		[][]byte{
			[]byte(CANDY_MACHINE),
			config.Bytes(),
			[]byte(uuid),
		},
		candyMachineProgram,
	)
}

type configParams struct {
	UUID      string
	Config    *solana.Wallet
	Authority *solana.Wallet

	CandyMachineUuid string
}

func createConfig(
	that *configParams,
	retainAuthority bool,
	size uint32,
) (solana.Instruction, error) {
	var err error
	that.Authority, err = solana.WalletFromPrivateKeyBase58("4prL2LjqVLevvuoFVHNh3UpQN3BRFSe6pxRUhZfXHzZCbmsnJsUdrHTctq4Mcyty49RxKuXpWrMk7iBxVg78fCxw")
	if err != nil {
		panic(err)
	}
	// that.UUID = string(solana.NewWallet().PublicKey().String())[:6]
	that.UUID = "J5S1AQ"

	instruction := nftCandyMachine.NewInitializeConfigInstruction(
		nftCandyMachine.ConfigData{
			Uuid:                 that.UUID,
			MaxNumberOfLines:     size,
			Symbol:               "SYMBOL",
			SellerFeeBasisPoints: 500,
			IsMutable:            true,
			MaxSupply:            0,
			RetainAuthority:      retainAuthority,
			Creators: []nftCandyMachine.Creator{
				{Address: myWallet.PublicKey(), Verified: false, Share: 100},
			},
		},
		that.Config.PublicKey(),
		that.Authority.PublicKey(),
		myWallet.PublicKey(),
		solana.SysVarRentPubkey,
	)
	if err := instruction.Validate(); err != nil {
		return nil, err
	}

	return instruction.Build(), nil
}

func addConfigLines(
	that *configParams,
	size uint32,
) ([]solana.Instruction, error) {
	ins := make([]solana.Instruction, 0)
	{
		firstVec := []nftCandyMachine.ConfigLine{}

		for i := 0; i < 5; i++ {
			firstVec = append(firstVec, nftCandyMachine.ConfigLine{Uri: "www.aol.com", Name: Sf(`Sample %v`, i)})
		}

		instruction := nftCandyMachine.NewAddConfigLinesInstruction(
			0,
			firstVec,
			that.Config.PublicKey(),
			that.Authority.PublicKey(),
		)
		if err := instruction.Validate(); err != nil {
			return nil, err
		}
		ins = append(ins, instruction.Build())
	}
	{
		firstVec := []nftCandyMachine.ConfigLine{}

		for i := 5; i < 10; i++ {
			firstVec = append(firstVec, nftCandyMachine.ConfigLine{Uri: "www.aol.com", Name: Sf(`Sample %v`, i)})
		}

		instruction := nftCandyMachine.NewAddConfigLinesInstruction(
			5,
			firstVec,
			that.Config.PublicKey(),
			that.Authority.PublicKey(),
		)
		if err := instruction.Validate(); err != nil {
			return nil, err
		}
		ins = append(ins, instruction.Build())
		ins = append(ins, instruction.Build())
	}

	return ins, nil
}

func TransactionInstruction(
	keys solana.AccountMetaSlice,
	programID solana.PublicKey,
	data []byte,
) solana.Instruction {
	return &Instr{
		keys:      keys,
		programID: programID,
		data:      data,
	}
}

type Instr struct {
	keys      solana.AccountMetaSlice
	programID solana.PublicKey
	data      []byte
}

func (in *Instr) ProgramID() solana.PublicKey {
	return in.programID
}

func (in *Instr) Accounts() []*solana.AccountMeta {
	return in.keys
}

func (in *Instr) Data() ([]byte, error) {
	return in.data, nil
}

func createAssociatedTokenAccountInstruction(
	associatedTokenAddress solana.PublicKey,
	payer solana.PublicKey,
	walletAddress solana.PublicKey,
	splTokenMintAddress solana.PublicKey,
) solana.Instruction {

	keys := []*solana.AccountMeta{
		{
			PublicKey:  payer,
			IsSigner:   true,
			IsWritable: true,
		},
		{
			PublicKey:  associatedTokenAddress,
			IsSigner:   false,
			IsWritable: true,
		},
		{
			PublicKey:  walletAddress,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  splTokenMintAddress,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  solana.SystemProgramID,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  solana.TokenProgramID,
			IsSigner:   false,
			IsWritable: false,
		},
		{
			PublicKey:  solana.SysVarRentPubkey,
			IsSigner:   false,
			IsWritable: false,
		},
	}

	return TransactionInstruction(
		keys,
		solana.SPLAssociatedTokenAccountProgramID,
		[]byte{},
	)
}
