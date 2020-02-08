package algo

import (
	"github.com/smartystreets/goconvey/convey"
	"math/rand"
	"testing"
	"time"
)

func TestSimpleRand(t *testing.T) {
	ForTest("简单随机算法", t, SimpleRand)
}

func TestDoubleRandom(t *testing.T) {
	ForTest("改进先洗牌算法", t, DoubleRandom)
}

func TestDoubleAverage(t *testing.T) {
	ForTest("2倍平均数算法", t, DoubleAverage)
}

func TestBeforeShuffle(t *testing.T) {
	ForTest("先洗牌算法", t, BeforeShuffle)
}

func TestAfterShuffle(t *testing.T) {
	count, amount := int64(10), int64(100) * 100
	rand.Seed(time.Now().UnixNano())
	sum := int64(0)

	list, err := AfterShuffle(count, amount)
	if err != nil {
		panic(err.Error())
	}
	for _, v := range list {
		sum += v
	}

	convey.Convey("后洗牌算法", t, func() {
		convey.So(sum, convey.ShouldEqual, amount)
	})
}

func ForTest(message string, t *testing.T, fn func(int64, int64) (int64, error)){
	count, amount := int64(10), int64(100) * 100
	rand.Seed(time.Now().UnixNano())
	sum := int64(0)
	testAmount := amount
	for i := int64(0); i < count; i++ {
		money, err := fn(count - i, amount)
		if err != nil {
			panic(err.Error())
		}
		amount = amount - money
		sum += money
	}

	convey.Convey(message, t, func() {
		convey.So(sum, convey.ShouldEqual, testAmount)
	})
}