package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/luckyOrangeGo/studyLog/structs/02-FilePathTree/list"
	"github.com/luckyOrangeGo/studyLog/structs/02-FilePathTree/saveLog"
	"github.com/luckyOrangeGo/studyLog/structs/02-FilePathTree/strTimeNow"
)

type FileTreeNode struct {
	Name         string          `json:"name"`
	Created      time.Time       `json:"created"`
	Updated      time.Time       `json:"updated"`
	Path         string          `json:"path"`
	FileTreeNode []*FileTreeNode `json:"children"`
}

func main() {

	// flag.Parse()
	// root := flag.Arg(0)

	flag.Parse()
	rootpath := flag.Arg(0)

	fmt.Println("-------------showFileList--------------")
	err := list.ShowFileList(rootpath)
	list.Check(err)

	fileInfo, _ := os.Lstat(rootpath)

	root := list.FileNode{
		Name:      path.Base(rootpath),
		Path:      rootpath,
		Size:      fileInfo.Size(),
		Updated:   fileInfo.ModTime(),
		FileNodes: []*list.FileNode{},
	}

	list.Walk(rootpath, fileInfo, &root)

	data, _ := json.Marshal(root)

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, data, "", "\t")
	if err != nil {
		log.Println("JSON parse error: ", err)
		return
	}

	//log.Println("CSP Violation:", string(prettyJSON.Bytes()))

	tm := strTimeNow.StrTm()

	fileName := "FilePath-" + rootpath + "-At-" + tm + ".json"

	err = saveLog.SaveFile(fileName, prettyJSON.Bytes())

	if err != nil {
		log.Println("Save Faild! Error: ", err)
	} else {
		log.Println(fileName, "Save Success!")
	}

}
