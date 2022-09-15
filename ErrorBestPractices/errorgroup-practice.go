package main

import (
	"time"
	"golang.org/x/sync/errgroup"
	"context"
	"fmt"
	"log"
)

// https://marksuper.xyz/2021/10/15/error_group/

func main() {
	TestErrgroup()
}


/*
	1.通过 WithContext 可以创建一个带取消的 Group
	2.当然除此之外也可以零值的 Group 也可以直接使用，但是出错之后就不会取消其他的 goroutine 了
	3.Go 方法传入一个 func() error 内部会启动一个 goroutine 去处理
	4.Wait 类似 WaitGroup 的 Wait 方法，等待所有的 goroutine 结束后退出，返回的错误是一个出错的 err

*/

func TestErrgroup() {
	// errgroup建立上下文
	eg, ctx := errgroup.WithContext(context.Background())

	// 启100个goroutine
	for i := 0; i < 100; i++ {
		i := i
		// 启
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			select {
			case <-ctx.Done():
				fmt.Println("Canceled:", i)
				return nil
			default:
				fmt.Println("End:", i)
				return nil
			}})}
	// 会阻塞在这
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}
}

// errgroup拓展包
