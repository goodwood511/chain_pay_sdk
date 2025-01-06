package api

import "github.com/goodwood511/chain_pay_sdk/request_define"

// GetBlockHeader get block header
// @param block block
// @param chainID chain id
// @return data, timestamp, sign, clientSign, error
func (s *Sdk) GetBlockHeader(block, chainID string) ([]byte, string, string, string, error) {

	return s.signPack(
		request_define.RequestGetBlockHeader{
			Block:   block,
			ChainID: chainID,
		},
	)
}
