package common

import (
	"golang.org/x/crypto/sha3"
	"golang.org/x/crypto/ed25519"
	"log"
)

type CryptoKeypair struct {
	PublicKey PublicKey
	PrivateKey PrivateKey
}


func HashData(data string) [32]byte {
	return sha3.Sum256([]byte(data))
}

func GenerateKeyPair() CryptoKeypair {
	pub, pri, err := ed25519.GenerateKey(nil)
	if err != nil {
		log.Println(err)
	}
	
	return CryptoKeypair{PrivateKey: pri, PublicKey: pub}
}

type PublicKey ed25519.PublicKey


type PrivateKey ed25519.PrivateKey
