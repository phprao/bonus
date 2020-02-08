package main

import (
	"fmt"
	"github.com/segmentio/ksuid"
)

func main(){
	uid := ksuid.New().Next().String()
	fmt.Println(uid)
}

