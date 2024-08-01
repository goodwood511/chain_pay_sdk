package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/goodwood511/chain_pay_sdk/rsa_utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	r := gin.Default()

	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to load config: %s", err))
	}
	apiObj := api.NewSDK(api.SDKConfig{
		ApiKey:         viper.GetString("ApiKey"),
		ApiSecret:      viper.GetString("ApiSecret"),
		PlatformPubKey: viper.GetString("PlatformPubKey"),
		RsaPrivateKey:  viper.GetString("RsaPrivateKey"),
	})

	r.POST("/withdraw_callback", func(c *gin.Context) {
		req := response_define.RequestTokenCb{}
		err := c.ShouldBindJSON(&req)
		if err != nil {
			logrus.Warnln("/withdraw_callback fail")
			response_define.FailWithMessage("parse json param fail "+err.Error(), c)
			return
		}

		mapData, err := rsa_utils.StructToMap(req)
		if err != nil {
			logrus.Warnln("StructToMap fail, err", err.Error())
			response_define.FailWithMessage("StructToMap fail "+err.Error(), c)
			return
		}

		err = apiObj.VerifyRSAsignature(mapData, req.Sign)
		if err != nil {
			logrus.Warnln("VerifyRSAsignature fail, err", err.Error())
			response_define.FailWithMessage("verify RSA signature fail "+err.Error(), c)
			return
		}
		logrus.Infoln(req)
		response_define.Ok(c)
	})
	r.Run(":9003")
}
