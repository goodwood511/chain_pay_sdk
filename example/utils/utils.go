package utils

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goodwood511/chain_pay_sdk/api"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/goodwood511/chain_pay_sdk/rsa_utils"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var (
	// Client is a shared resty client with best-practice settings
	Client *resty.Client
)

func init() {
	Client = resty.New().
		SetTimeout(15 * time.Second).
		SetRetryCount(3).
		SetRetryWaitTime(2 * time.Second).
		SetHeader("Content-Type", "application/json")
}

// InitSDK loads config and initializes the SDK
func InitSDK() *api.Sdk {
	viper.SetConfigFile("config.yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("..")
	viper.AddConfigPath("../../")
	
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("Failed to load config: %s", err)
	}

	return api.NewSDK(api.SDKConfig{
		ApiKey:             viper.GetString("ApiKey"),
		ApiSecret:          viper.GetString("ApiSecret"),
		PlatformPubKey:     viper.GetString("PlatformPubKey"),
		PlatformRiskPubKey: viper.GetString("PlatformRiskPubKey"),
		RsaPrivateKey:      viper.GetString("RsaPrivateKey"),
	})
}

// CallAndVerify performs a signed request and verifies the response
func CallAndVerify(apiObj *api.Sdk, path string, reqBody []byte, timestamp, sign, clientSign string, target interface{}) error {
	finalURL, err := url.JoinPath(api.DevNetEndpoint, path)
	if err != nil {
		return fmt.Errorf("failed to join path: %w", err)
	}

	resp, err := Client.R().
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		SetHeader("clientSign", clientSign).
		SetBody(reqBody).
		Post(finalURL)

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return ProcessResponse(apiObj, resp.Body(), target, true)
}

// CallAndVerifyMD5 performs an MD5 signed request
func CallAndVerifyMD5(apiObj *api.Sdk, path string, timestamp, sign string, target interface{}) error {
	finalURL, err := url.JoinPath(api.DevNetEndpoint, path)
	if err != nil {
		return fmt.Errorf("failed to join path: %w", err)
	}

	resp, err := Client.R().
		SetHeader("key", apiObj.GetApiKey()).
		SetHeader("timestamp", timestamp).
		SetHeader("sign", sign).
		Post(finalURL)

	if err != nil {
		return fmt.Errorf("request failed: %w", err)
	}

	return ProcessResponse(apiObj, resp.Body(), target, false)
}

// ProcessResponse handles common response parsing and verification
func ProcessResponse(apiObj *api.Sdk, body []byte, target interface{}, verifyRSA bool) error {
	// 1. Parse common response to check code
	var rspCommon response_define.ResponseCommon
	if err := json.Unmarshal(body, &rspCommon); err != nil {
		return fmt.Errorf("failed to unmarshal common response: %w", err)
	}

	if rspCommon.Code != response_define.SUCCESS {
		return fmt.Errorf("API error: code=%s, msg=%s", rspCommon.Code, rspCommon.Msg)
	}

	// 2. Parse into target struct
	if err := json.Unmarshal(body, target); err != nil {
		return fmt.Errorf("failed to unmarshal specific response: %w", err)
	}

	// 3. Verify RSA signature if required
	if verifyRSA {
		mapObj := rsa_utils.ToStringMap(body)
		
		var signHolder struct {
			Sign string `json:"sign"`
		}
		_ = json.Unmarshal(body, &signHolder) // Ignore error as we'll fail verification anyway if missing

		if err := apiObj.VerifyRSAsignature(mapObj, signHolder.Sign); err != nil {
			return fmt.Errorf("RSA signature verification failed: %w", err)
		}
	}

	return nil
}
// CallAPIPath joins the base endpoint with a path
func CallAPIPath(path string) (string, error) {
	return url.JoinPath(api.DevNetEndpoint, path)
}

// VerifyCallback verifies the signature of an incoming callback request
func VerifyCallback(apiObj *api.Sdk, body []byte, sign string) error {
	mapObj := rsa_utils.ToStringMap(body)
	return apiObj.VerifyRSAsignature(mapObj, sign)
}

// VerifyRiskCallback verifies the signature of an incoming risk callback request
func VerifyRiskCallback(apiObj *api.Sdk, body []byte, sign string) error {
	mapObj := rsa_utils.ToStringMap(body)
	return apiObj.VerifyRiskRSAsignature(mapObj, sign)
}
