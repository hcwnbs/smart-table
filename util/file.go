package util

import "os"

func PathExists(path string) (bool, error) {
	s, err := os.Stat(path)
	if err == nil && !s.IsDir(){
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}


