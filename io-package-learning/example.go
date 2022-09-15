package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	UseReader1()
	UseWrite1()
	//UseFileWrite()
	UseFileRead()
	//UseCopy1()
	//UseCopy2()
	UseWriteString()
	UseBufferRead()
	UseIOutil()


}


/*
	https://blog.csdn.net/qq_39780174/article/details/115318438?spm=1001.2101.3001.6650.2&utm_medium=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-115318438-blog-110822066.pc_relevant_default&depth_1-utm_source=distribute.pc_relevant.none-task-blog-2%7Edefault%7ECTRLIST%7ERate-2-115318438-blog-110822066.pc_relevant_default&utm_relevant_index=4
	io.Reader
	io.Reader 表示一个读取器，它将数据从某个资源读取到传输缓冲区。在缓冲区中，数据可以被流式传输和使用。
    对于要用作读取器的类型，它必须实现 io.Reader 接口的唯一一个方法 Read(p []byte)。
	换句话说，只要实现了 Read(p []byte) ，那它就是一个读取器。

	Read() 方法有两个返回值，一个是读取到的字节数，一个是发生错误时的错误。
	同时，如果资源内容已全部读取完毕，应该返回 io.EOF 错误。
 */

func UseReader1() {
	str := "I love Golang!! djaklfja;ldf"
	// strings包中的Reader就实现了 Read方法
	reader := strings.NewReader(str)
	buffer := make([]byte, 5)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取结束!", n)
				break
			}
			panic(err)
		}
		fmt.Println(n, string(buffer[:n]))
	}

}

/*
 	io.Writer
    io.Writer 表示一个编写器，它从缓冲区读取数据，并将数据写入目标资源。
	对于要用作编写器的类型，必须实现 io.Writer 接口的唯一一个方法 Write(p []byte)
	同样，只要实现了 Write(p []byte) ，那它就是一个编写器。

    Write() 方法有两个返回值，一个是写入到目标资源的字节数，一个是发生错误时的错误。
 */


func UseWrite1() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}

	var writer bytes.Buffer

	for _, v := range proverbs {
		n, err := writer.Write([]byte(v))
		if err != nil {
			panic(err)
		}

		if n != len(v) {
			panic(err)
		}
	}

	fmt.Println(writer.String())

}

/*
	os.File
	类型 os.File 表示本地系统上的文件。它实现了 io.Reader 和 io.Writer ，因此可以在任何 io 上下文中使用。
 */

func UseFileWrite() {
	proverbs := []string{
		"Channels orchestrate mutexes serialize\n",
		"Cgo is not Go\n",
		"Errors are values\n",
		"Don't panic\n",
	}
	file, err := os.Create("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/input.text")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	for _, v := range proverbs {
		n, err := file.Write([]byte(v))
		if err != nil {
			panic(err)
		}
		if n != len(v) {
			fmt.Println("failed to write data")
			os.Exit(1)
		}
	}
	fmt.Println("写操作结束")

}

func UseFileRead() {

	file, err := os.Open("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/input.text")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	buffer := make([]byte, 4)
	for  {
		n, err := file.Read(buffer)
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取结束")
				break
			}

		}
		fmt.Println(string(buffer[:n]))
	}
	fmt.Println("写操作结束")

}


/*
	io.Copy()
	io.Copy() 可以轻松地将数据从一个 Reader 拷贝到另一个 Writer。
	它抽象出 for 循环模式（我们上面已经实现了）并正确处理 io.EOF 和 字节计数。
 */

func UseCopy1() {

	proverbs := new(bytes.Buffer)
	proverbs.WriteString("Channels orchestrate mutexes serialize\n")
	proverbs.WriteString("Cgo is not Go\n")
	proverbs.WriteString("Errors are values\n")
	proverbs.WriteString("Don't panic\n")

	file, err := os.Create("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	// io.Copy 完成了从 proverbs 读取数据并写入 file 的流程
	if _, err := io.Copy(file, proverbs); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("file created")

}

// 可以使用 io.Copy() 函数重写从文件读取并打印到标准输出的先前程序
func UseCopy2() {

	file, err := os.Open("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := io.Copy(os.Stdout, file); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

// io.WriteString()方法可以将字符串类型写入一个 Writer
func UseWriteString() {
	file, err := os.Create("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/magic_msg.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()
	if _, err := io.WriteString(file, "Go is fun!"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

/*
	缓冲区 io
    标准库中 bufio 包支持 缓冲区 io 操作，可以轻松处理文本内容。
	例如，以下程序逐行读取文件的内容，并以值 '\n' 分隔：
 */

func UseBufferRead() {
	file, err := os.Open("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/proverbs.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("读取完毕")
				break
			}
			panic(err)

		}
		fmt.Print(line)
	}
}

// 使用函数 ReadFile 将文件内容加载到 []byte 中
func UseIOutil() {
	bytes, err := ioutil.ReadFile("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/proverbs.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s", string(bytes))
}


