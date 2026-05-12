package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"

	"github.com/goodwood511/chain_pay_sdk/rsa_utils"
	"github.com/sirupsen/logrus"
)

func verifySignature(publicKey *rsa.PublicKey, data, signature string) error {
	decodedSignature, err := base64.URLEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("error decoding signature: %v", err)
	}
	hash := sha256.New()
	io.WriteString(hash, data)
	hashInBytes := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashInBytes[:], decodedSignature)
	if err != nil {
		return fmt.Errorf("signature verification failed: %v", err)
	}

	return nil
}

func signData(privateKey *rsa.PrivateKey, data string) (string, error) {
	hash := sha256.New()
	io.WriteString(hash, data)
	hashInBytes := hash.Sum(nil)

	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashInBytes[:])
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(signature), nil
}

func main() {
	rawStr := "aaaaa"
	rsaPrivateKeyB64 := "YOUR_PRIVATE_KEY_BASE64" // Replace with real key for testing

	privateKey, err := rsa_utils.LoadPrivateKeyFromBase64(rsaPrivateKeyB64)
	if err != nil {
		logrus.Errorf("Failed to load private key: %v", err)
		return
	}
	publicKey := &privateKey.PublicKey

	// 1. Sign data
	goSign, err := signData(privateKey, rawStr)
	if err != nil {
		logrus.Errorf("Sign data failed: %v", err)
	} else {
		logrus.Infof("Sign data success: %s", goSign)
	}

	// 2. Verify signature
	err = verifySignature(publicKey, rawStr, goSign)
	if err != nil {
		logrus.Errorf("Verify signature failed: %v", err)
	} else {
		logrus.Info("Verify signature success")
	}
}
