// Package offchain describes Metaplex off-chain NFT metadata.
//
// Docs: https://docs.metaplex.com/nft-standard#uri-json-schema
package offchain

// NFT is the root level struct of NFT metadata.
type NFT struct {
	Name                 string       `json:"name"`
	Symbol               string       `json:"symbol"`
	Description          string       `json:"description"`
	SellerFeeBasisPoints int          `json:"seller_fee_basis_points"`
	Image                string       `json:"image"`
	AnimationURL         string       `json:"animation_url"`
	Attributes           []Attributes `json:"attributes,omitempty"`
	Collection           *Collection  `json:"collection,omitempty"`
	ExternalURL          string       `json:"external_url"`
	Properties           Properties   `json:"properties"`
}

type Attributes struct {
	TraitType   string      `json:"trait_type"`
	Value       interface{} `json:"value"`
	DisplayType string      `json:"display_type,omitempty"`
}

type Collection struct {
	Name   string `json:"name"`
	Family string `json:"family"`
}

type File struct {
	URI  string `json:"uri"`
	Type string `json:"type"`
	CDN  bool   `json:"cdn,omitempty"`
}

type Creator struct {
	Address string `json:"address"`
	Share   int    `json:"share"`
}

type Properties struct {
	Files    []File    `json:"files"`
	Category string    `json:"category"`
	Creators []Creator `json:"creators"`
}
