package response_define

type ResponseUserWithdrawByOpenID struct {
	Sign      string                           `json:"sign"`
	Timestamp string                           `json:"timestamp"`
	Code      int                              `json:"Code"`
	Msg       string                           `json:"Msg"`
	Data      ResponseUserWithdrawByOpenIDData `json:"data"`
}

type ResponseUserWithdrawByOpenIDData struct {
	UserPayChargeRecords []UserPayChargeRecords `json:"userPayChargeRecords"`
	TotalCount           int                    `json:"TotalCount"`
}

type UserPayChargeRecords struct {
	OpenID        string  `json:"OpenID"`
	PayOrCharge   int     `json:"PayOrCharge"`
	FromAddress   string  `json:"FromAddress"`
	ToAddress     string  `json:"ToAddress"`
	HashCode      string  `json:"HashCode"`
	SafeCode      string  `json:"SafeCode"`
	Amount        float64 `json:"Amount"`
	Status        int     `json:"Status"`
	NoticeTimes   int     `json:"NoticeTimes"`
	NoticeURL     string  `json:"NoticeUrl"`
	NoticeRespone string  `json:"NoticeRespone"`
	CreateTime    string  `json:"CreateTime"`
}
