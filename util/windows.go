package util

import "fmt"

func WelcomeWindow()  {
	fmt.Println("****************************")
	fmt.Println("***欢迎使用智能表格小程序***")
	fmt.Println("*********版本:0.0.1*********")
	fmt.Println("****************************")
}

func FunctionWindow() (int, error){
	var funcNum int
	fmt.Println("请输入序号来选择你要使用的功能：")
	fmt.Println("1.抓取相同产品的数量")
	fmt.Println("......")
	fmt.Println("剩余功能待开发中，敬请期待。。。")
	fmt.Println("温馨提示：每次输入完成后请按回车键")
	_, err := fmt.Scanln(&funcNum)
	if err != nil {
		return 0, err
	}
	return funcNum, nil
}