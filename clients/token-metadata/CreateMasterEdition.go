// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package token_metadata

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// Register a Metadata as a Master Edition V2, which means Edition V2s can be minted.
// Henceforth, no further tokens will be mintable from this primary mint. Will throw an error if more than one
// token exists, and will throw an error if less than one token exists in this primary mint.
type CreateMasterEdition struct {
	Args *CreateMasterEditionArgs

	// [0] = [WRITE] unallocatedEditionV2
	// ··········· Unallocated edition V2 account with address as pda of ['metadata', program id, mint, 'edition']
	//
	// [1] = [WRITE] metadataMint
	// ··········· Metadata mint
	//
	// [2] = [SIGNER] updateAuthority
	// ··········· Update authority
	//
	// [3] = [SIGNER] mintAuthority
	// ··········· Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
	//
	// [4] = [SIGNER] payer
	// ··········· payer
	//
	// [5] = [] metadataAccount
	// ··········· Metadata account
	//
	// [6] = [] tokenProgram
	// ··········· Token program
	//
	// [7] = [] systemProgram
	// ··········· System program
	//
	// [8] = [] rentInfo
	// ··········· Rent info
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewCreateMasterEditionInstructionBuilder creates a new `CreateMasterEdition` instruction builder.
func NewCreateMasterEditionInstructionBuilder() *CreateMasterEdition {
	nd := &CreateMasterEdition{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 9),
	}
	return nd
}

// SetArgs sets the "args" parameter.
func (inst *CreateMasterEdition) SetArgs(args CreateMasterEditionArgs) *CreateMasterEdition {
	inst.Args = &args
	return inst
}

// SetUnallocatedEditionV2Account sets the "unallocatedEditionV2" account.
// Unallocated edition V2 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *CreateMasterEdition) SetUnallocatedEditionV2Account(unallocatedEditionV2 ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(unallocatedEditionV2).WRITE()
	return inst
}

// GetUnallocatedEditionV2Account gets the "unallocatedEditionV2" account.
// Unallocated edition V2 account with address as pda of ['metadata', program id, mint, 'edition']
func (inst *CreateMasterEdition) GetUnallocatedEditionV2Account() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetMetadataMintAccount sets the "metadataMint" account.
// Metadata mint
func (inst *CreateMasterEdition) SetMetadataMintAccount(metadataMint ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(metadataMint).WRITE()
	return inst
}

// GetMetadataMintAccount gets the "metadataMint" account.
// Metadata mint
func (inst *CreateMasterEdition) GetMetadataMintAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetUpdateAuthorityAccount sets the "updateAuthority" account.
// Update authority
func (inst *CreateMasterEdition) SetUpdateAuthorityAccount(updateAuthority ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(updateAuthority).SIGNER()
	return inst
}

// GetUpdateAuthorityAccount gets the "updateAuthority" account.
// Update authority
func (inst *CreateMasterEdition) GetUpdateAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetMintAuthorityAccount sets the "mintAuthority" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *CreateMasterEdition) SetMintAuthorityAccount(mintAuthority ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(mintAuthority).SIGNER()
	return inst
}

// GetMintAuthorityAccount gets the "mintAuthority" account.
// Mint authority on the metadata's mint - THIS WILL TRANSFER AUTHORITY AWAY FROM THIS KEY
func (inst *CreateMasterEdition) GetMintAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetPayerAccount sets the "payer" account.
// payer
func (inst *CreateMasterEdition) SetPayerAccount(payer ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(payer).SIGNER()
	return inst
}

// GetPayerAccount gets the "payer" account.
// payer
func (inst *CreateMasterEdition) GetPayerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetMetadataAccount sets the "metadataAccount" account.
// Metadata account
func (inst *CreateMasterEdition) SetMetadataAccount(metadataAccount ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(metadataAccount)
	return inst
}

// GetMetadataAccount gets the "metadataAccount" account.
// Metadata account
func (inst *CreateMasterEdition) GetMetadataAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetTokenProgramAccount sets the "tokenProgram" account.
// Token program
func (inst *CreateMasterEdition) SetTokenProgramAccount(tokenProgram ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(tokenProgram)
	return inst
}

// GetTokenProgramAccount gets the "tokenProgram" account.
// Token program
func (inst *CreateMasterEdition) GetTokenProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetSystemProgramAccount sets the "systemProgram" account.
// System program
func (inst *CreateMasterEdition) SetSystemProgramAccount(systemProgram ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(systemProgram)
	return inst
}

// GetSystemProgramAccount gets the "systemProgram" account.
// System program
func (inst *CreateMasterEdition) GetSystemProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

// SetRentInfoAccount sets the "rentInfo" account.
// Rent info
func (inst *CreateMasterEdition) SetRentInfoAccount(rentInfo ag_solanago.PublicKey) *CreateMasterEdition {
	inst.AccountMetaSlice[8] = ag_solanago.Meta(rentInfo)
	return inst
}

// GetRentInfoAccount gets the "rentInfo" account.
// Rent info
func (inst *CreateMasterEdition) GetRentInfoAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[8]
}

func (inst CreateMasterEdition) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_CreateMasterEdition),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst CreateMasterEdition) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *CreateMasterEdition) Validate() error {
	// Check whether all (required) parameters are set:
	{
		if inst.Args == nil {
			return errors.New("Args parameter is not set")
		}
	}

	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.UnallocatedEditionV2 is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.MetadataMint is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.UpdateAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.MintAuthority is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Payer is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.MetadataAccount is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.TokenProgram is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.SystemProgram is not set")
		}
		if inst.AccountMetaSlice[8] == nil {
			return errors.New("accounts.RentInfo is not set")
		}
	}
	return nil
}

func (inst *CreateMasterEdition) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("CreateMasterEdition")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=1]").ParentFunc(func(paramsBranch ag_treeout.Branches) {
						paramsBranch.Child(ag_format.Param("Args", *inst.Args))
					})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=9]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("unallocatedEditionV2", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("        metadataMint", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("     updateAuthority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("       mintAuthority", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("               payer", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("            metadata", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("        tokenProgram", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("       systemProgram", inst.AccountMetaSlice[7]))
						accountsBranch.Child(ag_format.Meta("            rentInfo", inst.AccountMetaSlice[8]))
					})
				})
		})
}

func (obj CreateMasterEdition) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	// Serialize `Args` param:
	err = encoder.Encode(obj.Args)
	if err != nil {
		return err
	}
	return nil
}
func (obj *CreateMasterEdition) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	// Deserialize `Args`:
	err = decoder.Decode(&obj.Args)
	if err != nil {
		return err
	}
	return nil
}

// NewCreateMasterEditionInstruction declares a new CreateMasterEdition instruction with the provided parameters and accounts.
func NewCreateMasterEditionInstruction(
	// Parameters:
	args CreateMasterEditionArgs,
	// Accounts:
	unallocatedEditionV2 ag_solanago.PublicKey,
	metadataMint ag_solanago.PublicKey,
	updateAuthority ag_solanago.PublicKey,
	mintAuthority ag_solanago.PublicKey,
	payer ag_solanago.PublicKey,
	metadataAccount ag_solanago.PublicKey,
	tokenProgram ag_solanago.PublicKey,
	systemProgram ag_solanago.PublicKey,
	rentInfo ag_solanago.PublicKey) *CreateMasterEdition {
	return NewCreateMasterEditionInstructionBuilder().
		SetArgs(args).
		SetUnallocatedEditionV2Account(unallocatedEditionV2).
		SetMetadataMintAccount(metadataMint).
		SetUpdateAuthorityAccount(updateAuthority).
		SetMintAuthorityAccount(mintAuthority).
		SetPayerAccount(payer).
		SetMetadataAccount(metadataAccount).
		SetTokenProgramAccount(tokenProgram).
		SetSystemProgramAccount(systemProgram).
		SetRentInfoAccount(rentInfo)
}