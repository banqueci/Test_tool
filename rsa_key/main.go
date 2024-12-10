package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
)

// 生成RSA公私钥对
func generateRSAKeyPair(bits int) (*rsa.PrivateKey, *rsa.PublicKey, error) {
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		return nil, nil, err
	}
	publicKey := &privateKey.PublicKey
	return privateKey, publicKey, nil
}

// 将RSA私钥转换为PEM格式
func privateKeyToPEM(privateKey *rsa.PrivateKey) ([]byte, error) {
	privBytes := x509.MarshalPKCS1PrivateKey(privateKey)
	privBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privBytes,
	}
	return pem.EncodeToMemory(privBlock), nil
}

// 将RSA公钥转换为PEM格式
func publicKeyToPEM(publicKey *rsa.PublicKey) ([]byte, error) {
	pubBytes := x509.MarshalPKCS1PublicKey(publicKey)
	pubBlock := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubBytes,
	}
	return pem.EncodeToMemory(pubBlock), nil
}

// 使用RSA公钥加密
func encryptWithRSA(publicKey *rsa.PublicKey, data []byte) ([]byte, error) {
	return rsa.EncryptOAEP(sha256.New(), rand.Reader, publicKey, data, nil)
}

// 使用RSA私钥解密
func decryptWithRSA(privateKey *rsa.PrivateKey, ciphertext []byte) ([]byte, error) {
	return rsa.DecryptOAEP(sha256.New(), rand.Reader, privateKey, ciphertext, nil)
}

func main() {
	// 生成RSA公私钥对
	privateKey, publicKey, err := generateRSAKeyPair(2048)
	if err != nil {
		log.Fatal(err)
	}

	// 打印生成的公钥和私钥
	privPEM, err := privateKeyToPEM(privateKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Private Key:")
	fmt.Println(string(privPEM))

	pubPEM, err := publicKeyToPEM(publicKey)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Public Key:")
	fmt.Println(string(pubPEM))

	// 要加密的消息
	message := []byte("This is a secret message!")

	// 使用RSA公钥加密消息
	ciphertext, err := encryptWithRSA(publicKey, message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nEncrypted Message: %x\n", ciphertext)

	// 使用RSA私钥解密消息
	plaintext, err := decryptWithRSA(privateKey, ciphertext)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nDecrypted Message: %s\n", plaintext)
}
