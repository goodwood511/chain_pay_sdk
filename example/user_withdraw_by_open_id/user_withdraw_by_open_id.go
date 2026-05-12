package main

import (
	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/example/utils"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/sirupsen/logrus"
)

func main() {
	apiObj := utils.InitSDK()

	openId := "HASH13900000010"
	tokenId := "4"
	amount := "0.01"
	addressTo := "TQdL5yttJPTx7hJmBhGfo2LcE7AXLPtHSg"
	callbackUrl := " "
	safeCheckCode := "202408132056333"

	reqBody, timestamp, sign, clientSign, err := apiObj.UserWithdrawByOpenID(openId,
		tokenId,
		amount,
		addressTo,
		callbackUrl,
		safeCheckCode,
	)
	if err != nil {
		logrus.Fatalf("Failed to prepare request: %v", err)
	}

	var resp response_define.ResponseUserWithdrawByOpenID
	err = utils.CallAndVerify(apiObj, api.PathUserWithdrawByOpenID, reqBody, timestamp, sign, clientSign, &resp)
	if err != nil {
		logrus.Fatalf("API call failed: %v", err)
	}

	logrus.Infof("UserWithdrawByOpenID Success: %+v", resp)
}
