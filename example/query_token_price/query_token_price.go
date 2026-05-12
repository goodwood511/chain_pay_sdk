package main

import (
	"encoding/json"

	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/example/utils"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/sirupsen/logrus"
)

func main() {
	apiObj := utils.InitSDK()

	tokens := "BTC,ETH,TRX"

	reqBody, timestamp, sign, clientSign, err := apiObj.GetTokenPrice(tokens)
	if err != nil {
		logrus.Fatalf("Failed to prepare request: %v", err)
	}

	finalURL, _ := utils.CallAPIPath(api.PathGetTokenPrice)
	resp, err := utils.Client.R().
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		SetBody(reqBody).
		Post(finalURL)

	if err != nil {
		logrus.Fatalf("API call failed: %v", err)
	}

	var prices []response_define.ResponseQueryTokenPrice
	if err := json.Unmarshal(resp.Body(), &prices); err != nil {
		logrus.Fatalf("Failed to unmarshal response: %v", err)
	}

	logrus.Infof("Token Prices: %+v", prices)
}
