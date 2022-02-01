// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Register a Metadata as a Master Edition V1, which means Editions can be minted.
// Henceforth, no further tokens will be mintable from this primary mint. Will throw an error if more than one
// token exists, and will throw an error if less than one token exists in this primary mint.
type DeprecatedCreateMasterEdition struct {
	Args *CreateMasterEditionArgs

	// [0] = [WRITE] unallocatedEditionV1
	// ··········· Unallocated edition V1 account with address as pda of ['metadata', program id, mint, 'edition']
	//
	// [1] = [WRITE] metadataMint
	// ··········· Metadata mint
	//
	// [2] = [WRITE] printingMint
	// ··········· Printing mint - A mint you control that can mint tokens that can be exchanged for limited editions of your
	// ··········· master edition via the MintNewEditionFromMasterEditionViaToken endpoint
	//
	// [3] = [WRITE] oneTimeAuthPrintingMint
	// ··········· One time authorization printing mint - A mint you control that prints tokens that gives the bearer permission to mint any
	// ··········· number of tokens from the printing mint one time via an endpoint with the token-metadata program for your metadata. Also burns the token.
	//
	// [4] = [SIGNER] currentUpdateAuthority
	// ··········· Current Update authority key
	//
	// [5] = [SIGNER] printingMintAuthority
	// ··········· Printing mint authority - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
	//
	// [6] = [SIGNER] mintAuthorityOnMetadataMint
	// ··········· Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
	//
	// [7] = [] metadata
	// ··········· Metadata account
	//
	// [8] = [SIGNER] payer
	// ··········· payer
	//
	// [9] = [] tokenProgram
	// ··········· Token program
	//
	// [10] = [] system
	// ··········· System program
	//
	// [11] = [] rent
	// ··········· Rent info
	//
	// [12] = [SIGNER] oneTimeAuthorizationPrintingMintAuthority
	// ··········· One time authorization printing mint authority - must be provided if using max supply. THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
	ag_solanago.AccountMetaSlice `bin:"-"`
}

// NewDeprecatedCreateMasterEditionInstructionBuilder creates a new `DeprecatedCreateMasterEdition` instruction builder.
func NewDeprecatedCreateMasterEditionInstructionBuilder() *DeprecatedCreateMasterEdition {
	nd := &DeprecatedCreateMasterEdition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 13),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *DeprecatedCreateMasterEdition) SetArgs(args CreateMasterEditionArgs) *DeprecatedCreateMasterEdition {
	inst.Args = &args
	return inst
}

// SetUnallocatedEditionV1Account sets the "unallocatedEditionV1" account.
// Unallocated edition V1 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *DeprecatedCreateMasterEdition) SetUnallocatedEditionV1Account(unallocatedEditionV1 ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(unallocatedEditionV1).WRITE()
	return inst
}

// GetUnallocatedEditionV1Account gets the "unallocatedEditionV1" account.
// Unallocated edition V1 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *DeprecatedCreateMasterEdition) GetUnallocatedEditionV1Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(0)
}

// SetMetadataMintAccount sets the "metadataMint" account.
// Metadata mint
func (inst *DeprecatedCreateMasterEdition) SetMetadataMintAccount(metadataMint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(metadataMint).WRITE()
	return inst
}

// GetMetadataMintAccount gets the "metadataMint" account.
// Metadata mint
func (inst *DeprecatedCreateMasterEdition) GetMetadataMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(1)
}

// SetPrintingMintAccount sets the "printingMint" account.
// Printing mint - A mint you control that can mint tokens that can be exchanged for limited editions of your
// master edition via the MintNewEditionFromMasterEditionViaToken endpoint
func (inst *DeprecatedCreateMasterEdition) SetPrintingMintAccount(printingMint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(printingMint).WRITE()
	return inst
}

// GetPrintingMintAccount gets the "printingMint" account.
// Printing mint - A mint you control that can mint tokens that can be exchanged for limited editions of your
// master edition via the MintNewEditionFromMasterEditionViaToken endpoint
func (inst *DeprecatedCreateMasterEdition) GetPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(2)
}

// SetOneTimeAuthPrintingMintAccount sets the "oneTimeAuthPrintingMint" account.
// One time authorization printing mint - A mint you control that prints tokens that gives the bearer permission to mint any
// number of tokens from the printing mint one time via an endpoint with the token-metadata program for your metadata. Also burns the token.
func (inst *DeprecatedCreateMasterEdition) SetOneTimeAuthPrintingMintAccount(oneTimeAuthPrintingMint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(oneTimeAuthPrintingMint).WRITE()
	return inst
}

// GetOneTimeAuthPrintingMintAccount gets the "oneTimeAuthPrintingMint" account.
// One time authorization printing mint - A mint you control that prints tokens that gives the bearer permission to mint any
// number of tokens from the printing mint one time via an endpoint with the token-metadata program for your metadata. Also burns the token.
func (inst *DeprecatedCreateMasterEdition) GetOneTimeAuthPrintingMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(3)
}

// SetCurrentUpdateAuthorityAccount sets the "currentUpdateAuthority" account.
// Current Update authority key
func (inst *DeprecatedCreateMasterEdition) SetCurrentUpdateAuthorityAccount(currentUpdateAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(currentUpdateAuthority).SIGNER()
	return inst
}

// GetCurrentUpdateAuthorityAccount gets the "currentUpdateAuthority" account.
// Current Update authority key
func (inst *DeprecatedCreateMasterEdition) GetCurrentUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(4)
}

// SetPrintingMintAuthorityAccount sets the "printingMintAuthority" account.
// Printing mint authority - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) SetPrintingMintAuthorityAccount(printingMintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(printingMintAuthority).SIGNER()
	return inst
}

// GetPrintingMintAuthorityAccount gets the "printingMintAuthority" account.
// Printing mint authority - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) GetPrintingMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(5)
}

// SetMintAuthorityOnMetadataMintAccount sets the "mintAuthorityOnMetadataMint" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *DeprecatedCreateMasterEdition) SetMintAuthorityOnMetadataMintAccount(mintAuthorityOnMetadataMint ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(mintAuthorityOnMetadataMint).SIGNER()
	return inst
}

// GetMintAuthorityOnMetadataMintAccount gets the "mintAuthorityOnMetadataMint" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *DeprecatedCreateMasterEdition) GetMintAuthorityOnMetadataMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(6)
}

// SetMetadataAccount sets the "metadata" account.
// Metadata account
func (inst *DeprecatedCreateMasterEdition) SetMetadataAccount(metadata ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(metadata)
	return inst
}

// GetMetadataAccount gets the "metadata" account.
// Metadata account
func (inst *DeprecatedCreateMasterEdition) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(7)
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *DeprecatedCreateMasterEdition) SetPayerAccount(payer ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *DeprecatedCreateMasterEdition) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(8)
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *DeprecatedCreateMasterEdition) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[9] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *DeprecatedCreateMasterEdition) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(9)
}

// SetSystemAccount sets the "system" account.
// System program
func (inst *DeprecatedCreateMasterEdition) SetSystemAccount(system ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[10] = ag_solanago.Meta(system)
	return inst
}

// GetSystemAccount gets the "system" account.
// System program
func (inst *DeprecatedCreateMasterEdition) GetSystemAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(10)
}

// SetRentAccount sets the "rent" account.
// Rent info
func (inst *DeprecatedCreateMasterEdition) SetRentAccount(rent ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[11] = ag_solanago.Meta(rent)
	return inst
}

// GetRentAccount gets the "rent" account.
// Rent info
func (inst *DeprecatedCreateMasterEdition) GetRentAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(11)
}

// SetOneTimeAuthorizationPrintingMintAuthorityAccount sets the "oneTimeAuthorizationPrintingMintAuthority" account.
// One time authorization printing mint authority - must be provided if using max supply. THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) SetOneTimeAuthorizationPrintingMintAuthorityAccount(oneTimeAuthorizationPrintingMintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	inst.AccountMetaSlice[12] = ag_solanago.Meta(oneTimeAuthorizationPrintingMintAuthority).SIGNER()
	return inst
}

// GetOneTimeAuthorizationPrintingMintAuthorityAccount gets the "oneTimeAuthorizationPrintingMintAuthority" account.
// One time authorization printing mint authority - must be provided if using max supply. THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY.
func (inst *DeprecatedCreateMasterEdition) GetOneTimeAuthorizationPrintingMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice.Get(12)
}

func (inst DeprecatedCreateMasterEdition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DeprecatedCreateMasterEdition),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DeprecatedCreateMasterEdition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DeprecatedCreateMasterEdition) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UnallocatedEditionV1 is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MetadataMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.PrintingMint is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.OneTimeAuthPrintingMint is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.CurrentUpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.PrintingMintAuthority is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.MintAuthorityOnMetadataMint is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.Metadata is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[9] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[10] == nil {
			return errors.New("accounts.System is not set")
		}
		if inst.AccountMetaSlice[11] == nil {
			return errors.New("accounts.Rent is not set")
		}
		if inst.AccountMetaSlice[12] == nil {
			return errors.New("accounts.OneTimeAuthorizationPrintingMintAuthority is not set")
		}
	}
	return nil
}

func (inst *DeprecatedCreateMasterEdition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DeprecatedCreateMasterEdition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=13]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("                     unallocatedEditionV1", inst.AccountMetaSlice.Get(0)))
						accountsBranch.Child(ag_format.Meta("                             metadataMint", inst.AccountMetaSlice.Get(1)))
						accountsBranch.Child(ag_format.Meta("                             printingMint", inst.AccountMetaSlice.Get(2)))
						accountsBranch.Child(ag_format.Meta("                  oneTimeAuthPrintingMint", inst.AccountMetaSlice.Get(3)))
						accountsBranch.Child(ag_format.Meta("                   currentUpdateAuthority", inst.AccountMetaSlice.Get(4)))
						accountsBranch.Child(ag_format.Meta("                    printingMintAuthority", inst.AccountMetaSlice.Get(5)))
						accountsBranch.Child(ag_format.Meta("              mintAuthorityOnMetadataMint", inst.AccountMetaSlice.Get(6)))
						accountsBranch.Child(ag_format.Meta("                                 metadata", inst.AccountMetaSlice.Get(7)))
						accountsBranch.Child(ag_format.Meta("                                    payer", inst.AccountMetaSlice.Get(8)))
						accountsBranch.Child(ag_format.Meta("                             tokenProgram", inst.AccountMetaSlice.Get(9)))
						accountsBranch.Child(ag_format.Meta("                                   system", inst.AccountMetaSlice.Get(10)))
						accountsBranch.Child(ag_format.Meta("                                     rent", inst.AccountMetaSlice.Get(11)))
						accountsBranch.Child(ag_format.Meta("oneTimeAuthorizationPrintingMintAuthority", inst.AccountMetaSlice.Get(12)))
					})
				})
		})
}

func (obj DeprecatedCreateMasterEdition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *DeprecatedCreateMasterEdition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewDeprecatedCreateMasterEditionInstruction declares a new DeprecatedCreateMasterEdition instruction with the provided parameters and accounts.
func NewDeprecatedCreateMasterEditionInstruction(
	// Parameters:
	args CreateMasterEditionArgs,
	// Accounts:
	unallocatedEditionV1 ag_solanago.PublicKey,
	metadataMint ag_solanago.PublicKey,
	printingMint ag_solanago.PublicKey,
	oneTimeAuthPrintingMint ag_solanago.PublicKey,
	currentUpdateAuthority ag_solanago.PublicKey,
	printingMintAuthority ag_solanago.PublicKey,
	mintAuthorityOnMetadataMint ag_solanago.PublicKey,
	metadata ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	system ag_solanago.PublicKey,
	rent ag_solanago.PublicKey,
	oneTimeAuthorizationPrintingMintAuthority ag_solanago.PublicKey) *DeprecatedCreateMasterEdition {
	return NewDeprecatedCreateMasterEditionInstructionBuilder().
		SetArgs(args).
		SetUnallocatedEditionV1Account(unallocatedEditionV1).
		SetMetadataMintAccount(metadataMint).
		SetPrintingMintAccount(printingMint).
		SetOneTimeAuthPrintingMintAccount(oneTimeAuthPrintingMint).
		SetCurrentUpdateAuthorityAccount(currentUpdateAuthority).
		SetPrintingMintAuthorityAccount(printingMintAuthority).
		SetMintAuthorityOnMetadataMintAccount(mintAuthorityOnMetadataMint).
		SetMetadataAccount(metadata).
		SetPayerAccount(payer).
		SetTokenProgramAccount(tokenProgram).
		SetSystemAccount(system).
		SetRentAccount(rent).
		SetOneTimeAuthorizationPrintingMintAuthorityAccount(oneTimeAuthorizationPrintingMintAuthority)
}
