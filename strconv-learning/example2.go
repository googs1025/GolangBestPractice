package main

import (
	"fmt"
	"strconv"
)

/*
	strconv包简介：
    https://www.cnblogs.com/f-ck-need-u/p/9863915.html
    strconv包提供了字符串与简单数据类型之间的类型转换功能。可以将简单类型转换为字符串，也可以将字符串转换为其它简单类型。

	字符串转int：Atoi()
	int转字符串: Itoa()
	ParseTP类函数将string转换为TP类型：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。因为string转其它类型可能会失败，所以这些函数都有第二个返回值表示是否转换成功
	FormatTP类函数将其它类型转string：FormatBool()、FormatFloat()、FormatInt()、FormatUint()
	AppendTP类函数用于将TP转换成字符串后append到一个slice中：AppendBool()、AppendFloat()、AppendInt()、AppendUint()
 */

func main() {

	println("a" + strconv.Itoa(4222))

	i, _ := strconv.Atoi("3")
	println(5 + i)

	res, err := strconv.Atoi("a")
	if err != nil {
		fmt.Errorf("converted failed")
	}
	println(res)


	// Parse类函数用于转换字符串为给定类型的值：ParseBool()、ParseFloat()、ParseInt()、ParseUint()。
	b, err := strconv.ParseBool("true")
	f, err := strconv.ParseFloat("3.1415", 64)
	ii, err := strconv.ParseInt("-42", 10, 64)
	u, err := strconv.ParseUint("42", 10, 64)
	fmt.Println(b, f, ii, u)



	// 将给定类型格式化为string类型：FormatBool()、FormatFloat()、FormatInt()、FormatUint()。
	s1 := strconv.FormatBool(true)
    // func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	s2 := strconv.FormatFloat(3.1415, 'E', -1, 64)
	// func FormatInt(i int64, base int) string
	// func FormatUint(i uint64, base int) string
	// 第二个参数base指定将第一个参数转换为多少进制
	s3 := strconv.FormatInt(-42, 16)
	s4 := strconv.FormatUint(42, 16)
	fmt.Println(s1, s2, s3, s4)
}
