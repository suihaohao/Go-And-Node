

import (
	"crypto/sha1"
	"fmt"
  	"log"
)


func main(){
	text := "123456text"
  sha1Sum := fmt.Sprintf("%x", sha1.Sum([]byte(text)))
  log.Println(sha1Sum)
}
