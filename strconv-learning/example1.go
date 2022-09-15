package main

/*
	go需要自行做类型转换  转换数据类型
    ex: valueTypeB = typeB(valueTypeA)

 */

/*
	总结：
	1. 不是所有数据类型都能转换的，例如字母格式的string类型"abcd"转换为int肯定会失败
    2. 这种简单的转换方式不能对int(float)和string进行互转，要跨大类型转换，可以使用strconv包提供的函数
    3. 低精度转换为高精度时是安全的，高精度的值转换为低精度时会丢失精度。例如int32转换为int16，float32转换为int

 */

func main()  {

	// 事例一：简单类型转换
	a := 5.0
	b := int(a)
	println(b)

	// 事例二：类型转换
	// IT底层的类型还是 int
	type IT int
	// a 的类型是IT 但是底层还是为int
	var aa IT = 5
	println(aa)
	// 把 a(IT)转回 int类型
	bb := int(aa)
	println(bb)
	// 把 b(int)转成 IT类型
	c := IT(bb)
	println(c)



}
