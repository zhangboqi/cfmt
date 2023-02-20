package cfmt

import (
	"fmt"
)

type Cfmt struct {
	Code       int
	Background int
	Foreground int
}

func New(code int, background int, foreground int) *Cfmt {
	return &Cfmt{
		Code:       code,
		Background: background,
		Foreground: foreground,
	}
}

func (f *Cfmt) Println(a ...any) {
	s := fmt.Sprintln(a...)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1b, f.Code, f.Background, f.Foreground, s, 0x1b)
}

func (f *Cfmt) Print(a ...any) {
	s := fmt.Sprint(a...)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1b, f.Code, f.Background, f.Foreground, s, 0x1b)
}

func (f *Cfmt) Printf(format string, a ...any) {
	s := fmt.Sprintf(format, a...)
	fmt.Printf("%c[%d;%d;%dm%s%c[0m", 0x1b, f.Code, f.Background, f.Foreground, s, 0x1b)
}

var gfmt = New(CODE_HIGH_LIGHT, BG_BLACK, FG_GREEN)

func Gprintln(a ...any) {
	gfmt.Println(a...)
}
func Gprint(a ...any) {
	gfmt.Print(a...)
}
func Gprintf(format string, a ...any) {
	gfmt.Printf(format, a...)
}

var rfmt = New(CODE_HIGH_LIGHT, BG_BLACK, FG_RED)

func Rprintln(a ...any) {
	rfmt.Println(a...)
}
func Rprint(a ...any) {
	rfmt.Print(a...)
}
func Rprintf(format string, a ...any) {
	rfmt.Printf(format, a...)
}

var yfmt = New(CODE_HIGH_LIGHT, BG_BLACK, FG_YELLOW)

func Yprintln(a ...any) {
	yfmt.Println(a...)
}
func Yprint(a ...any) {
	yfmt.Print(a...)
}
func Yprintf(format string, a ...any) {
	yfmt.Printf(format, a...)
}
