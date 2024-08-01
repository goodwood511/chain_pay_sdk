package response_define

type ResponseCreateUserData struct {
	OpenId string `json:"OpenId"`
}

type ResponseCreateUser struct {
	Sign      string                 `json:"sign"`
	Timestamp string                 `json:"timestamp"`
	Code      int                    `json:"Code"`
	Msg       string                 `json:"Msg"`
	Data      ResponseCreateUserData `json:"Data"`
}
