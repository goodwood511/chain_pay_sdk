package api

import (
	"github.com/goodwood511/chain_pay_sdk/request_define"
)

// CreateWallet create wallet
// @param openId user open id
// @param chainID chain id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) CreateWallet(openId, chainID string) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestCreateWallet{
			OpenID:  openId,
			ChainID: chainID,
		},
	)
}
