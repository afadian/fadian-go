package main

import (
	"fmt"
)

func interactiveMode(typee bool, num int) {
	fmt.Println("进入交互模式")
	var name string
	fmt.Println("请输入对方名字：")
	fmt.Scanln(&name)
	if name == "" {
		fmt.Println("请输入对方名字！输入exit退出。")
	}
	if name == "exit" {
		return
	}
	if typee == true {
		//if num == "" {
		fmt.Println(FaBing(name, num))
		//      } else {
		//              fmt.Println(FaBing(name, num))
		//}
	} else {
		//if num == "" {
		fmt.Println(dog(name, num))
		//      } else {
		//              fmt.Println(dog(name, num))
		// }
	}
}
