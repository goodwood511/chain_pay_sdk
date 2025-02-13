package response_define

type ResponseQueryTokenList struct {
	Msg       string               `json:"msg"`
	Code      int64                `json:"code"`
	Data      []QueryTokenListData `json:"data"`
	Sign      string               `json:"sign"`
	Timestamp string               `json:"timestamp"`
}

type QueryTokenListData struct {
	Symbol             string  `json:"symbol"`
	Autocheck          bool    `json:"autocheck"`
	Tokenid            int64   `json:"tokenid"`
	Ordermax           string  `json:"ordermax"`
	Payoutamount       string  `json:"payoutamount"`
	CollectEnable      int64   `json:"collectEnable"`
	CollectFee         string  `json:"collectFee"`
	PartnerFeePaid     string  `json:"partnerFeePaid"`
	Paylimitday        string  `json:"paylimitday"`
	ParnterFeeTotal    string  `json:"parnterFeeTotal"`
	Ordercount         int64   `json:"ordercount"`
	Coldaddress1       *string `json:"coldaddress1,omitempty"`
	TokenAddress       string  `json:"tokenAddress"`
	Checkmin           string  `json:"checkmin"`
	Paytotalday        string  `json:"paytotalday"`
	Paylimitone        string  `json:"paylimitone"`
	Statuscharge       int64   `json:"statuscharge"`
	Payinamount        string  `json:"payinamount"`
	ChainBalance       string  `json:"chainBalance"`
	ID                 int64   `json:"id"`
	Partnerid          int64   `json:"partnerid"`
	Txcallback         string  `json:"txcallback"`
	Feeaddress         *string `json:"feeaddress,omitempty"`
	Chainnodeid        int64   `json:"chainnodeid"`
	CollectTime        int64   `json:"collectTime"`
	CollectWaitBalance string  `json:"collectWaitBalance"`
	Payoutquota        string  `json:"payoutquota"`
	Coldaddress        string  `json:"coldaddress"`
	Allownew           int64   `json:"allownew"`
	TokenPrice         string  `json:"tokenPrice"`
	Collectfee         string  `json:"collectfee"`
	CollectBalance     string  `json:"collectBalance"`
	UpdateTime         string  `json:"updateTime"`
	Statuspay          int64   `json:"statuspay"`
	Activetime         string  `json:"activetime"`
	CollectType        int64   `json:"collectType"`
	TokenDecimals      int64   `json:"tokenDecimals"`
	Payoutaddress      string  `json:"payoutaddress"`
	Paycharge          string  `json:"paycharge"`
	Feecharge          string  `json:"feecharge"`
	CollectMin         string  `json:"collectMin"`
	ChainUpdateTime    string  `json:"chainUpdateTime"`
	Status             int64   `json:"status"`
	TokenBalance       string  `json:"tokenBalance"`
}
