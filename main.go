package main

import (
	"flag"
)

func main() {
	var (
		name  string
		mode  bool
		typee bool
		num   int
	)
	flag.StringVar(&name, "name", "", "发癫对象")
	flag.BoolVar(&mode, "interactive", false, "是否进入交互模式")
	flag.BoolVar(&typee, "fabing", false, "")
	flag.IntVar(&num, "num", 0, "")
	flag.Parse()

	if mode {
		interactiveMode(typee, num)
	} else {
		normalmode(name, typee, num)
	}
}
