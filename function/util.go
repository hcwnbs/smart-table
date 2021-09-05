package function

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strconv"
)

// SetMap set the number of the same name in the map
func SetMap(filePath string, sheetName string,nameRowIndex string, numRowIndex string, results map[string]string)error{
	xFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return err
	}
	nameIndex, _ := strconv.Atoi(nameRowIndex)
	nameIndex = nameIndex-1
	numIndex, _ := strconv.Atoi(numRowIndex)
	numIndex = numIndex-1
	for _, sheet := range xFile.Sheets{
		if sheet.Name == sheetName {
			for name, num := range results{
				for _, row := range xFile.Sheets[0].Rows{
					if len(row.Cells) < 1{
						break
					}
					if row.Cells[nameIndex].String() == name{
						getNum, err := strconv.Atoi(num)
						if err != nil{
							continue
						}
						row.Cells[numIndex].SetValue(getNum)
					}
				}
			}
			err = xFile.Save(filePath)
			if err != nil {
				return err
			}
			return nil
		}
	}
	return fmt.Errorf("not found the sheet")
}

// GetMap get the map of name and number like map[name]=num
func GetMap (filePath string, sheetName string,nameRowIndex string, numRowIndex string)(map[string]string, error){
	xFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		return nil, err
	}
	nameIndex, _ := strconv.Atoi(nameRowIndex)
	nameIndex = nameIndex-1
	numIndex, _ := strconv.Atoi(numRowIndex)
	numIndex = numIndex-1
	results := make(map[string]string)
	for _, sheet := range xFile.Sheets{
		if sheet.Name == sheetName {
			for _, row := range sheet.Rows{
				if len(row.Cells) < 1{
					break
				}
				name := row.Cells[nameIndex].String()
				num := row.Cells[numIndex].String()
				if name == "" || num == "" {
					continue
				}
				results[name] = num
			}
			return results, nil
		}
	}
	return nil, fmt.Errorf("not found the sheet")
}
