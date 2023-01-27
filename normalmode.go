package main

import (
	"fmt"
)

func normalmode(name string, typee bool, num int) {
	if typee == true {
		//if num == "" {
		fmt.Println(FaBing(name, num))
		//      } else {
		//              fmt.Println(FaBing(name, num))
		//}
		return
	} else {
		//if num == "" {
		fmt.Println(dog(name, num))
		//      } else {
		//              fmt.Println(dog(name, num))
		// }
		return
	}

}
