package offchain

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// testData contains the off-chain metadata of a random NFT on Solana.
//
// Token Metadata Account: https://explorer.solana.com/address/GfihrEYCPrvUyrMyMQPdhGEStxa9nKEK2Wfn9iK4AZq2
// Arweave: https://fs6f64jxxtd6hbvrydkag5orgv5c5xxg3fmazr5k4ziicoqmatuq.arweave.net/LLxfcTe8x-OGscDUA3XRNXou3ubZWAzHquZQgToMBOk/
var testData = `{"name":"Zuuper Grapes","symbol":"TOILET","description":"In the year 3035 when humanity and the earth are falling and almost ravaged, a fruit is found as a sign of hope of a new civilization. The fruit is called Zuuper Grapes. /// Soilets are digital collectibles of 30 unique audiovisual Toilets vibing on the Solana blockchain. As a homage to the greatest sanitation revolution in the past 200 years and as the symbol of the degradation of the inner self, each Toilet has 5 editions of 5 minted biweekly until the end of November 2021 in the auction format. /// Visual by Peul & Audio by Bagvs.","seller_fee_basis_points":1000,"image":"https://www.arweave.net/eAmrcRbfeAzO8VYoKYe6eKtL--Dn9y7EuhS2fWevGXo?ext=png","animation_url":"https://www.arweave.net/U-IINiksuUdmUz2JXDYseVKidsqIHVEC1Sp_7zrbSpo?ext=mp4","attributes":[{"trait_type":"Toilet Type","value":"Zuuper Grapes"},{"trait_type":"Visual Artist","value":"Peul"},{"trait_type":"Audio Artist","value":"Bagvs"},{"display_type":"number","trait_type":"Batch","value":1}],"collection":{"name":"Soilets Batch #1","family":"Soilets"},"external_url":"https://soilets.art","properties":{"files":[{"uri":"https://www.arweave.net/eAmrcRbfeAzO8VYoKYe6eKtL--Dn9y7EuhS2fWevGXo?ext=png","type":"image/png"},{"uri":"https://www.arweave.net/U-IINiksuUdmUz2JXDYseVKidsqIHVEC1Sp_7zrbSpo?ext=mp4","type":"video/mp4"}],"category":"video","creators":[{"address":"BmpjtUChunwTUqkSPaq3YvZ2nLnF9SScTeUZ4Mc6CLrq","share":43},{"address":"FTP6tvEaKXi9rDW1y8NdHd4GpfMGcJKJVLJdBaXJpr9j","share":28},{"address":"3WzjQZeBybrMn1jGskeAhnjX2m4RTUPXDkDDTJJmjcVZ","share":10},{"address":"6ftJ84D4rgJDCyna5cPWPguiWPpGTYrb4uuhocBCQmXq","share":19}]}}`

func TestDataUnmarshal(t *testing.T) {
	var actual NFT
	require.NoError(t, json.Unmarshal([]byte(testData), &actual))
	expected := &NFT{
		Name:                 "Zuuper Grapes",
		Symbol:               "TOILET",
		Description:          "In the year 3035 when humanity and the earth are falling and almost ravaged, a fruit is found as a sign of hope of a new civilization. The fruit is called Zuuper Grapes. /// Soilets are digital collectibles of 30 unique audiovisual Toilets vibing on the Solana blockchain. As a homage to the greatest sanitation revolution in the past 200 years and as the symbol of the degradation of the inner self, each Toilet has 5 editions of 5 minted biweekly until the end of November 2021 in the auction format. /// Visual by Peul & Audio by Bagvs.",
		SellerFeeBasisPoints: 1000,
		Image:                "https://www.arweave.net/eAmrcRbfeAzO8VYoKYe6eKtL--Dn9y7EuhS2fWevGXo?ext=png",
		AnimationURL:         "https://www.arweave.net/U-IINiksuUdmUz2JXDYseVKidsqIHVEC1Sp_7zrbSpo?ext=mp4",
		Attributes: []Attributes{
			{
				TraitType: "Toilet Type",
				Value:     "Zuuper Grapes",
			},
			{
				TraitType: "Visual Artist",
				Value:     "Peul",
			},
			{
				TraitType: "Audio Artist",
				Value:     "Bagvs",
			},
			{
				TraitType:   "Batch",
				Value:       float64(1),
				DisplayType: "number",
			},
		},
		Collection: &Collection{
			Name:   "Soilets Batch #1",
			Family: "Soilets",
		},
		ExternalURL: "https://soilets.art",
		Properties: Properties{
			Files: []File{
				{
					URI:  "https://www.arweave.net/eAmrcRbfeAzO8VYoKYe6eKtL--Dn9y7EuhS2fWevGXo?ext=png",
					Type: "image/png",
				},
				{
					URI:  "https://www.arweave.net/U-IINiksuUdmUz2JXDYseVKidsqIHVEC1Sp_7zrbSpo?ext=mp4",
					Type: "video/mp4",
				},
			},
			Category: "video",
			Creators: []Creator{
				{
					Address: "BmpjtUChunwTUqkSPaq3YvZ2nLnF9SScTeUZ4Mc6CLrq",
					Share:   43,
				},
				{
					Address: "FTP6tvEaKXi9rDW1y8NdHd4GpfMGcJKJVLJdBaXJpr9j",
					Share:   28,
				},
				{
					Address: "3WzjQZeBybrMn1jGskeAhnjX2m4RTUPXDkDDTJJmjcVZ",
					Share:   10,
				},
				{
					Address: "6ftJ84D4rgJDCyna5cPWPguiWPpGTYrb4uuhocBCQmXq",
					Share:   19,
				},
			},
		},
	}
	assert.Equal(t, expected, &actual)
}
