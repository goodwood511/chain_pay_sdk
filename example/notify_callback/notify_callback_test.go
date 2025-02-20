package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func TestCallbackVerify(t *testing.T) {

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})

	body := []byte(viper.GetString("CBRsp"))

	req := response_define.RequestTokenCb{}
	fmt.Println("Raw JSON:", string(body))
	err := json.Unmarshal(body, &req)

	mapData := make(map[string]string)

	err = json.Unmarshal(body, &mapData)
	if err != nil {
		logrus.Warnln("json.Unmarshal fail, err", err.Error())
		return
	}

	err = apiObj.VerifyRSAsignature(mapData, req.Sign)
	if err != nil {
		logrus.Warnln("VerifyRSAsignature fail, err", err.Error())
		return
	}
	logrus.Infoln(req)

}
