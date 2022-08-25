package Count

import (
	"fmt"
	"io/ioutil"
	"log"
)

var FileNums int = 0

func CountFiles(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal("【ERROR】目录读取失败，请检查目录路径\n", err)
	}
	for _, file := range files {
		if file.IsDir() {
			continue //遇到文件夹直接略过
		}
		//属于文件，计数器加1
		FileNums++
		fmt.Println(file.Name())
	}

	fmt.Printf("【tips】该目录下共有文件  %d  个\n", FileNums)
}
