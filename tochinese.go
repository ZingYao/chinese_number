package chinese_number

import (
	"fmt"
	"github.com/ZingYao/chinese_number/simplified"
	"github.com/ZingYao/chinese_number/traditional"
	"strconv"
	"strings"
)

func Number2Chinese(langType LangType, number int64) string {
	var zero string
	var toChineseUnitArr, numArr []string
	var tenThousands string
	switch langType {
	case Simplified:
		numArr = simplified.NumArr
		zero = simplified.Zero
		toChineseUnitArr = simplified.ToChineseUintArr
		tenThousands = simplified.TenThousands
	case Traditional:
		numArr = traditional.NumArr
		zero = traditional.Zero
		toChineseUnitArr = traditional.ToChineseUintArr
		tenThousands = traditional.TenThousands
	}
	if number == 0 {
		return zero
	}
	numStr := strconv.FormatInt(number, 10)
	lessThenZero := false
	if strings.Contains(numStr, "-") {
		lessThenZero = true
		numStr = strings.ReplaceAll(numStr, "-", "")
	}
	isBetweenZero := false
	var chineseNumStrArr []string
	arrLen := len(numStr)
	for index, item := range numStr {
		if item != '0' {
			if isBetweenZero {
				isBetweenZero = false
				chineseNumStrArr = append(chineseNumStrArr, fmt.Sprintf("%s%s%s", zero, numArr[item-'0'], toChineseUnitArr[arrLen-1-index]))
			} else {
				chineseNumStrArr = append(chineseNumStrArr, fmt.Sprintf("%s%s", numArr[item-'0'], toChineseUnitArr[arrLen-1-index]))
			}
		} else {
			if toChineseUnitArr[arrLen-1-index] == tenThousands {
				chineseNumStrArr = append(chineseNumStrArr, tenThousands)
			} else {
				isBetweenZero = true
			}
		}
	}
	str := strings.Join(chineseNumStrArr, "")
	if lessThenZero {
		str = "负" + str
	}
	return str
}

func Number2Simplified(number int64) string {
	return Number2Chinese(Simplified, number)
}

func Number2Traditional(number int64) string {
	return Number2Chinese(Traditional, number)
}
