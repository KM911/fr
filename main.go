package main

import (
	"fmt"
	"github.com/KM911/oslib/fs"
	"github.com/KM911/oslib/lg"
	"github.com/panjf2000/ants/v2"
	"os"
	"path/filepath"
	"sync"
	"syscall"
	"time"
)

func RemovePool(foldername string) {
	fmt.Println("remove with pool")
	files, folders := fs.DeepDir(foldername)
	filejobs := make(chan string, 8)
	folderjobs := make(chan string, 8)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				os.Remove(filepath.Join(foldername, <-filejobs))
			}
		}()
		go func() {
			for {
				os.RemoveAll(filepath.Join(foldername, <-folderjobs))
			}
		}()
	}
	for i := 0; i < len(files); i++ {
		filejobs <- files[i]
	}
	for i := 0; i < len(folders); i++ {
		folderjobs <- folders[i]
	}
	time.Sleep(1 * time.Second)
	os.Remove(foldername)
}

func RemoveAntsPool(foldername string) {
	files, folders := fs.DeepDir(foldername)
	defer ants.Release()
	var wg sync.WaitGroup
	p, _ := ants.NewPoolWithFunc(8, func(filename interface{}) {
		remove(filepath.Join(foldername, filename.(string)))
		wg.Done()
	})
	defer p.Release()
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		p.Invoke(files[i])
	}
	for i := 0; i < len(folders); i++ {
		wg.Add(1)
		p.Invoke(folders[i])
	}
	wg.Wait()
	remove(foldername)
}

func Remove(foldername string) {
	fmt.Println("remove")
	files, folders := fs.DeepDir(foldername)
	for i := 0; i < len(files); i++ {
		os.Remove(filepath.Join(foldername, files[i]))
	}
	for i := 0; i < len(folders); i++ {
		os.RemoveAll(filepath.Join(foldername, folders[i]))
	}
	os.Remove(foldername)
}

//func remove(src string) {
//	//os.Remove()
//	if fs.IsDir(src) {
//		syscall.Rmdir(src)
//	} else {
//		syscall.Unlink(src)
//	}
//}

func remove(src string) {
	syscall.Rmdir(src)
	syscall.Unlink(src)
}
func main() {
	lg.SingleLogger(filepath.Join(fs.ExecutePath(), "fr.log"))
	defer lg.Recover()
	fr()
}
func fr() {
	//panic("test painc")
	argsLens := len(os.Args)
	// 将权限提升为管理员权限
	switch argsLens {
	case 1:
		fmt.Println("需要有效的参数")
		fmt.Println("例如 fr.exe E:/trash/test")
		fmt.Println("将会删除 E:/trash/test 文件夹 和其下的所有文件和文件夹")
	default:
		for i := 1; i < argsLens; i++ {
			if fs.IsDir(os.Args[i]) {
				if filepath.IsAbs(os.Args[i]) {
					RemoveAntsPool(os.Args[i])
				} else {
					RemoveAntsPool(filepath.Join(fs.CmdPath(), os.Args[i]))
				}
			} else {
				//fmt.Println("参数需要为文件夹路径 而不是文件路径")
				os.Remove(os.Args[i])
				//time.Sleep(3 * time.Second)
			}
		}
	}
}
