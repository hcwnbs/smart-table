package main

import (
	"fmt"
	"learnfast/function"
	"learnfast/util"
)

func main() {
	util.WelcomeWindow()
	for {
		num, err := util.FunctionWindow()
		if err != nil {
			return
		}
		switch num {
		case 1:
			err = function.SetNumByName()
			if err != nil {
				fmt.Println(err)
				continue
			}
		default:
			fmt.Println("您输入的功能不存在或数字有误，请重新输入！")
			continue
		}
	}
}
