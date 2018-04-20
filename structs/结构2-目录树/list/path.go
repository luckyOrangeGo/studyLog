package list

import (
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

func walkFunc(path string, info os.FileInfo, err error) error {
	if info == nil {

		// 文件名称超过限定长度等其他问题也会导致info == nil
		// 如果此时return err 就会显示找不到路径，并停止查找。
		println("can't find:(" + path + ")")
		return nil
	}

	if info.IsDir() {
		tree := make(map[string]FileTreeNode)
		tree[path] = FileTreeNode{}
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
