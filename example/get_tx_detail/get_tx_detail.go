package main

import (
	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/example/utils"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/sirupsen/logrus"
)

func main() {
	apiObj := utils.InitSDK()

	txHash := "bdb327a18d543d2307525d433a1649e2525e74c677046583f7577311c6a3cb9f"
	chainID := "2"

	reqBody, timestamp, sign, clientSign, err := apiObj.GetTXDetail(txHash, chainID)
	if err != nil {
		logrus.Fatalf("Failed to prepare request: %v", err)
	}

	var resp response_define.ResponseTXDetail
	err = utils.CallAndVerify(apiObj, api.PathGetTXDetail, reqBody, timestamp, sign, clientSign, &resp)
	if err != nil {
		logrus.Fatalf("API call failed: %v", err)
	}

	logrus.Infof("GetTXDetail Success: %+v", resp)
}
