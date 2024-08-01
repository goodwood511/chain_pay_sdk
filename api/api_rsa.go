package api

import (
	"github.com/goodwood511/chain_pay_sdk/rsa_utils"
)

func (s *Sdk) GenerateRSASignature(mapData map[string]string) (string, error) {

	rawStr := rsa_utils.ComposeParams(mapData)

	privateKey, err := rsa_utils.LoadPrivateKeyFromBase64(s.config.RsaPrivateKey)
	if err != nil {
		return "", err
	}

	return rsa_utils.SignData(privateKey, rawStr)
}

func (s *Sdk) VerifyRSAsignature(mapData map[string]string, sign string) error {

	rawStr := rsa_utils.ComposeParams(mapData)

	publicKey, err := rsa_utils.ParsePublicKey(s.config.PlatformPubKey)
	if err != nil {
		return err
	}

	err = rsa_utils.VerifySignature(publicKey, rawStr, sign)
	if err != nil {
		return err
	}

	return nil
}
