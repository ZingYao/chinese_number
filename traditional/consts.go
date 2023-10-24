package traditional

// 大写数字定义
const (
	Zero  = "零"
	One   = "壹"
	Two   = "贰"
	Three = "叁"
	Four  = "肆"
	Five  = "伍"
	Six   = "陆"
	Seven = "柒"
	Eight = "捌"
	Nine  = "玖"
)

// 大写单位定义
const (
	Ten          = "拾"
	Hundred      = "佰"
	Thousands    = "仟"
	TenThousands = "万"
	TenMillion   = "亿"
)

var (
	NumArr           = []string{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine}
	ChineseNum       = map[string]int64{One: 1, Two: 2, Three: 3, Four: 4, Five: 5, Six: 6, Seven: 7, Eight: 8, Nine: 9, Zero: 0}
	ChineseUnit      = map[string]int64{Ten: 10, Hundred: 100, Thousands: 1000, TenThousands: 10000, TenMillion: 100000000}
	UnitArr          = []string{TenMillion, TenThousands, Thousands, Hundred, Ten}
	ToChineseUintArr = []string{"", Ten, Hundred, Thousands, TenThousands, Ten, Hundred, Thousands, TenMillion, Ten, Hundred, Thousands, TenThousands, Ten, Hundred, Thousands, TenMillion, Ten, Hundred}
)
