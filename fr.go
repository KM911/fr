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
)

func RemovePool(foldername string) {
	fmt.Println("remove with pool")
	files, folders := fs.DeepDir(foldername)
	//这里chan的大小会影响吗? 我没有试过 这里需要大量的测试工作了
	filejobs := make(chan string, 1)
	for i := 0; i < 8; i++ {
		go func() {
			for {
				syscall.Unlink(filepath.Join(foldername, <-filejobs))
			}
		}()
	}
	for i := 0; i < len(files); i++ {
		filejobs <- files[i]
	}
	// 这里理论上会因为删除文件夹还没有完成导致文件夹删除失败
	for i := 0; i < len(folders); i++ {
		syscall.Rmdir(filepath.Join(foldername, folders[i]))
	}
	syscall.Rmdir(foldername)
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

//
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
	argsLens := len(os.Args)
	switch argsLens {
	case 1:
		fmt.Println("需要有效的参数")
		fmt.Println("例如 fr.exe E:/trash/test")
		fmt.Println("将会删除 E:/trash/test 文件夹 和其下的所有文件和文件夹")
	default:
		for i := 1; i < argsLens; i++ {
			if fs.IsDir(os.Args[i]) {
				if filepath.IsAbs(os.Args[i]) {
					RemovePool(os.Args[i])
				} else {
					RemovePool(filepath.Join(fs.CmdPath(), os.Args[i]))
				}
			} else {
				os.Remove(os.Args[i])
			}
		}
	}
}
