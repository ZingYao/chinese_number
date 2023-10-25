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
	const thread = 128
	index := uint64(0)
	wg := sync.WaitGroup{}
	maxNum := uint64(math.MaxUint64)
	length := int64(maxNum / 128)
	for i := int64(1); i <= thread; i++ {
		wg.Add(1)
		go func(startNum int64, endNum int64) {
			defer func() {
				wg.Done()
			}()
			source := rand.NewSource(time.Now().UnixNano())
			r := rand.New(source)
			for i := startNum; i < endNum; i++ {
				index++
				// 生成1到100000000的随机数
				n := r.Int63n(100000000) + 1
				if n != 1 {
					continue
				}
				fmt.Printf("[%s] %.010f%%\tnum:%d\n", time.Now().Format("2006-01-02 15:04:05"), float64(index)/math.MaxInt64*100*2, i)
				str := chinese_number.Number2Simplified(i)
				num, err := chinese_number.Simplified2Number(str)
				if err != nil {
					fmt.Println("err: ", str, i, err)
					os.Exit(-1)
				} else if num != i {
					fmt.Println("err: ", str, i, num)
					os.Exit(-1)
				}
			}
		}(math.MinInt64+length*(i-1), math.MinInt64+i*length)
		time.Sleep(time.Duration(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(1200)) * time.Millisecond)
	}
	wg.Wait()
	fmt.Println("done")
}
