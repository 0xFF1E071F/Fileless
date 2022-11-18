package main

import (
  "fmt"
  "io/ioutil"
  "encoding/base64"

  crypto "crypto/ed25519"
  "crypto/md5"
)

var PubKey crypto.PublicKey
var PrivKey crypto.PrivateKey
var PubKeyBase string
var PrivKeyBase string

func generateKeypairs() {
  /*
    TODO:

    - Optimizing the body of this function
  */
  pubKey, privKey, err := crypto.GenerateKey(nil)
  CheckErr(err)

  PubKey = pubKey
  PrivKey = privKey

  PubKeyBase = base64.StdEncoding.EncodeToString([]byte(PubKey))
  PrivKeyBase = base64.StdEncoding.EncodeToString([]byte(PrivKey))
  // Writing the keypairs to the keys folder

  err = ioutil.WriteFile("keys/ed25519.pub", []byte(PubKeyBase), 0644)
  CheckErr(err)

  err = ioutil.WriteFile("keys/ed25519", []byte(PrivKeyBase), 0644)
  CheckErr(err)
}

func SignCmd(cmd string) string {
  /* 
    TODO:
    
    - Optimizing the body of this function
  */

  fmt.Println("Private key: ", PrivKey)
  hashSlice := md5.Sum([]byte(cmd))
  cmdMD5 := hashSlice[:]
  signedCmd := crypto.Sign(PrivKey, cmdMD5)

  return base64.StdEncoding.EncodeToString(signedCmd) 
}

func InitCrypto() {
  generateKeypairs()
}

