package token_utils

import "fmt"

type TokenInfo struct {
	ChainID      int64
	TokenID      int64
	TokenAddress string
	Symbol       string
	Decimal      int16
}

var (
	mapTokenInfo = make(map[string]*TokenInfo)
)

func GetTokenInfo(chainID, tokenId int64) *TokenInfo {
	key := GenerateKey(chainID, tokenId)
	if value, ok := mapTokenInfo[key]; ok {
		return value
	}
	return nil
}

func GenerateKey(chainID, tokenId int64) string {
	return fmt.Sprintf("%d_%d", chainID, tokenId)
}
