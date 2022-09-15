package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"time"
)

/*
    https://blog.csdn.net/xixihahalelehehe/article/details/110822066
	io.Pipe使用:
    io.Pipe方法的返回值是一个读取器Reader与写入器Writer，当对Reader读取 or 对Writer写入，goroutine会被锁住，直到Writer有新的数据进来或关闭 or Reader把数据读走
 */


func main() {
	//UsePipeErrorWay()
	//UsePipe1()
	//UsePipe2()
	UsePipe3()
}

// 此方法会造成死锁产生。
func UsePipeErrorWay() {
	reader, writer := io.Pipe()

	// writer写入后会锁住进程
	writer.Write([]byte("hello"))
	defer writer.Close()
	// Reader读取后会锁住进程
	buffer := make([]byte, 100)
	reader.Read(buffer)
	fmt.Println(string(buffer))

}

// 有两种解决方法：
// 1. 建立子goroutine 来写writer

func UsePipe1() {
	reader, writer := io.Pipe()
	stopC := make(chan struct{})

	go func() {

		for i := 0; i < 10; i++ {
			writer.Write([]byte("hello world"))
		}
		defer writer.Close()
		//wg.Done()
		stopC <-struct{}{}
	}()

	for {
		select {

		case <-stopC:
			fmt.Println("写请求已经结束!")
			fmt.Println("主goroutine退出")
			return

		default:
			buffer := make([]byte, 100)
			_, err := reader.Read(buffer)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(buffer))
		}


	}

}

func UsePipe2() {

	reader, writer := io.Pipe()
	defer writer.Close()
	stopC := make(chan struct{})


	go func() {
		buffer := make([]byte, 100)
		for {
			select {
			case <-stopC:
				fmt.Println("写请求已结束，关闭读goroutine！")
				return
			case <-time.After(time.Second):
				_, err := reader.Read(buffer)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(buffer))

			}

		}

	}()

	for i := 0; i < 100; i++ {
		writer.Write([]byte("hello world"))
	}

	stopC <- struct{}{}

}

func UsePipe3() {
	//res, _ := os.Getwd()
	//fmt.Println(res)
	file, err := os.OpenFile("/Users/zhenyu.jiang/go/src/golanglearning/new_project/for-gong-zhong-hao/BestPractices/io-package-learning/input-try", os.O_RDONLY, 0755)
	if err != nil {
		panic(err)
	}
	cmd := exec.Command("grep", "world")
	reader, writer := io.Pipe()
	cmd.Stdin = reader
	cmd.Stdout = os.Stdout

	go func() {
		_, err := io.Copy(writer, file)
		if err != nil {
			fmt.Println(err)
		}
		defer writer.Close()
	}()

	cmd.Run()
}
