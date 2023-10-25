package test

import (
	"fmt"
	"github.com/ZingYao/chinese_number"
	"math"
	"math/rand"
	"os"
	"sync"
	"testing"
	"time"
)

func TestStr2Num(t *testing.T) {
	strNum := "负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八"
	num, err := chinese_number.Simplified2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	} else if num != math.MinInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	} else {
		fmt.Printf("simplified: %s number is: %d\n", strNum, num)
	}

	strNum = "九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零七"
	num, err = chinese_number.Simplified2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	} else if num != math.MaxInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	} else {
		fmt.Printf("simplified: %s number is: %d\n", strNum, num)
	}

	strNum = "负玖佰贰拾贰亿叁仟叁佰柒拾贰万零叁佰陆拾捌亿伍仟肆佰柒拾柒万伍仟捌佰零捌"
	num, err = chinese_number.Traditional2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	} else if num != math.MinInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	} else {
		fmt.Printf("traditional: %s number is: %d\n", strNum, num)
	}

	strNum = "玖佰贰拾贰亿叁仟叁佰柒拾贰万零叁佰陆拾捌亿伍仟肆佰柒拾柒万伍仟捌佰零柒"
	num, err = chinese_number.Traditional2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	} else if num != math.MaxInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	} else {
		fmt.Printf("traditional: %s number is: %d\n", strNum, num)
	}
}

func TestChinese2NumberUnit(t *testing.T) {
	strNum := "负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十万零九千九百九十九"
	num, err := chinese_number.Chinese2Number(chinese_number.Simplified, strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	} else if num != -9223372036854709999 {
		t.Errorf("result is not correct")
		t.Fail()
	} else {
		fmt.Printf("simplified: %s number is: %d\n", strNum, num)
	}
}

func TestNumber2Chinese(t *testing.T) {
	num := int64(math.MinInt64)
	str := chinese_number.Number2Simplified(num)
	if str != "负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八" {
		t.Errorf("result is not correct")
		t.Fail()
	} else {
		fmt.Printf("number: %d simplified is: %s\n", num, str)
	}
}

func TestAllNum(t *testing.T) {
	index := uint64(0)
	limit := make(chan int, 128)
	wg := sync.WaitGroup{}
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := int64(math.MinInt64); i < math.MaxInt64; i++ {
		index++
		// 生成1到100000000的随机数
		n := r.Intn(100000000) + 1
		if n == 1 {
			wg.Add(1)
			limit <- 1
			go func(number int64) {
				defer func() {
					wg.Done()
					<-limit
				}()
				fmt.Printf("[%s] %.010f%%\tnum:%d\n", time.Now().Format("2006-01-02 15:04:05"), float64(index)/math.MaxInt64*100*2, number)
				str := chinese_number.Number2Simplified(number)
				num, err := chinese_number.Simplified2Number(str)
				if err != nil {
					fmt.Println("err: ", str, number, err)
					os.Exit(-1)
				} else if num != number {
					fmt.Println("err: ", str, number, num)
					os.Exit(-1)
				}
			}(i)
		}
	}
	wg.Wait()
	fmt.Println("done")
}
