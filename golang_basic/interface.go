package main

import "fmt"

// interface receiver是value还是pointer => https://sanyuesha.com/2017/07/22/how-to-understand-go-interface/

// 类型定义都用type开头

// 空interface作用很多类似于指针(*void)
type emptyInterface interface{}

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

	// switch case
	var data interface{}
	switch value := data.(type) {
	case []byte:
		return value, nil
	case string:
		return []byte(value), nil
	case proto.Message:
		return proto.Marshal(value)
	case gojce.Message:
		return gojce.Marshal(value)
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return []byte(fmt.Sprintf("%d", value)), nil
	}
}

// ========================================================================
//                      for
// ========================================================================
func main() {
	fmt.Println("vim-go")

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
