package test

import (
	"fmt"
	"github.com/ZingYao/chinese_number"
	"math"
	"os"
	"sync"
	"testing"
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
	limit := make(chan int, 64)
	wg := sync.WaitGroup{}
	for i := int64(math.MinInt64); i < math.MaxInt64; i++ {
		index++
		wg.Add(1)
		limit <- 1
		go func(number int64) {
			defer func() {
				wg.Done()
				<-limit
			}()
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
		if index%10000000 == 0 {
			fmt.Printf("%.010f%%\n", float64(index)/math.MaxInt64*100*2)
		}
	}
	wg.Wait()
	fmt.Println("done")
}
