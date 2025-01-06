package api

import "github.com/goodwood511/chain_pay_sdk/request_define"

// GetTXDetail get transfer detail
// @param txHash transfer hash
// @param chainID chain id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) GetTXDetail(txHash, chainID string) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestGetTXDetail{
			TxHash:  txHash,
			ChainID: chainID,
		},
	)
}
