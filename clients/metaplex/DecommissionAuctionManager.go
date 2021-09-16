// Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.

package metaplex

import (
	"errors"
	ag_binary "github.com/gagliardetto/binary"
	ag_solanago "github.com/gagliardetto/solana-go"
	ag_format "github.com/gagliardetto/solana-go/text/format"
	ag_treeout "github.com/gagliardetto/treeout"
)

// If you have an auction manager in an Initialized state and for some reason you can't validate it, you want to retrieve
// The items inside of it. This will allow you to move it straight to Disbursing, and then you can, as Auctioneer,
// Redeem those items using the RedeemUnusedWinningConfigItemsAsAuctioneer endpoint.
//
// If you pass the vault program account, authority over the vault will be returned to you, so you can unwind the vault
// to get your items back that way instead.
//
// Be WARNED: Because the boxes have not been validated, the logic for redemptions may not work quite right. For instance,
// if your validation step failed because you provided an empty box but said there was a token in it, when you go
// and try to redeem it, you yourself will experience quite the explosion. It will be up to you to tactfully
// request the bids that can be properly redeemed from the ones that cannot.
//
// If you had a FullRightsTransfer token, and you never validated (and thus transferred) ownership, when the redemption happens
// it will skip trying to transfer it to you, so that should work fine.
type DecommissionAuctionManager struct {

	// [0] = [WRITE] auctionManager
	// ··········· Auction Manager
	//
	// [1] = [WRITE] auction
	// ··········· Auction
	//
	// [2] = [SIGNER] auctionManagerAuthority
	// ··········· Authority of the Auction Manager
	//
	// [3] = [WRITE] vault
	// ··········· Vault
	//
	// [4] = [] store
	// ··········· Store
	//
	// [5] = [] auctionProgram
	// ··········· Auction program
	//
	// [6] = [] clockSysvar
	// ··········· Clock sysvar
	//
	// [7] = [] vaultProgram
	// ··········· Vault program (Optional)
	ag_solanago.AccountMetaSlice `bin:"-" borsh_skip:"true"`
}

// NewDecommissionAuctionManagerInstructionBuilder creates a new `DecommissionAuctionManager` instruction builder.
func NewDecommissionAuctionManagerInstructionBuilder() *DecommissionAuctionManager {
	nd := &DecommissionAuctionManager{
		AccountMetaSlice: make(ag_solanago.AccountMetaSlice, 8),
	}
	return nd
}

// SetAuctionManagerAccount sets the "auctionManager" account.
// Auction Manager
func (inst *DecommissionAuctionManager) SetAuctionManagerAccount(auctionManager ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[0] = ag_solanago.Meta(auctionManager).WRITE()
	return inst
}

// GetAuctionManagerAccount gets the "auctionManager" account.
// Auction Manager
func (inst *DecommissionAuctionManager) GetAuctionManagerAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[0]
}

// SetAuctionAccount sets the "auction" account.
// Auction
func (inst *DecommissionAuctionManager) SetAuctionAccount(auction ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[1] = ag_solanago.Meta(auction).WRITE()
	return inst
}

// GetAuctionAccount gets the "auction" account.
// Auction
func (inst *DecommissionAuctionManager) GetAuctionAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[1]
}

// SetAuctionManagerAuthorityAccount sets the "auctionManagerAuthority" account.
// Authority of the Auction Manager
func (inst *DecommissionAuctionManager) SetAuctionManagerAuthorityAccount(auctionManagerAuthority ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[2] = ag_solanago.Meta(auctionManagerAuthority).SIGNER()
	return inst
}

// GetAuctionManagerAuthorityAccount gets the "auctionManagerAuthority" account.
// Authority of the Auction Manager
func (inst *DecommissionAuctionManager) GetAuctionManagerAuthorityAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[2]
}

// SetVaultAccount sets the "vault" account.
// Vault
func (inst *DecommissionAuctionManager) SetVaultAccount(vault ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[3] = ag_solanago.Meta(vault).WRITE()
	return inst
}

// GetVaultAccount gets the "vault" account.
// Vault
func (inst *DecommissionAuctionManager) GetVaultAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[3]
}

// SetStoreAccount sets the "store" account.
// Store
func (inst *DecommissionAuctionManager) SetStoreAccount(store ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[4] = ag_solanago.Meta(store)
	return inst
}

// GetStoreAccount gets the "store" account.
// Store
func (inst *DecommissionAuctionManager) GetStoreAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[4]
}

// SetAuctionProgramAccount sets the "auctionProgram" account.
// Auction program
func (inst *DecommissionAuctionManager) SetAuctionProgramAccount(auctionProgram ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[5] = ag_solanago.Meta(auctionProgram)
	return inst
}

// GetAuctionProgramAccount gets the "auctionProgram" account.
// Auction program
func (inst *DecommissionAuctionManager) GetAuctionProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[5]
}

// SetClockSysvarAccount sets the "clockSysvar" account.
// Clock sysvar
func (inst *DecommissionAuctionManager) SetClockSysvarAccount(clockSysvar ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[6] = ag_solanago.Meta(clockSysvar)
	return inst
}

// GetClockSysvarAccount gets the "clockSysvar" account.
// Clock sysvar
func (inst *DecommissionAuctionManager) GetClockSysvarAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[6]
}

// SetVaultProgramAccount sets the "vaultProgram" account.
// Vault program (Optional)
func (inst *DecommissionAuctionManager) SetVaultProgramAccount(vaultProgram ag_solanago.PublicKey) *DecommissionAuctionManager {
	inst.AccountMetaSlice[7] = ag_solanago.Meta(vaultProgram)
	return inst
}

// GetVaultProgramAccount gets the "vaultProgram" account.
// Vault program (Optional)
func (inst *DecommissionAuctionManager) GetVaultProgramAccount() *ag_solanago.AccountMeta {
	return inst.AccountMetaSlice[7]
}

func (inst DecommissionAuctionManager) Build() *Instruction {
	return &Instruction{BaseVariant: ag_binary.BaseVariant{
		Impl:   inst,
		TypeID: ag_binary.TypeIDFromUint8(Instruction_DecommissionAuctionManager),
	}}
}

// ValidateAndBuild validates the instruction parameters and accounts;
// if there is a validation error, it returns the error.
// Otherwise, it builds and returns the instruction.
func (inst DecommissionAuctionManager) ValidateAndBuild() (*Instruction, error) {
	if err := inst.Validate(); err != nil {
		return nil, err
	}
	return inst.Build(), nil
}

func (inst *DecommissionAuctionManager) Validate() error {
	// Check whether all (required) accounts are set:
	{
		if inst.AccountMetaSlice[0] == nil {
			return errors.New("accounts.AuctionManager is not set")
		}
		if inst.AccountMetaSlice[1] == nil {
			return errors.New("accounts.Auction is not set")
		}
		if inst.AccountMetaSlice[2] == nil {
			return errors.New("accounts.AuctionManagerAuthority is not set")
		}
		if inst.AccountMetaSlice[3] == nil {
			return errors.New("accounts.Vault is not set")
		}
		if inst.AccountMetaSlice[4] == nil {
			return errors.New("accounts.Store is not set")
		}
		if inst.AccountMetaSlice[5] == nil {
			return errors.New("accounts.AuctionProgram is not set")
		}
		if inst.AccountMetaSlice[6] == nil {
			return errors.New("accounts.ClockSysvar is not set")
		}
		if inst.AccountMetaSlice[7] == nil {
			return errors.New("accounts.VaultProgram is not set")
		}
	}
	return nil
}

func (inst *DecommissionAuctionManager) EncodeToTree(parent ag_treeout.Branches) {
	parent.Child(ag_format.Program(ProgramName, ProgramID)).
		//
		ParentFunc(func(programBranch ag_treeout.Branches) {
			programBranch.Child(ag_format.Instruction("DecommissionAuctionManager")).
				//
				ParentFunc(func(instructionBranch ag_treeout.Branches) {

					// Parameters of the instruction:
					instructionBranch.Child("Params[len=0]").ParentFunc(func(paramsBranch ag_treeout.Branches) {})

					// Accounts of the instruction:
					instructionBranch.Child("Accounts[len=8]").ParentFunc(func(accountsBranch ag_treeout.Branches) {
						accountsBranch.Child(ag_format.Meta("         auctionManager", inst.AccountMetaSlice[0]))
						accountsBranch.Child(ag_format.Meta("                auction", inst.AccountMetaSlice[1]))
						accountsBranch.Child(ag_format.Meta("auctionManagerAuthority", inst.AccountMetaSlice[2]))
						accountsBranch.Child(ag_format.Meta("                  vault", inst.AccountMetaSlice[3]))
						accountsBranch.Child(ag_format.Meta("                  store", inst.AccountMetaSlice[4]))
						accountsBranch.Child(ag_format.Meta("         auctionProgram", inst.AccountMetaSlice[5]))
						accountsBranch.Child(ag_format.Meta("            clockSysvar", inst.AccountMetaSlice[6]))
						accountsBranch.Child(ag_format.Meta("           vaultProgram", inst.AccountMetaSlice[7]))
					})
				})
		})
}

func (obj DecommissionAuctionManager) MarshalWithEncoder(encoder *ag_binary.Encoder) (err error) {
	return nil
}
func (obj *DecommissionAuctionManager) UnmarshalWithDecoder(decoder *ag_binary.Decoder) (err error) {
	return nil
}

// NewDecommissionAuctionManagerInstruction declares a new DecommissionAuctionManager instruction with the provided parameters and accounts.
func NewDecommissionAuctionManagerInstruction(
	// Accounts:
	auctionManager ag_solanago.PublicKey,
	auction ag_solanago.PublicKey,
	auctionManagerAuthority ag_solanago.PublicKey,
	vault ag_solanago.PublicKey,
	store ag_solanago.PublicKey,
	auctionProgram ag_solanago.PublicKey,
	clockSysvar ag_solanago.PublicKey,
	vaultProgram ag_solanago.PublicKey) *DecommissionAuctionManager {
	return NewDecommissionAuctionManagerInstructionBuilder().
		SetAuctionManagerAccount(auctionManager).
		SetAuctionAccount(auction).
		SetAuctionManagerAuthorityAccount(auctionManagerAuthority).
		SetVaultAccount(vault).
		SetStoreAccount(store).
		SetAuctionProgramAccount(auctionProgram).
		SetClockSysvarAccount(clockSysvar).
		SetVaultProgramAccount(vaultProgram)
}