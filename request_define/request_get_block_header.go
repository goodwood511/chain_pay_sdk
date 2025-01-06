package request_define

type RequestGetBlockHeader struct {
	Block   string `json:"block"`
	ChainID string `json:"chainId"`
}
