package main

import (
  "log"
  "io/ioutil"
  "os"

  "encoding/base64"
  "encoding/json"
)

func CheckErr(err error) {
  if err != nil {
    log.Fatal(err)
    return
  }
}

func WriteToFile(txt []byte, path string) {
  err := ioutil.WriteFile(path, txt, 0644)
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
  pubKey, err := ioutil.ReadFile("keys/ed25519.pub")
  CheckErr(err)

  privKey, err := ioutil.ReadFile("keys/ed25519")
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
