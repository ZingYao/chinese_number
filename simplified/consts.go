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
)

var (
	ChineseNum  = map[string]int64{One: 1, Two: 2, Three: 3, Four: 4, Five: 5, Six: 6, Seven: 7, Eight: 8, Nine: 9, Zero: 0}
	ChineseUnit = map[string]int64{Ten: 10, Hundred: 100, Thousands: 1000, TenThousands: 10000, TenMillion: 100000000}
	UnitArr     = []string{TenMillion, TenThousands, Thousands, Hundred, Ten}
)
