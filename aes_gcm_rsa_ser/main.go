package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"
)

func main() {
	//密文
	cipherText := "AkK5KHb+N9A2+9ExM+KwbaQ4DS2Xmz5a3iwt3K9CfWDIrr/3AAbQ6FTSm7Ol6KyyaSczj4MU7b11d38zpiEP"
	//aesGcm对称加密/解密的Key
	key := "thisis32bitlongkey1234567890abcd"
	//数据rsa签名
	signature := "RgeK/pY0qUikCtzD8mPk0Trs03MccsfMNL9Poc4iWfeUM/zLlL7xw8VYxxNnlfZ8/c9F0cEA012ppNZDBQQ3iq28bnWoDZq0sUR2yVX9D6Jpx8iEuhy3FQywr3l0K0oa2JyiAyAdVwTEPYngnsGJzVtI9IA8l0qPs3BbX0hwx3LMSk1Pj2HKcDtjsRlriDXCPQPIH3df1xiakMO5msIcJhZPRXgG4mQeuweiTUbVj3gSVFHniXocn2PbjYEbeMsMxiIsoH3emJl1jF+ww3wi3Ymofag7XTILzdRcBzi+q0n2gJUHPR1GGpceAZIveMR1HzsVz0UX5PWYoAmUTpWCpg=="
	//加载rsa公钥
	publicKey, err := loadPublicKey("../rsa_key/public.pem")
	if err != nil {
		log.Fatalf("Error loading public key: %v", err)
	}

	err = VerifyAuthRequest(key, cipherText, signature, publicKey)
	if err != nil {
		fmt.Println("Verify failed:", err)
		return
	}
}

// VerifyAuthRequest 验证网关发送的auth request
func VerifyAuthRequest(key string, ciphertext string, signature string, publicKey *rsa.PublicKey) error {

	// 对密文进行对称解密，获取明文
	plainText, err := aesGCMDecrypt(ciphertext, key)
	if err != nil {
		log.Errorf("AGENT_AUTH: decryptAESGCM err: %v", err)
		return err
	}

	// 对签名使用公钥进行验签
	if err := rsaVerifySign(plainText, signature, publicKey); err != nil {
		log.Errorf("AGENT_AUTH: rsaVerifySign err: %v", err)
		return err
	} else {
		fmt.Printf("rsaVerifySign ok")
		fmt.Println("=========cipherText=========", ciphertext)
		fmt.Println("=========originalText=========", plainText)
	}

	return nil
}

// rsaVerifySign rsa解签名并验签
func rsaVerifySign(originalText string, signText string, publicKey *rsa.PublicKey) error {
	sign, err := base64.StdEncoding.DecodeString(signText)
	if err != nil {
		log.Errorf("AGENT_AUTH: decode signData to byte err: %v", err)
		return err
	}

	hash := sha256.New()
	_, err = hash.Write([]byte(originalText))
	if err != nil {
		log.Errorf("AGENT_AUTH: sha256 write err: %v", err)
		return err
	}
	hashText := hash.Sum(nil)

	return rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hashText, sign)
}

// aesGCMDecrypt aes gcm对称解密
func aesGCMDecrypt(cipherText, key string) (string, error) {
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		log.Errorf("AGENT_AUTH: decrpty aes gcm failed to creates new block: %v", err)
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		log.Errorf("AGENT_AUTH: decrpty aes gcm failed to creates new gcm: %v", err)
		return "", err
	}

	cipherData, err := base64.StdEncoding.DecodeString(cipherText)
	if err != nil {
		log.Errorf("AGENT_AUTH: ciphertext decode err: %v", err)
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherData) < nonceSize {
		log.Errorf("AGENT_AUTH: ciphertext is too short")
		return "", errors.New("ciphertext is invalid")
	}

	nonce, cipherContent := cipherData[:nonceSize], cipherData[nonceSize:]

	plainData, err := gcm.Open(nil, nonce, cipherContent, nil)
	if err != nil {
		log.Errorf("AGENT_AUTH: decrypts ciphertext failed: %v", err)
		return "", err
	}

	return string(plainData), nil
}

// 加载 RSA 公钥
func loadPublicKey(filename string) (*rsa.PublicKey, error) {
	pubKeyBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read public key file: %v", err)
	}
	fmt.Println("pubKeyBytes:", string(pubKeyBytes))
	// 解析 PEM 文件
	block, _ := pem.Decode(pubKeyBytes)
	if block == nil {
		return nil, fmt.Errorf("failed to parse PEM block containing the public key")
	}

	pubKey, err := x509.ParsePKCS1PublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse public key: %v", err)
	}

	return pubKey, nil
}
