package cfmt

// 代码 意义
// -------------------------
//
//	0  终端默认设置
//	1  高亮显示
//	4  使用下划线
//	5  闪烁
//	7  反白显示
//	8  不可见
const (
	CODE_DEFAULT    = 0
	CODE_HIGH_LIGHT = 1
	CODE_UNDER_LINE = 4
	CODE_TWINKLE    = 5
	CODE_ALBEDO     = 7
	CODE_INVISIBLE  = 8
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  黄色
// 34  44  蓝色
// 35  45  紫红色
// 36  46  青蓝色
// 37  47  白色
const (
	BG_BLACK   = 40
	BG_RED     = 41
	BG_GREEN   = 42
	BG_YELLOW  = 43
	BG_BLUE    = 44
	BG_MAGENTA = 45
	BG_INDIGO  = 46
	BG_WHITE   = 47
)

const (
	FG_BLACK   = 30
	FG_RED     = 31
	FG_GREEN   = 32
	FG_YELLOW  = 33
	FG_BLUE    = 34
	FG_MAGENTA = 35
	FG_INDIGO  = 36
	FG_WHITE   = 37
)
