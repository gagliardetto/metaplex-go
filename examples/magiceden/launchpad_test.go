package magiceden

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gagliardetto/metaplex-go/clients/nft_magiceden"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/rpc"
	"testing"
)

func Test_GetAccountInfo(t *testing.T) {
	client := rpc.New(rpc.MainNetBeta_RPC)
	candyAddress, _ := solana.PublicKeyFromBase58("")
	//programId, _ := solana.PublicKeyFromBase58("")
	var cm nft_magiceden.CandyMachine
	err := client.GetAccountDataBorshInto(context.TODO(), candyAddress, &cm)
	if err != nil {
		panic(err)
	}
	b, _ := json.Marshal(cm)
	fmt.Println(string(b))
}
