package token_utils

func init() {
	mapTokenInfo[GenerateKey(ChainID_BNB, TokenID_BNB)] = &TokenInfo{
		ChainID:      ChainID_BNB,
		TokenID:      TokenID_BNB,
		TokenAddress: TokenSymbol_BNB,
		Symbol:       TokenSymbol_BNB,
		Decimal:      18,
	}

	mapTokenInfo[GenerateKey(ChainID_BNB, TokenID_BNB_USDT)] = &TokenInfo{
		ChainID:      ChainID_BNB,
		TokenID:      TokenID_BNB_USDT,
		TokenAddress: TokenAddress_BNB_USDT,
		Symbol:       TokenSymbol_BNB_USDT,
		Decimal:      18,
	}
	mapTokenInfo[GenerateKey(ChainID_BNB, TokenID_BNB_USDC)] = &TokenInfo{
		ChainID:      ChainID_BNB,
		TokenID:      TokenID_BNB_USDC,
		TokenAddress: TokenAddress_BNB_USDC,
		Symbol:       TokenSymbol_BNB_USDC,
		Decimal:      18,
	}
}
