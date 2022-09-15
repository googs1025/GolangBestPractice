package DoSomethingWithFile

import (
	"bufio"
	"log"
	"os"
)

// os/ioutil包都提供了WriteFile方法可以快速处理创建/打开文件/写数据/关闭文件
// 1. 快速写文件
func WriteSomethingForFile1(filename string) error {

	// os模块有WriteFile ReadFile方法
	err := os.WriteFile(filename, []byte("hello world\n"), 0666)
	if err != nil {
		log.Fatal("写文件失败！ err=%s\n", err)
		return err
	}
	return nil
}

// 2. 按行写文件
// os、buffo写数据都没有提供按行写入的方法，所以我们可以在调用os.WriteString、bufio.WriteString方法是在数据中加入换行符即可

// 直接进行IO操作
func WriteSomethingForFile2(filename string) error {
	// 数据
	data := []string{
		"asong",
		"test",
		"123",
	}

	// 创建文件
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil{
		return err
	}

	for _, line := range data{
		// 记得加入换行符
		_,err := file.WriteString(line + "\n")
		if err != nil{
			return err
		}
	}
	file.Close()
	return nil
}
// 使用缓存区写入
func WriteSomethingForFile3(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 为这个文件创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	for i:=0; i < 2; i++{
		// 写字符串到buffer
		bytesWritten, err := bufferedWriter.WriteString(
			"hello world\n",
		)
		if err != nil {
			return err
		}
		log.Printf("Bytes written: %d\n", bytesWritten)
	}
	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil{
		return err
	}

	file.Close()
	return nil
}


// 偏移量写入
// 想根据给定的偏移量写入数据，可以使用os中的writeAt方法
func WriteSomethingForFile4(filename string) error {
	// 数据
	data := []byte{
		0x41, // A
		0x73, // s
		0x20, // space
		0x20, // space
		0x67, // g
	}

	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil{
		return err
	}

	_, err = file.Write(data)
	if err != nil{
		return err
	}

	replaceSplace := []byte{
		0x6F, // o
		0x6E, // n
	}

	// 使用偏移位写入
	_, err = file.WriteAt(replaceSplace, 2)
	if err != nil{
		return err
	}
	file.Close()
	return nil
}

// 缓存区写入
// os库中的方法对文件都是直接的IO操作，频繁的IO操作会增加CPU的中断频率，
// 所以我们可以使用内存缓存区来减少IO操作，在写字节到硬盘前使用内存缓存，
// 当内存缓存区的容量到达一定数值时在写内存数据buffer到硬盘
func WriteSomethingForFile5(filename string) error {

	file, err := os.OpenFile(filename, os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	// 为file创建buffered writer
	bufferedWriter := bufio.NewWriter(file)

	// 写字符串到buffer
	bytesWritten, err := bufferedWriter.WriteString(
		"hello world\n",
	)
	if err != nil {
		return err
	}
	log.Printf("Bytes written: %d\n", bytesWritten)

	// 检查缓存中的字节数
	unflushedBufferSize := bufferedWriter.Buffered()
	log.Printf("Bytes buffered: %d\n", unflushedBufferSize)

	// 还有多少字节可用（未使用的缓存大小）
	bytesAvailable := bufferedWriter.Available()
	log.Printf("Available buffer: %d\n", bytesAvailable)

	// 写内存buffer到硬盘
	err = bufferedWriter.Flush()
	if err != nil{
		return err
	}

	file.Close()
	return nil
}

