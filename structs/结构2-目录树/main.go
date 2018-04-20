package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type FileTreeNode struct {
	Name         string          `json:"name"`
	Created      time.Time       `json:"created"`
	Updated      time.Time       `json:"updated"`
	Path         string          `json:"path"`
	FileTreeNode []*FileTreeNode `json:"children"`
}

func main() {
	// fmt.Println("请输入完整根目录: ")
	// inputReader := bufio.NewReader(os.Stdin)
	// rootPath, err := inputReader.ReadString('\n')
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Printf("输入的根目录为：%s\n", rootPath)
	// getFilelist(rootPath)

	flag.Parse()
	root := flag.Arg(0)

	fmt.Println("-------------showFileList--------------")
	showFileList(root)
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if info == nil {

		// 文件名称超过限定长度等其他问题也会导致info == nil
		// 如果此时return err 就会显示找不到路径，并停止查找。
		println("can't find:(" + path + ")")
		return nil
	}

	if info.IsDir() {
		tree := make(map[string]FileTreeNode)
		println("Folder:" + path)
		return nil
	} else {
		println("-------" + path)
		return nil
	}

}

func showFileList(root string) {
	err := filepath.Walk(root, walkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	return
}
