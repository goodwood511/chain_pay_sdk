package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/goodwood511/chain_pay_sdk/example/utils"
	"github.com/goodwood511/chain_pay_sdk/response_define"
	"github.com/goodwood511/chain_pay_sdk/rsa_utils"
	"github.com/sirupsen/logrus"
)

func main() {
	apiObj := utils.InitSDK()

	r := gin.Default()

	r.POST("/withdraw_callback", func(c *gin.Context) {
		body, err := c.GetRawData()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read request body"})
			return
		}

		var req response_define.RequestTokenCb
		if err := json.Unmarshal(body, &req); err != nil {
			logrus.Warnf("Unmarshal failed: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid JSON"})
			return
		}

		if err := utils.VerifyCallback(apiObj, body, req.Sign); err != nil {
			logrus.Warnf("VerifyRSAsignature failed: %v", err)
			c.JSON(http.StatusForbidden, gin.H{"message": "Invalid signature"})
			return
		}

		logrus.Infof("Callback received: %+v", req)

		// Generate response signature
		timestamp := strconv.FormatInt(time.Now().Unix(), 10)
		rsp := response_define.ResponseWithdraw{
			Code:      "0",
			Timestamp: timestamp,
			Message:   "success",
		}

		// Re-marshal to get map for signing (as per original logic)
		jStr, _ := json.Marshal(&req)
		reqMap := rsa_utils.ToStringMap(jStr)
		// Original used rsa_utils.ToStringMap(jStr)
		// Let's stick to what was there but clean
		
		clientSign, err := apiObj.GenerateRSASignature(reqMap)
		if err != nil {
			logrus.Errorf("GenerateRSASignature failed: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal error"})
			return
		}
		rsp.Sign = clientSign

		c.JSON(http.StatusOK, rsp)
	})

	logrus.Info("Starting notify_callback server on :9003")
	r.Run(":9003")
}
