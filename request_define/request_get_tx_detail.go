package request_define

type RequestGetTXDetail struct {
	TxHash  string `json:"txHash"`
	ChainID string `json:"chainId"`
}
