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

	reqBody, timestamp, sign, clientSign, err := apiObj.CreateUser(openId)
	if err != nil {
		logrus.Fatalf("Failed to prepare request: %v", err)
	}

	var resp response_define.ResponseCreateUser
	err = utils.CallAndVerify(apiObj, api.PathCreateUser, reqBody, timestamp, sign, clientSign, &resp)
	if err != nil {
		logrus.Fatalf("API call failed: %v", err)
	}

	logrus.Infof("CreateUser Success: %+v", resp)
}
