# fadian-go

[afadian/fadian](https://github.com/afadian/fadian) 的 Golang 版本，按 ~~1:114514~~ 的比例还原了原版。

本工具可以帮助您快速发癫。

## Usage

```
Usage of fadian-go:
  -d            是否开启调试模式
  -f            是否进入发病模式
  -i            是否进入交互模式
  -name string  发癫对象
  -num int      发病/发癫次数 (default 1)
```

## Build with source code

1. Clone this repository: `git clone https://github.com/afadian/fadian-go.git`
2. Run `go mod download` to download dependencies
3. Run `go build -o fadian` to build the project
4. Run `./fadian` to run the project
