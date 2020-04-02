package golang_basic

import (
	"fmt"
)

// 避坑好文：https://zhuanlan.zhihu.com/p/41418726

func MapOperation() {
	var ma map[int]rune

	ma := make(map[int32]rune)

	m := map[int]string{0: "0", 1: "1"} // 还可以这样初始化
	m[0] = "first"
	m[1] = "second"
	delete(m, 0)

	ms := make(map[int]map[int]string) //只是初始化最外层的map
	mv, ok := ms[0][0]                 //判断里面的map有没有初始化,如果没有,返回false
	if !ok {
		ms[0] = make(map[int]string) //现在初始化value的map
	}
	ms[0][0] = "赋值成功"
	fm.Println("ms is:", ms, " mv is:", mv)

	// 判断是否存在某个key
	if val, ok := dict["foo"]; ok {
		//do something here
	}

	// 遍历map
	// *不能保证多次遍历顺序一致
	for k, v := range m {
		fmt.Printf("key:%d, val:%s", k, v)
	}

	// 遍历map的key
	var keys []string
	for k := range m {
		keys = append(keys, k)
	}

	// Map的Value为结构体时候不能原地修改
	// see: https://www.jianshu.com/p/30e86473bdce

}
