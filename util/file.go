package util

import (
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

func SendToTestFile(arr ...interface{}) error {
	out, _ := os.Open("./test_out.text")

	_, err := fmt.Fprintln(out, arr...)
	return err
}
