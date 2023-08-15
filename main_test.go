package main

import (
	"os"
	"testing"
)

func DeepDir(path string) ([]string, []string) {
	files, err := os.ReadDir(path)
	if err != nil {
		panic(err)
	}
	var fileNames []string
	var dirNames []string
	for _, file := range files {
		if file.IsDir() {
			dirNames = append(dirNames, file.Name())
			deepFileNames, deepDirNames := DeepDir(path + "/" + file.Name())
			for _, deepFileName := range deepFileNames {
				fileNames = append(fileNames, file.Name()+"/"+deepFileName)
			}
			for _, deepDirName := range deepDirNames {
				dirNames = append(dirNames, file.Name()+"/"+deepDirName)
			}
		} else {
			fileNames = append(fileNames, file.Name())
		}
	}
	return fileNames, dirNames
}

func Test_Main(t *testing.T) {
	//os.Remove("C:/Temp/test")
	//fmt.Println(path.IsAbs("C:/Temp/test"))
	//fmt.Println(path.IsAbs("C:\\Temp\\test"))
	RemoveAntsPool("E:/trash/2/0")
}
