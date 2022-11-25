package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"log"
	"math/big"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatal(err)
		return
	}
}

func WriteToFile(txt []byte, path string) {
	err := os.WriteFile(path, txt, 0644)
	CheckErr(err)
}

func FileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func ReadKeys() {
	pubKey, err := os.ReadFile("keys/ed25519.pub")
	CheckErr(err)

	privKey, err := os.ReadFile("keys/ed25519")
	CheckErr(err)

	pubKey, err = base64.StdEncoding.DecodeString(string(pubKey))
	CheckErr(err)

	privKey, err = base64.StdEncoding.DecodeString(string(privKey))
	CheckErr(err)

	PubKey = pubKey
	PrivKey = privKey
}

func StructToJson(cmd AgentCommand) []byte {
	res, err := json.Marshal(cmd)
	CheckErr(err)

	return res
}

func GenerateRandomString(n int) string {
	const letters = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz-"
	ret := make([]byte, n)
	for i := 0; i < n; i++ {
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
		CheckErr(err)

		ret[i] = letters[num.Int64()]
	}

	return string(ret)
}
