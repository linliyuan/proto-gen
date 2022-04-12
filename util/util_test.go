package util

import "testing"

func Test_SendToTestFile(t *testing.T) {
	err := SendToTestFile("test_out.txt", "leihou", "hahahah", "heihie")
	if err != nil {
		panic(err)
	}
}
