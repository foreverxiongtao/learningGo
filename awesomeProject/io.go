package main

import (
	"os"
	"fmt"
)

func main() {
	//createFile()
	readFile()
}

//创建文件
func createFile() {

	userFile := "test.txt"
	fout, err := os.Create(userFile)
	defer fout.Close()
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < 5; i++ {
		fout.WriteString("hello" + string(i) + "\n")
	}
}

//读取文件
func readFile() {
	userFile := "test.txt"
	f, err := os.Open(userFile)
	if err != nil {
		fmt.Println(err)
	}
	buffer := make([]byte, 1024)
	for {
		n, _ := f.Read(buffer)
		if n == 0 {
			break
		}
		os.Stdout.Write(buffer[:n])
	}
}

//删除文件
func deleteFile() {
	os.Remove("test.txt")
}
