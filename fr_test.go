package main

import (
	"os"
	"runtime"
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
	//RemoveAntsPool("E:/trash/2/0")

	// get the number of cores

	// this will give the logical core count
	// we know that if the thread number is more than the logical
	//core count, the performance will be bad

	// so it is not a good strategy to create a lot of threads
	// especially when you need to sync them .
	//It may cause the performance to be worse than single thread

	println(runtime.NumCPU())

}
