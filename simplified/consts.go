package simplified

// 大写数字定义
const (
	Zero  = "零"
	One   = "一"
	Two   = "二"
	Three = "三"
	Four  = "四"
	Five  = "五"
	Six   = "六"
	Seven = "七"
	Eight = "八"
	Nine  = "九"
)

// 大写单位定义
const (
	Ten          = "十"
	Hundred      = "百"
	Thousands    = "千"
	TenThousands = "万"
	TenMillion   = "亿"
	Trillion     = "兆"
	Quadrillion  = "京"
)

// 常用数组定义
var (
	// NumArr 数字转字符串数组
	NumArr = []string{Zero, One, Two, Three, Four, Five, Six, Seven, Eight, Nine}
	// ChineseNum 字符串转数字map
	ChineseNum = map[string]int64{One: 1, Two: 2, Three: 3, Four: 4, Five: 5, Six: 6, Seven: 7, Eight: 8, Nine: 9, Zero: 0}
	// ChineseUnit 单位转数字map
	ChineseUnit = map[string]int64{Ten: 10, Hundred: 100, Thousands: 1000, TenThousands: 10000, TenMillion: 100000000, Trillion: 1000000000000, Quadrillion: 10000000000000000}
	// UnitArr 单位排序数组
	UnitArr = []string{Trillion, TenMillion, TenThousands, Thousands, Hundred, Ten}
	// ToChineseUintArr 转中文单位数组
	ToChineseUintArr = []string{"", Ten, Hundred, Thousands, TenThousands, Ten, Hundred, Thousands, TenMillion, Ten, Hundred, Thousands, TenThousands, Ten, Hundred, Thousands, TenMillion, Ten, Hundred}
)
