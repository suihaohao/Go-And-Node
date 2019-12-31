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
	
	getDaysOfCurrMonth()

}

func getDaysOfCurrMonth() int {
	t , _ := time.Parse("200601", time.Now().Local().Format("200601"))
	t1 , _ := time.Parse("200601", t.AddDate(0, 1, 0).Local().Format("200601"))
	days := int(t1.Sub(t).Hours() / 24)
	return days
}



