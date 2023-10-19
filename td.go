package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

var (
	rootPath = "G:/trash"
	sb       = strings.Builder{}
	runtime  = 10 * 10000
)

func main() {
	expr, err := strconv.Atoi(os.Args[1])
	if err != nil {
		expr = 0
	}
	switch expr {
	case 0:
		TestDataType0()
		time.Sleep(1 * time.Second)
	case 1:
		TestDataType1()
	case 2:
		TestDataType2()
	default:
		fmt.Println("需要有效的参数")
	}
}

func TestDataType0() {
	// 创建目录
	fmt.Println("创建测试数据 类型为单文件夹 " + strconv.Itoa(runtime) + "文件")
	os.Mkdir(filepath.Join(rootPath, "0"), 0777)
	controlChan := make(chan string, 10)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				fileName := <-controlChan
				file, _ := os.Create(fileName)
				file.Close()
			}
		}()
	}
	for i := 0; i < runtime; i++ {
		controlChan <- "G:/trash/0/" + strconv.Itoa(i) + ".txt"
	}
}

func TestDataType1() {
	fmt.Println("创建测试数据 类型为100文件夹1000文件")
	fileChan := make(chan string, 8)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				os.Create(<-fileChan)
			}
		}()
	}
	os.Mkdir("G:/trash/1", 0777)
	for i := 0; i < 100; i++ {
		os.Mkdir("G:/trash/1/"+strconv.Itoa(i), 0777)
		for j := 0; j < runtime/100; j++ {
			fileChan <- "G:/trash/1/" + strconv.Itoa(i) + "/" + strconv.Itoa(
				j) + ".txt"
		}
	}
}

func TestDataType2() {
	fmt.Println("创建测试数据 类型为100文件夹100子文件夹10文件")
	fileChan := make(chan string, 8)
	folderChan := make(chan string, 8)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				os.Create(<-fileChan)
			}
		}()
		go func() {
			for {
				os.Mkdir(<-folderChan, 0777)
			}
		}()
	}
	os.Mkdir("G:/trash/2", 0777)
	for i := 0; i < 100; i++ {
		folderChan <- "G:/trash/2/" + strconv.Itoa(i)
		for j := 0; j < 100; j++ {
			folderChan <- "G:/trash/2/" + strconv.Itoa(i) + "/" + strconv.Itoa(j)
			for k := 0; k < 10; k++ {
				fileChan <- "G:/trash/2/" + strconv.Itoa(i) + "/" + strconv.Itoa(j) + "/" + strconv.Itoa(k) + ".txt"
			}
		}
	}
	time.Sleep(1 * time.Second)
}
