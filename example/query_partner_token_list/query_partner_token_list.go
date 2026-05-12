package main

import (
	"strconv"
	"time"

	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/example/utils"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/sirupsen/logrus"
)

func main() {
	apiObj := utils.InitSDK()

	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign := apiObj.GenerateMD5Sign("", timestamp)

	var resp response_define.ResponseQueryTokenList
	err := utils.CallAndVerifyMD5(apiObj, api.PathQueryTokenList, timestamp, sign, &resp)
	if err != nil {
		logrus.Fatalf("API call failed: %v", err)
	}

	logrus.Infof("QueryTokenList Success: %+v", resp)
}
