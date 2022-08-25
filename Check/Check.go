package Check

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

var NameListPath = "./NameList.csv"

var SubmittedStatus = map[string]bool{} //作业提交状态

var NotSubmitted []string //未提交作业名单

func NameCheck(fileName string) (string, string, bool) {
	file, err := os.Open(NameListPath) //只读状态打开名单
	defer file.Close()                 //函数退出前关闭文件

	if err != nil {
		log.Fatal("【ERROR】名单读取失败，请在对应目录中创建 NameList.csv \n", err)
	}

	fileScanner := csv.NewReader(file)

	//逐行读取文件
	for {
		row, err := fileScanner.Read()

		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		//fmt.Println(row)
		//匹配
		if strings.Contains(fileName, row[0]) {
			SubmittedStatus[row[0]] = true //加入已提交字典，并将状态置1 已提交
			fmt.Printf("%s Check\n", row[0])
			return row[0], row[1], true
		}
	}

	//无匹配
	return "", "", false
}

func SubmittedCheck() {
	file, err := os.Open(NameListPath) //只读状态打开名单
	defer file.Close()                 //函数退出前关闭文件

	if err != nil {
		log.Fatal("【ERROR】名单读取失败，请在对应目录中创建 NameList.csv \n", err)
	}

	fileScanner := csv.NewReader(file)

	//逐行读取文件
	for {
		row, err := fileScanner.Read()

		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}

		if !SubmittedStatus[row[0]] { //若名字未存在于已提交名单
			//加入未提交作业名单
			NotSubmitted = append(NotSubmitted, row[0])
		}
	}

	//输出未提交作业的名单
	if len(NotSubmitted) > 0 {
		fmt.Printf("[NotSubmitted] 未提交作业 %v\n", NotSubmitted)
	} else {
		fmt.Printf("[AllSubmitted] 作业全部提交完毕！ %v\n", NotSubmitted)
	}
}
