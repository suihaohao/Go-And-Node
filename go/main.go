

import (
	"crypto/sha1"
	"fmt"
  "log"
)


func main(){
  sha1Sum := fmt.Sprintf("%x", sha1.Sum([]byte(tempStr)))
  log.Println(sha1Sum)
}
