package algo

import (
	"errors"
	"math/rand"
)

/**
2倍平均数算法
 */

func DoubleAverage(count, amount int64) (int64, error) {
	if count == 1 {
		return amount, nil
	}
	// 计算最大可调度金额
	max := amount - count * MIN_MONEY
	if max == 0 {
		return 0, nil
	}
	if max < 0 {
		return 0, errors.New("红包金额与个数不匹配")
	}
	// 计算均值
	avg := max / count
	// 计算2倍均值
	avg *= 2
	// 以2倍均值作为基数来随机红包
	x := rand.Int63n(avg)

	return x + MIN_MONEY, nil
}