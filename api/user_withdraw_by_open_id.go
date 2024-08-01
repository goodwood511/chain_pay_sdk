package api

import (
	"strconv"

	"github.com/goodwood511/chain_pay_sdk/request_define"
)

// UserWithdrawByOpenID user withdraw by openId
// @param tokenId token
// @param amount Amount of tokens
// @param addressTo address to
// @param callbackUrl callback url
// @param safeCheckCode safe check code
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) UserWithdrawByOpenID(openId string,
	tokenId int64,
	amount float64,
	addressTo,
	callbackUrl,
	safeCheckCode string,
) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestWithdrawByOpenID{
			OpenID:        openId,
			TokenID:       strconv.FormatInt(tokenId, 10),
			Amount:        strconv.FormatFloat(amount, 'f', 2, 64),
			AddressTo:     addressTo,
			CallBackURL:   callbackUrl,
			SafeCheckCode: safeCheckCode,
		},
	)
}
