package cfmt

import (
	"testing"
)

func TestPrintln(t *testing.T) {
	cfmt := &Cfmt{
		Code:       CODE_HIGH_LIGHT,
		Background: BG_BLACK,
		Foreground: FG_GREEN,
	}
	cfmt.Println("ccc", "ddd")
}

func TestPrint(t *testing.T) {
	cfmt := &Cfmt{
		Code:       0,
		Background: 40,
		Foreground: 32,
	}
	cfmt.Print("ccc", "ddd")
}

func TestPrintf(t *testing.T) {
	cfmt := &Cfmt{
		Code:       0,
		Background: 40,
		Foreground: 32,
	}
	cfmt.Printf("%s %d", "ddd", 2)
}

func TestGPrint(t *testing.T) {
	Gprint("green")
}

func TestRPrint(t *testing.T) {
	Rprint("green")
}

func TestYPrint(t *testing.T) {
	Yprint("yellow")
}
