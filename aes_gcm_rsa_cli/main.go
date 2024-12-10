package main

import (
	"crypto"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	log "github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
)

func main() {
	//原始数据
	originalText := "Hello, this is the data to encrypt!"
	//aesGcm对称加密/解密的Key
	key := "thisis32bitlongkey1234567890abcd"
	// 加载rsa签名的私钥
	privateKey, err := loadPrivateKey("../rsa_key/private.pem")
	if err != nil {
		log.Fatalf("Error loading private key: %v", err)
	}
	//通过私钥对originalText进行签名
	signText, err := rsaGenerateSign(originalText, privateKey)
	if err != nil {
		log.Fatalf("Error generating sign: %v", err)
	}
	//通过aesGcm对originalText进行加密
	cipherText, err := aesGCMEncrypts(originalText, key)
	if err != nil {
		log.Fatalf("Error encrypting data: %v", err)
		return
	}
	fmt.Println("=======originalText========", originalText)
	fmt.Println("========cipherText=========", cipherText)
	fmt.Println("============key============", key)
	fmt.Println("=========signature=========", signText)
	privKeyBytes, err := ioutil.ReadFile("../rsa_key/private.pem")
	if err != nil {
		fmt.Errorf("failed to read private key file: %v", err)
	}
	fmt.Println("========privKey========", string(privKeyBytes))
}

// aesGCMEncrypts aes gcm对称加密
func aesGCMEncrypts(plaintext, key string) (string, error) {
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

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	cipherContent := gcm.Seal(nonce, nonce, []byte(plaintext), nil)
	cipherText := base64.StdEncoding.EncodeToString(cipherContent)

	return cipherText, nil
}

// rsaGenerateSign rsa签名
func rsaGenerateSign(originalText string, privateKey *rsa.PrivateKey) (string, error) {

	hash := sha256.New()
	_, err := hash.Write([]byte(originalText))
	if err != nil {
		log.Errorf("AGENT_AUTH: sha256 write err: %v", err)
		return "", err
	}
	hashText := hash.Sum(nil)

	sign, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashText)
	if err != nil {
		log.Errorf("AGENT_AUTH: rsa SignPKCS1v15 err: %v", err)
		return "", err
	}

	signText := base64.StdEncoding.EncodeToString(sign)

	return signText, nil
}

func loadPrivateKey(filePath string) (*rsa.PrivateKey, error) {
	// 读取私钥文件
	privKeyBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Errorf("failed to read private key file: %v", err)
		return nil, err
	}

	// 解码PEM格式
	block, _ := pem.Decode(privKeyBytes)
	if block == nil {
		fmt.Errorf("failed to parse PEM block containing the private key")
		return nil, err
	}

	// 解析私钥
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		fmt.Errorf("failed to parse RSA private key: %v", err)
		return nil, err
	}

	return privKey, nil
}
