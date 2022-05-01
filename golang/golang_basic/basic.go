package main

import "fmt"

// ===========================================================================
//                             variable initialization
// ===========================================================================

type People struct {
	name  string
	child *People
}

relation := &People{
	name: "爷爷",
	child: &People{
		name: "爸爸",
		child: &People{
			name: "我",
		},
	},
}


// 类型定义都用type开头

type typeInterface interface {
	name() string
}

type typeStructure struct {
	name string
}

type typeAlias *string

// type impliment interface

func (t *typeAlias) name() string {
	return "typeAlias implement interface"
}

func (c typeStructure) name() string {
	return "typeStructure implement interface"
}

// ========================================================================
//                               类型转换
// ========================================================================
func typeConv() {

	// 使用type(variable)的方式进行类型转换
	var a int32 = 10
	var b int64 = int64(a)
	var c float32 = 12.3
	var d float64 = float64(c)

	var p *int = &a
	var c *int64
	c = (*int64)(p)

	var foo string = "lunzhoufei"
	var bar []byte = []byte(foo)
	var feliz string = string(bar)
}

// ========================================================================
//                         类型判定
// ========================================================================

func typeAssert() {
	var foo interface{}

	foo = "lunzhoufei"
	bar = foo.(*String)  // will panic
	feliz = foo.(String) // ok

	if value, ok = foo.(String); !ok {
		fmt.Println("foo is not the type of string")
	}
}

// ========================================================================
//                      for
// ========================================================================
func main() {
	fmt.Println("vim-go")


	// switch case
	switch strings.ToUpper(command) {
	case "SET", "ADD":
		store, extends, err := parseDoArg(data...)
		if err != nil {
			return nil, err
		}
		switch command {
		case "SET":
			c.DoSet(ctx, key, store, extends...)
		case "ADD":
			c.DoAdd(ctx, key, store, extends...)
		}
	case "GET":
		c.DoGet(ctx, key)
	case "DEL":
		c.DoDel(ctx, key)
	default:
		return nil, ErrorParamInvalid
	}

	a := []string{}
	a = append(a, "hello")
	a = append(a, ", ")
	a = append(a, "world")
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(&a[i])
	}
	fmt.Println(a)

	for _, aa := range a {
		aa += "@"
		fmt.Println(&aa)
	}
	fmt.Println(a)

	for i := 0; i < len(a); i++ {
		a[i] += "@"
		fmt.Println(&a[i])
	}
	fmt.Println(a)

}
