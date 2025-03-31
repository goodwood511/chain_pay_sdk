package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	client := resty.New()

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
	timestamp := strconv.FormatInt(time.Now().UnixMilli(), 10)
	sign := apiObj.GenerateMD5Sign("", timestamp)

	var Tokens string = "BTC,ETH,TRX"

	reqBody, timestamp, sign, clientSign, err := apiObj.GetTokenPrice(Tokens)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	fmt.Println("reqBody: ", string(reqBody))

	finalURL, err := url.JoinPath(api.DevNetEndpoint, api.PathGetTokenPrice)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}

	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetBody(reqBody).
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		Post(finalURL)

	if err != nil {
		logrus.Warnln("Error: ", err, "finalURL", finalURL)
		return
	}

	body := resp.Body()
	fmt.Println(string(body))

	rspQueryTokenList := []response_define.ResponseQueryTokenPrice{}
	err = json.Unmarshal(body, &rspQueryTokenList)
	if err != nil {
		logrus.Warnln("Error: ", err)
		return
	}
	logrus.Infoln("ResponseQueryTokenList: ", rspQueryTokenList)

}
