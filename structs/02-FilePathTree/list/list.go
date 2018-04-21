package list

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"time"
)

//文件节点结构体
type FileNode struct {
	Name      string      `json:"name"`
	Size      int64       `json:"size"`
	Updated   time.Time   `json:"updated"`
	Path      string      `json:"path"`
	FileNodes []*FileNode `json:"children"`
}

//直接展示文件序列
func ShowFileList(root string) error {
	err := filepath.Walk(root, WalkFunc)
	if err != nil {
		fmt.Printf("filepath.Walk() error: %v\n", err)
	}
	return err
}

func WalkFunc(pathOfFile string, info os.FileInfo, err error) error {

	if info == nil {

		// 文件名称超过限定长度等其他问题也会导致info == nil
		// 如果此时return err 就会显示找不到路径，并停止查找。
		println("can't find:(" + pathOfFile + ")")
		return err
	}

	if info.IsDir() {
		println("Folder:" + pathOfFile)
		return nil
	} else {
		println("-------" + pathOfFile)
		return nil
	}

}

//对文件夹遍历
func Walk(pathOfFile string, info os.FileInfo, node *FileNode) {

	//首先判断是否为文件夹
	if !info.IsDir() {
		fmt.Println(pathOfFile, "is NOT a directory")
		return
	}

	// 列出当前目录下的所有目录、文件
	files := ListFiles(pathOfFile)

	// 遍历这些文件
	for _, filename := range files {
		// 拼接全路径
		fpath := filepath.Join(pathOfFile, filename)

		// 构造文件结构
		fio, _ := os.Lstat(fpath)

		// 将当前文件作为子节点添加到目录下
		child := FileNode{
			Name:      filename,
			Size:      fio.Size(),
			Updated:   fio.ModTime(),
			Path:      fpath,
			FileNodes: []*FileNode{},
		}
		node.FileNodes = append(node.FileNodes, &child)

		// 如果遍历的当前文件是个目录，则进入该目录进行递归
		if fio.IsDir() {
			Walk(fpath, fio, &child)
		}
	}

	return
}

//列出并排序所有文件名
func ListFiles(dirname string) []string {
	f, err := os.Open(dirname)
	defer f.Close()
	Check(err)

	//展示所有文件
	names, err := f.Readdirnames(0)
	Check(err)

	//按文件名排序
	sort.Strings(names)

	return names
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
