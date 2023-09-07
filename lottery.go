package main

import (
	"crypto/rand"
	"flag"
	"log"
	"math/big"
)

func main() {
	nums := flag.Int64("n", 0, "lottery number")
	flag.Parse()
	lottery(*nums)
}

func lottery(nums int64) {
	// 假设楼层为 100，确保 rand 函数输出 0 - 99
	if nums != 0 {
		nums -= 1
	}
	n, err := rand.Int(rand.Reader, big.NewInt(nums))
	if err != nil {
		panic(err)
	}
	// rand 结果加 1 后得出正确楼层
	log.Println(n.Int64() + 1)
}
