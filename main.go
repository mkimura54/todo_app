package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"runtime"
	"strings"
)

const FILENAME string = "todo.json"

var todos TodoList

func main() {
	addContent := flag.String("a", "", "追加するTODOの内容")
	deleteNo := flag.Int("d", 0, "削除するTODOの番号")
	flag.Parse()

	if countArgs() > 1 {
		fmt.Println("複数のオプションを同時に指定することはできません。")
		return
	}

	loadFile()
	if *addContent != "" {
		todos.AddTodo(*addContent)
		saveFile(todos)
	} else if *deleteNo != 0 {
		todos.DeleteTodo(*deleteNo)
		saveFile(todos)
	} else {
		showTodoList()
	}
}

func countArgs() int {
	var count int
	for i := 1; i < len(os.Args); i++ {
		if strings.Contains(os.Args[i], "-") {
			count++
		}
	}
	return count
}

func showTodoList() {
	for _, todo := range todos.Todos {
		fmt.Printf("%d : %s\n", todo.No, todo.Content)
	}
}

func getDataFilePath() string {
	var fileDirPath string
	if runtime.GOOS == "linux" {
		fileDirPath = "/home/"
	}

	return path.Join(fileDirPath, FILENAME)
}

func loadFile() bool {
	filePath := getDataFilePath()
	_, err := os.Stat(filePath)
	if err != nil {
		saveFile(TodoList{})
	}

	bytes, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("データ読み込みエラー(%s)", err.Error())
		return false
	}

	err = json.Unmarshal(bytes, &todos)
	if err != nil {
		fmt.Printf("データ変換エラー(%s)", err.Error())
		return false
	}

	return true
}

func saveFile(todos TodoList) bool {
	bytes, err := json.Marshal(todos)
	if err != nil {
		fmt.Printf("データ変換エラー(%s)", err.Error())
		return false
	}

	fp, err := os.Create(getDataFilePath())
	if err != nil {
		fmt.Printf("ファイルOPENエラー(%s)", err.Error())
		return false
	}
	defer fp.Close()

	_, err = fp.Write(bytes)
	if err != nil {
		fmt.Printf("ファイル書き込みエラー(%s)", err.Error())
		return false
	}

	return true
}
