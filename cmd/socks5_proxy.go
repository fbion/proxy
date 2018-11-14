// wiki https://en.wikipedia.org/wiki/SOCKS#SOCKS5
package main

import (
	"log"
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())

}
