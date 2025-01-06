package response_define

type ResponseTXDetail struct {
	Msg       string `json:"msg"`
	Code      int64  `json:"code"`
	Data      Data   `json:"data"`
	Sign      string `json:"sign"`
	Timestamp string `json:"timestamp"`
}

type Data struct {
	BlockID     string      `json:"blockID"`
	BlockHeader BlockHeader `json:"block_header"`
}

type BlockHeader struct {
	RawData          RawData `json:"raw_data"`
	WitnessSignature string  `json:"witness_signature"`
}

type RawData struct {
	Number         int64  `json:"number"`
	TxTrieRoot     string `json:"txTrieRoot"`
	WitnessAddress string `json:"witness_address"`
	ParentHash     string `json:"parentHash"`
	Version        int64  `json:"version"`
	Timestamp      int64  `json:"timestamp"`
}
