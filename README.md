# cfmt

![Apache license](https://img.shields.io/badge/license-Apache%202-blue)
![Go Version](https://img.shields.io/badge/go%20version-%3E%3D1.16-brightgreen)

[README](README.md) | [中文文档](README_zh.md)

A lightweight packaging library for outputting color text on the terminal

## Install
```
go get github.com/zhangboqi/cfmt
```

## Usage
```golang
package main
import "github.com/zhangboqi/cfmt"

func main(){
	// green
	cfmt.Gprintln("green")

	// red
	cfmt.Rprintln("red")

	// yellow
	cfmt.Yprintln("yellow")

}
```

## Custom
```golang

package main

import (
	"github.com/zhangboqi/cfmt"
)

func main() {
	bfmt := cfmt.New(cfmt.CODE_DEFAULT, cfmt.BG_BLACK, cfmt.FG_BLUE)
	bfmt.Println("Blue")
}

```