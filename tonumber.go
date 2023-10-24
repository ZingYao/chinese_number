package chinese_number

import (
	"fmt"
	"github.com/ZingYao/chinese_number/simplified"
	"github.com/ZingYao/chinese_number/traditional"
	"github.com/gobwas/glob/util/runes"
	"strings"
)

func Chinese2Number(langType LangType, str string) (int64, error) {
	var zero string
	var unitArr []string
	var chineseUnit map[string]int64
	var chineseNum map[string]int64
	switch langType {
	case Simplified:
		zero = simplified.Zero
		unitArr = simplified.UnitArr
		chineseUnit = simplified.ChineseUnit
		chineseNum = simplified.ChineseNum
	case Traditional:
		zero = traditional.Zero
		unitArr = traditional.UnitArr
		chineseUnit = traditional.ChineseUnit
		chineseNum = traditional.ChineseNum
	}
	baseNum := int64(1)
	if strings.Contains(str, "负") {
		baseNum = -1
		str = strings.ReplaceAll(str, "负", "")
	}
	num := int64(0)
	tempStr := strings.ReplaceAll(str, zero, "")
	tempNumStr := []rune(tempStr)
	oldUnit := ""
	for i := 0; i < len(unitArr); i++ {
		unit := unitArr[i]
		if position := runes.IndexAny(tempNumStr, []rune(unit)); position != -1 {
			tempStr := string(tempNumStr[:position])
			tempNumStr = tempNumStr[position+1:]
			n, err := Chinese2Number(langType, tempStr)
			if err != nil {
				return 0, err
			}
			if oldUnit == unit {
				num = num * chineseUnit[unit]
			}
			num += n * chineseUnit[unit]
			oldUnit = unit
			i -= 1
		}
	}
	if len(tempNumStr) == 1 {
		if n, exists := chineseNum[string(tempNumStr)]; exists {
			num += n
		} else {
			s := string(tempNumStr)
			fmt.Println(s)
		}
	} else if len(tempNumStr) > 1 {
		return 0, fmt.Errorf("str:%s is not a simplified chinese number", str)
	}
	num *= baseNum
	return num, nil
}

func Simplified2Number(str string) (int64, error) {
	return Chinese2Number(Simplified, str)
}

func Traditional2Number(str string) (int64, error) {
	return Chinese2Number(Traditional, str)
}
