package token_utils

func init() {
	mapTokenInfo[GenerateKey(ChainID_BNB, TokenID_BNB)] = &TokenInfo{
		ChainID:      ChainID_BNB,
		TokenID:      TokenID_BNB,
		TokenAddress: TokenSymbol_BNB,
		Symbol:       TokenSymbol_BNB,
		Decimal:      18,
	}

}
