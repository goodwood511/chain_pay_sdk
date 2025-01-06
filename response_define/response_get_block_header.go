package response_define

type ResponseGetBlockHeader struct {
	Msg       string             `json:"msg"`
	Code      int64              `json:"code"`
	Data      GetBlockHeaderData `json:"data"`
	Sign      string             `json:"sign"`
	Timestamp string             `json:"timestamp"`
}

type GetBlockHeaderData struct {
	WithdrawAmount  int64                 `json:"withdraw_amount"`
	Log             []GetBlockHeaderLog   `json:"log"`
	Fee             int64                 `json:"fee"`
	BlockNumber     int64                 `json:"blockNumber"`
	ContractResult  []string              `json:"contractResult"`
	BlockTimeStamp  int64                 `json:"blockTimeStamp"`
	Receipt         GetBlockHeaderReceipt `json:"receipt"`
	ID              string                `json:"id"`
	ContractAddress string                `json:"contract_address"`
}

type GetBlockHeaderLog struct {
	Address string   `json:"address"`
	Data    string   `json:"data"`
	Topics  []string `json:"topics"`
}

type GetBlockHeaderReceipt struct {
	Result             string `json:"result"`
	EnergyPenaltyTotal int64  `json:"energy_penalty_total"`
	EnergyUsage        int64  `json:"energy_usage"`
	EnergyUsageTotal   int64  `json:"energy_usage_total"`
	NetUsage           int64  `json:"net_usage"`
}
