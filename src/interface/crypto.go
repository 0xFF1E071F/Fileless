package main

import (
  "fmt"
  "encoding/base64"

  crypto "crypto/ed25519"
  "crypto/md5"
)

var (
  PubKey      crypto.PublicKey
  PrivKey     crypto.PrivateKey
)

func generateKeypairs() {
  /*
    TODO:

    - Optimizing the body of this function
  */
  
  if FileExists("keys/ed25519") && FileExists("keys/ed25519.pub") {
    ReadKeys()
    return
  }

  pubKey, privKey, err := crypto.GenerateKey(nil)
  CheckErr(err)

  PubKey = pubKey
  PrivKey = privKey

  // Writing the keypairs to the keys folder
  WriteToFile([]byte(base64.StdEncoding.EncodeToString([]byte(PubKey))), "keys/ed25519.pub")
  WriteToFile([]byte(base64.StdEncoding.EncodeToString([]byte(PrivKey))), "keys/ed25519")
}

func SignCmd(cmd string) string {
  hashSlice := md5.Sum([]byte(cmd))
  cmdMD5 := hashSlice[:] // Converting from [size]byte to []byte
  signedCmd := crypto.Sign(PrivKey, cmdMD5)

  return base64.StdEncoding.EncodeToString(signedCmd) 
}

func InitCrypto() {
  generateKeypairs()
}

func main() {
  InitCrypto()
  fmt.Println(SignCmd("test123"))
}
