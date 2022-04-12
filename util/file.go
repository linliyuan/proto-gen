package util

import (
	"bufio"
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func SendToTestFile(filePath string, arr ...interface{}) error {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		return err
	}

	//及时关闭file句柄
	defer func() {
		_ = file.Close()
	}()

	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	_, err = fmt.Fprintln(write, arr...)
	if err != nil {
		return err
	}

	//Flush将缓存的文件真正写入到文件中
	return write.Flush()
}
