package DoSomethingWithFile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
	"log"
	"strings"
)

// os、io/ioutil中提供了readFile方法可以快速读取全文
// io/ioutil中提供了ReadAll方法在打开文件句柄后可以读取全文
// 读取全文件
func ReadSomethingForFile1(filename string) error {
	// 直接读
	content, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	log.Printf("read %s content is %s", filename, content)
	return nil
}

func ReadSomethingForFile2(filename string) error {
	// 创建文件
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	// 读文件
	content, err := ioutil.ReadAll(file)
	log.Printf("read %s content is %s\n", filename, content)

	file.Close()
	return nil
}

// 逐行读取
func ReadSomethingForFile3(filename string) error {

	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	// 为file建立buffered read
	bufferedReader := bufio.NewReader(file)
	// 不断读
	for {
		// bufio中提供了三种方法ReadLine、ReadBytes("\n")、ReadString("\n")可以按行读取数据。
		lineBytes, err := bufferedReader.ReadBytes('\n')
		bufferedReader.ReadLine()
		line := strings.TrimSpace(string(lineBytes))
		if err != nil && err != io.EOF {
			return err
		}
		// 读完了
		if err == io.EOF {
			break
		}
		log.Printf("readline %s every line data is %s\n", filename, line)
	}
	file.Close()
	return nil
}

// 按块读取文件
// os库配合bufio.NewReader调用Read方法
// os库的Read方法
// os库配合io库的ReadFull、ReadAtLeast方法

// use bufio.NewReader
func ReadSomethingForFile4(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}
	// 创建 Reader
	r := bufio.NewReader(file)

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := r.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// use os
func ReadSomethingForFile5(filename string) error{
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := file.Read(buf)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}


// use os and io.ReadAtLeast
func ReadSomethingForFile6(filename string) error{
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	// 每次读取 2 个字节
	buf := make([]byte, 2)
	for {
		n, err := io.ReadAtLeast(file, buf, 0)
		if err != nil && err != io.EOF {
			return err
		}

		if n == 0 {
			break
		}
		log.Printf("writeByte %s every read 2 byte is %s\n", filename, string(buf[:n]))
	}
	file.Close()
	return nil
}

// 分隔符读取
func ReadSomethingForFile7(filename string) error {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	// 可以定制Split函数做分隔函数
	// ScanWords 是scanner自带的分隔函数用来找空格分隔的文本字
	scanner.Split(bufio.ScanWords)
	for {
		success := scanner.Scan()
		if success == false {
			// 出现错误或者EOF是返回Error
			err = scanner.Err()
			if err == nil {
				log.Println("Scan completed and reached EOF")
				break
			} else {
				return err
			}
		}
		// 得到数据，Bytes() 或者 Text()
		log.Printf("readScanner get data is %s", scanner.Text())
	}
	file.Close()
	return nil
}