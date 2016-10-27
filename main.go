package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"io/ioutil"
	"log"
)

var (
	keySize int
	file    string
)

func init() {
	flag.IntVar(&keySize, "b", 1024, "Key size.")
	flag.StringVar(&file, "f", "rsa_id", "output")
	flag.Parse()
}

func main() {
	rsaPriv, err := rsa.GenerateKey(rand.Reader, keySize)
	if err != nil {
		log.Fatal(err)
	}

	privKey := x509.MarshalPKCS1PrivateKey(rsaPriv)
	privPem := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privKey,
	}

	err = ioutil.WriteFile(file, pem.EncodeToMemory(privPem), 0644)
	if err != nil {
		log.Fatal(err)
	}

	pubKey, err := x509.MarshalPKIXPublicKey(&rsaPriv.PublicKey)
	if err != nil {
		log.Fatal(err)
	}

	pubPem := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubKey,
	}

	err = ioutil.WriteFile(file+".pub", pem.EncodeToMemory(pubPem), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
