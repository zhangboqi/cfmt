# cfmt

![Apache license](https://img.shields.io/badge/license-Apache%202-blue)
![Go Version](https://img.shields.io/badge/go%20version-%3E%3D1.16-brightgreen)

[README](README.md) | [中文文档](README_zh.md)

一款轻量级的在终端上输出彩色文本的封装库

## 安装
```
go get github.com/zhangboqi/cfmt
```

## 使用
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

## 自定义方式
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