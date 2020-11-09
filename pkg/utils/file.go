package utils

import "os"

func PathExists(path string) (bool, error) {
	_, e := os.Stat(path)
	if e == nil {
		return true, nil
	}
	if os.IsNotExist(e) {
		return false, nil
	}
	return false, e
}
func CreateFile(fileName string) (f *os.File, e error) {
	exist, e := PathExists(fileName)
	if e != nil {
		return
	}
	if exist == false {
		f, e := os.Create(fileName)
		if e != nil {
			return f, e
		}
	}
	f, e = os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0777)
	if e != nil {
		return
	}
	return
}
func LogDir(savePath string) string {
	return savePath + "/" + GetCurrentDate() + "/"
}
