package util

import (
	"math/rand"
	"time"
)

func Rand(min, max int) int {
	//rand.NewRand(NewSource(seed))
	mtRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	//rand.Seed(time.Now().UnixNano())
	randomNumber := mtRand.Intn(max-min+1) + min
	// 生成一个0到100之间的随机整数
	// randomNumber := rand.Intn(101)
	//randomNumber := rand.Intn(max-min+1) + min
	//fmt.Println(randomNumber)
	return randomNumber
}
