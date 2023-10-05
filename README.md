# fadian-go

[afadian/fadian](https://github.com/afadian/fadian) 的 Golang 版本，按 ~~1:114514~~ 的比例还原了原版。

本工具可以帮助您快速发癫。

## Usage

### Command Line

```bash
$ ./fadian-go --help

USAGE:
   fadian-go [global options] command [command options] [arguments...]

COMMANDS:
   fadian   发癫
   fabing   发病
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --num value    重复次数 (default: 1)
   --help, -h     show help
   --version, -v  print the version

OPTIONS:
   --name value, -n value  对谁发病/发癫
   --interactive, -i       是否交互式发病 (default: false)
   --help, -h              show help
```

If you want to use interactive mode, you can use `-i` or `--interactive` option.

### Debug

You should set `DEBUG` environment variable to `true` to enable debug mode.

```bash
DEBUG=true ./fadian-go
```

## Build

1. Clone this repository: `git clone git@github.com:afadian/fadian.git`
2. Run `cd fadian-go` to enter the project directory.
3. Run `go mod download` to download dependencies.
4. Run `go build -o fadian-go` to build the project.
5. Run `./fadian-go` to run the project.

## License

This project is licensed under the GNU Affero General Public License v3.0. See [LICENSE](LICENSE) for more details.
