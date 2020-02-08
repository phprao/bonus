package main

import (
	"bonus/infra/algo"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	count, amount := int64(10), int64(100) * 100
	rand.Seed(time.Now().UnixNano())

	for i := int64(0); i < count; i++ {
		money, err := algo.SimpleRand(count - i, amount)
		if err != nil {
			panic(err.Error())
		}
		amount = amount - money
		fmt.Println(money)
	}
}
