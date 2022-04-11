package main

import (
	"bytes"
	"fmt"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/pluginpb"
	"io/ioutil"
	"os"
)

func sendToTestFile(arr ...interface{}) error {
	out, _ := os.Create("./test_out.text")

	_, err := fmt.Fprintln(out, arr...)
	return err
}

func main() {
	// Protoc 将protobuf文件编译为 plugin pb.CodeGeneratorRequest 结构，并输出到stdin中
	input, _ := ioutil.ReadAll(os.Stdin)
	var req pluginpb.CodeGeneratorRequest
	err := proto.Unmarshal(input, &req)
	if err != nil {
		panic(err)
	}

	// 使用默认选项初始化我们的插件
	opts := protogen.Options{}
	plugin, err := opts.New(&req)
	if err != nil {
		panic(err)
	}

	// protoc 将一组文件结构传递给程序处理
	for _, file := range plugin.Files {
		_ = sendToTestFile("file.GeneratedFilenamePrefix", file.GeneratedFilenamePrefix)
		_ = sendToTestFile("file.GoPackageName", file.GoPackageName)

		// 是时候生成代码了……！

		// 1. 初始化缓冲区以保存生成的代码
		var buf bytes.Buffer

		// 2. 生成包名称
		pkg := fmt.Sprintf("package %s", file.GoPackageName)
		buf.Write([]byte(pkg))

		// 3. 为每个message生成 Foo() 方法
		for _, msg := range file.Proto.MessageType {
			buf.Write([]byte(fmt.Sprintf(`
            func (x %s) Foo() string {
               return "bar"
            }`, *msg.Name)))
		}

		// 4. 指定输出文件名，在这种情况下为test.foo.go
		filename := file.GeneratedFilenamePrefix + ".foo.go"
		file := plugin.NewGeneratedFile(filename, ".")

		// 5. 将设概念车呢个的代码，从缓冲区写入到文件
		_, err = file.Write(buf.Bytes())
		if err != nil {
			panic(err)
		}
	}

	// 从我们的插件生成响应,并将其编组为protobuf
	stdout := plugin.Response()
	out, err := proto.Marshal(stdout)
	if err != nil {
		panic(err)
	}

	// 相应输出到stdout, 它将被 protoc 接收
	_, err = fmt.Fprintf(os.Stdout, string(out))
	if err != nil {
		panic(err)
	}
}
