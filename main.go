package main

import (
	"GoRename/Check"
	"GoRename/Count"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
	"sync"
	"time"
)

//"./WAIT_TO_RENAME"
var (
	myFolder     string //待处理目录
	id           string //学号
	class        string //班级
	peopleName   string //姓名
	chapter      string //作业章节
	formatChoice string //格式
	WG           sync.WaitGroup
)

func main() {

	fmt.Println("[guide] 请输入要处理的目录路径：(请仔细检查您所填写路径)")
	fmt.Scanln(&myFolder)
	_, err := ioutil.ReadDir(myFolder)
	for err != nil {
		fmt.Println("[warning] 无法读取文件夹，请再次检查路径")
		fmt.Println("[guide] 请再次输入要处理的目录路径：")
		fmt.Scanln(&myFolder)
		_, err = ioutil.ReadDir(myFolder)
	}

	fmt.Println("[guide] 请输入您的班级:")
	fmt.Scanln(&class)
	fmt.Println("[guide] 请输入作业章节:")
	fmt.Scanln(&chapter)
	fmt.Println("[guide] 请选择您希望的格式")
	fmt.Println("[menu] [0]学号 [1]姓名 [2]班级 [3]作业章节（备注）")
	fmt.Println("[example] 100001 张三 网络206 第一章作业 对应输入：0123")
	fmt.Scanln(&formatChoice)
	for len(strings.Split(formatChoice, "")) != 4 {
		fmt.Println("[warning] 参数不正确")
		fmt.Println("[guide] 请选择您希望的格式")
		fmt.Scanln(&formatChoice)
	}

	//计时
	startTime := time.Now()

	//统计指定目录下的文件数量
	Count.CountFiles(myFolder)

	//添加WaitGroup
	WG.Add(Count.FileNums)

	files, err := ioutil.ReadDir(myFolder)
	if err != nil {
		log.Fatal("【ERROR】目录读取失败，请检查目录路径\n", err)
	}

	for _, file := range files {
		if file.IsDir() {
			continue //遇到文件夹直接略过
		}
		f := file
		//开启协程
		go func() {
			//退出前关闭一个wg
			defer WG.Done()
			//fmt.Println(f.Name())
			//属于文件，先进行匹配查询，查询成功则重命名
			var ok bool
			peopleName, id, ok = Check.NameCheck(f.Name())
			if ok {

				var format = map[string]string{
					"0": id,
					"1": peopleName,
					"2": class,
					"3": chapter,
				}

				//fmt.Println(format)

				//新旧文件路径
				oldPath := fmt.Sprintf("%v/%v", myFolder, f.Name())
				newPath := fmt.Sprintf("%v/%v %v %v %v %v", myFolder,
					format[strings.Split(formatChoice, "")[0]],
					format[strings.Split(formatChoice, "")[1]],
					format[strings.Split(formatChoice, "")[2]],
					format[strings.Split(formatChoice, "")[3]],
					path.Ext(oldPath))
				//重命名
				err := os.Rename(oldPath, newPath)
				if err != nil {
					fmt.Println(err)
				}
			}

		}()
	}

	//main主协程等待其他协程完成再退出
	WG.Wait()

	//输出未交作业的名单
	Check.SubmittedCheck()

	endTime := time.Now()
	fmt.Printf("[ 总用时 ] %v\n", endTime.Sub(startTime))

	time.Sleep(2 * 60 * time.Second)
}
