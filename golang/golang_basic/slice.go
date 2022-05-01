package golang_basic

import (
	"fmt"
)

// https://zhuanlan.zhihu.com/p/34041570

func SliceOperation() {

	var a []int32 // slice == nil
	// var a1 []int32{}   // 不合法

	var b = make([]int32)
	var b1 = make([]int32, 0, 100) // 长度为0，容量为100 len() ==0; cap() == 100
	var b2 = make([]int32, 1)      // 未指定容量：长度为1，初始容量为1; len() = 1

	c := make([]int32)
	d := []int32{}
	d := []int32{11, 12, 13}
	e := []int32(nil) // slice == nil

	// -----------------------------------------------------------
	// 长度&容量
	var size int  // 下面返回的类型都是int
	size = len(c) // slice长度
	size = cap(c) // slice容量

	// -----------------------------------------------------------
	// 追加
	//
	c = append(c, 1)
	c = append(c, 1, 2, 3)
	c = append(c, b...)

	// -----------------------------------------------------------
	// 拷贝
	//
	var copiedSize int = copy(c, d) // 深拷贝
	cShadow := c                    // shadow copy

	// deep copy
	cDeep := make([]int32, len(c))
	copy(cDeep, c)

	// -----------------------------------------------------------
	// 清空
	//
	c = c[:0]

	// -----------------------------------------------------------
	// 删除某一部分
	//
	c = append(c[0:5], c[6:]) // 删除6
	c = append(c[0:5], c[7:]) // 删除6,7
	c = c[0 : len(c)-1]       // 删除最后1个，类似pop
	c = c[0 : len(c)-3]       // 删除最后3个，类似pop

	// -----------------------------------------------------------
	// 遍历
	//
	for index, item := range c {
		fmt.Println(item)
	}

}
