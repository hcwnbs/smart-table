package function

import (
	"fmt"
	"learnfast/util"
	"time"
)

// SetNumByName set the number by the same name
func SetNumByName() error{
	var (
		filePath string

		oldSheet string
		oldNameRowIndex string
		oldNumRowIndex string

		newSheet string
		newNameRowIndex string
		newNumRowIndex string
	)
	fmt.Println("请输入你要修改的xlsx文件路径：")
	fmt.Println("例如：D:/test.xlsx, 如果你的文件和程序在同一目录也可以是：./test.xlsx")
	_, err := fmt.Scanln(&filePath)
	if err != nil {
		return err
	}
	isExit, _ := util.PathExists(filePath)
	if !isExit {
		fmt.Println("该文件找不到或你输入的是个文件夹路径！请重新进入该功能")
		return err
	}

	fmt.Println("请输入你要获取的sheet名字：如Sheet1")
	_, err = fmt.Scanln(&oldSheet)
	if err != nil {
		return err
	}
	fmt.Println("请输入你要获取的产品名字所在列：如A列就是第1列，输入1")
	_, err = fmt.Scanln(&oldNameRowIndex)
	if err != nil {
		return err
	}
	fmt.Println("请输入你要获取的产品数量所在列：如A列就是第1列，输入1")
	_, err = fmt.Scanln(&oldNumRowIndex)
	if err != nil {
		return err
	}
	result,err := GetMap(filePath, oldSheet, oldNameRowIndex, oldNumRowIndex)
	if err != nil {
		fmt.Println("获取失败，请重新进入该功能")
		return err
	}

	fmt.Println("请输入你要保存的sheet名字：如Sheet1")
	_, err = fmt.Scanln(&newSheet)
	if err != nil {
		return err
	}
	fmt.Println("请输入你要保存的产品名字所在列：如A列就是第1列，输入1")
	_, err = fmt.Scanln(&newNameRowIndex)
	if err != nil {
		return err
	}
	fmt.Println("请输入你要保存的产品数量所在列：如A列就是第1列，输入1")
	_, err = fmt.Scanln(&newNumRowIndex)
	if err != nil {
		return err
	}
	for i := 10; i > 0; i-- {
		str := fmt.Sprintf("还有%d秒开始操作文件，请先将该文件关闭，以免保存失败\n", i)
		time.Sleep(time.Second)
		fmt.Println(str)
	}

	err = SetMap(filePath, newSheet, newNameRowIndex, newNumRowIndex, result)
	if err != nil {
		fmt.Println("保存失败，请重新进入该功能！")
		return err
	}

	fmt.Println("保存成功,请重新打开文件确认！")
	return nil
}
