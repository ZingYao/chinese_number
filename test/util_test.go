package test

import (
	"fmt"
	"github.com/ZingYao/chinese_number"
	"math"
	"testing"
)

func TestStr2Num(t *testing.T) {
	strNum := "负九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零八"
	num, err := chinese_number.Simplified2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	}
	if num != math.MinInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	}
	fmt.Printf("simplified: %s number is: %d\n", strNum, num)

	strNum = "九百二十二亿三千三百七十二万零三百六十八亿五千四百七十七万五千八百零七"
	num, err = chinese_number.Simplified2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	}
	if num != math.MaxInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	}
	fmt.Printf("simplified: %s number is: %d\n", strNum, num)

	strNum = "负玖佰贰拾贰亿叁仟叁佰柒拾贰万零叁佰陆拾捌亿伍仟肆佰柒拾柒万伍仟捌佰零捌"
	num, err = chinese_number.Traditional2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	}
	if num != math.MinInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	}
	fmt.Printf("traditional: %s number is: %d\n", strNum, num)

	strNum = "玖佰贰拾贰亿叁仟叁佰柒拾贰万零叁佰陆拾捌亿伍仟肆佰柒拾柒万伍仟捌佰零柒"
	num, err = chinese_number.Traditional2Number(strNum)
	if err != nil {
		t.Errorf("has an error:%+v", err)
		t.Fail()
	}
	if num != math.MaxInt64 {
		t.Errorf("result is not correct")
		t.Fail()
	}
	fmt.Printf("traditional: %s number is: %d\n", strNum, num)
}
