/*
  TODO:
  
  - Implementing asymmetric encryption (ed25519)
  - Documentation: https://pkg.go.dev/crypto/ed25519
*/

package main

import (
  "io/ioutil"
  crypto "crypto/ed25519"
  "encoding/base64"
)

func generateKeypairs() {
  pubKey, privKey, err := crypto.GenerateKey(nil)
  CheckErr(err)

  pubKeyBase := base64.StdEncoding.EncodeToString([]byte(pubKey))
  privKeyBase := base64.StdEncoding.EncodeToString([]byte(privKey))
  // Writing the keypairs to the keys folder

  err = ioutil.WriteFile("keys/ed25519.pub", []byte(pubKeyBase), 0644)
  CheckErr(err)

  err = ioutil.WriteFile("keys/ed25519", []byte(privKeyBase), 0644)
  CheckErr(err)
}

func InitCrypto() {
  generateKeypairs()
}

func main() {
  InitCrypto()
}
