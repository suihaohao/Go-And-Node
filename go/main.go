package main

import (
	"crypto/sha1"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
  "log"
)


func main(){
	
  text := "123456text"
  sha1Sum := fmt.Sprintf("%x", sha1.Sum([]byte(text)))
  log.Println(sha1Sum)
	
	
	
	//加盐哈希
	salt := "e9a51074898fb36d"
	h := hmac.New(sha256.New, []byte(salt))
	h.Write([]byte(text))
	hmacResult := hex.EncodeToString(h.Sum(nil))
	log.Println(hmacResult)
	

}




