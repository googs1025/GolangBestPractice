package main

import (
	"fmt"
	"github.com/pkg/errors"
)

func main()  {

	err := NewError("错误处理")
	fmt.Println(err)

}


/*
    Error vs Exception

	Go内部使用 error 接口类型作为 go 的错误标准的处理流程
	我们可以建立一个对象(ex：errorString)实现error接口(就是写出内部的Error() string 方法)。
    错误是程序中可能出现的问题，比如连接数据库失败，连接网络失败等，在程序设计中，错误处理是业务的一部分。
*/

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}


func NewError(text string) error {
	return &errorString{text}
}


// Exception异常
// 异常是指在不该出现问题的地方出现问题，是预料之外的，比如空指针引用，下标越界，向空 map 添加键值等
// 人为制造被自动触发的异常，比如：数组越界，向空 map 添加键值对等。
// 手工触发异常并终止异常，比如：连接数据库失败主动 panic。

// panic恐慌
// 对于真正意外的情况，那些表示不
//可恢复的程序错误，不可恢复才使用 panic。对于其他的错误情况，我们应该是期望使用 error 来进行判定
// 理论上 panic 只存在于 server 启动阶段，比如 config 文件解析失败，端口监听失败等等，所有业务逻辑禁止主动 panic，所有异步的 goroutine 都要用 recover 去兜底处理。




// ZooTour struct

type Panda struct {

}

// 接口
type ZooTour interface {
	Enter() error
	VisitPanda(panda *Panda) error
	Leave() error
}

// 分步处理，每个步骤可以针对具体返回结果进行处理
func Tour(t ZooTour, panda *Panda) error {

	// 调用第一个方法
	if err := t.Enter(); err != nil {
		return errors.WithMessage(err, "Enter failed.")
	}
	// 调用第二种方法
	if err := t.VisitPanda(panda); err != nil {
		return errors.WithMessage(err, "VisitPanda failed.")
	}
	// ...

	return nil
}



//

type ZooTour1 interface {
	Enter() error
	VisitPanda(panda *Panda) error
	Leave() error
	Err() error
}

func Tour(t ZooTour1, panda *Panda) error {

	// 先调用方法
	t.Enter()
	t.VisitPanda(panda)
	t.Leave()

	// 集中编写业务逻辑代码,最后统一处理error
	if err := t.Err(); err != nil {
		return errors.WithMessage(err, "ZooTour failed")
	}
	return nil
}

// 第三种方法

type MyFunc func(t ZooTour) error

type Walker interface {
	Next MyFunc
}
type SliceWalker struct {
	index int
	funs []MyFunc
}

func NewEnterFunc() MyFunc {
	return func(t ZooTour) error {
		return t.Enter()
	}
}

func BreakOnError(t ZooTour, walker Walker) error {
	for {
		f := walker.Next()
		if f == nil {
			break
		}
		if err := f(t); err := nil {
			// 遇到错误break或者continue继续执行
		}
	}
}

