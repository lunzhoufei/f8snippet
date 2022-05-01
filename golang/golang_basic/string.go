package golang_basic

import (
	"strconv"
)

func strTruncate() {
	var s string
	s = "a" + "b"

	// string2
	s.Atoi()
	s.ParseInt()
	s.ParseUint()
	s.ParseBool()
	s.ParseFloat()

	// 2string
	s.Itoa()
	s.FormatInt()
	s.FormatUint()
	s.FormatBool()
	s.FormatFloat()
}

func strSplit() {
	s := "Hello, 世界! Hello!"

	// Split
	ss := strings.Split(s, " ")
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.Split(s, ", ")
	fmt.Printf("%q\n", ss) // ["Hello" "世界! Hello!"]
	ss = strings.Split(s, "")
	fmt.Printf("%q\n", ss) // 单个字符列表

	// SplitAfter
	ss = strings.SplitAfter(s, " ")
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfter(s, ", ")
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfter(s, "")
	fmt.Printf("%q\n", ss) // 单个字符列表

	// SplitN
	ss = strings.SplitN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello," "世界! Hello!"]
	ss = strings.SplitN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.SplitN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]

	// SplitAfterN
	ss = strings.SplitAfterN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfterN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfterN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]

	// Fields 以连续的空白字符为分隔符，将 s 切分成多个子串，结果中不包含空白字符本身
	// 空白字符有：\t, \n, \v, \f, \r, ' ', U+0085 (NEL), U+00A0 (NBSP)
	// 如果 s 中只包含空白字符，则返回一个空列表
	ss := strings.Fields(s)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]

	// FieldsFunc 以一个或多个满足 f(rune) 的字符为分隔符，
	// 将 s 切分成多个子串，结果中不包含分隔符本身。
	// 如果 s 中没有满足 f(rune) 的字符，则返回一个空列表。
	func isSlash(r rune) bool {
		return r == '\\' || r == '/'
	}
	s = "C:\\Windows\\System32\\FileName"
	ss = strings.FieldsFunc(s, isSlash)
	fmt.Printf("%q\n", ss) // ["C:" "Windows" "System32" "FileName"]

}

func strContainCheck() {
	s := "Hello,世界!!!!!"

	// Contain
	b := strings.Contains(s, "!!")
	fmt.Println(b) // true
	b = strings.Contains(s, "!?")
	fmt.Println(b) // false
	b = strings.Contains(s, "")
	fmt.Println(b) // true

	b = strings.ContainsAny(s, "abc")
	fmt.Println(b) // false
	b = strings.ContainsAny(s, "def")
	fmt.Println(b) // true

	b = strings.ContainsRune(s, '\n')
	fmt.Println(b) // false
	b = strings.ContainsRune(s, '界')
	fmt.Println(b) // true
	b = strings.ContainsRune(s, 0)
	fmt.Println(b) // false
}

func strContainCheck() {
	s := "Hello,世界!!!!!"

	// Index
	i := strings.Index(s, "h")
	fmt.Println(i) // -1
	i = strings.Index(s, "!")
	fmt.Println(i) // 12
	i = strings.Index(s, "")
	fmt.Println(i) // 0

	i = strings.LastIndex(s, "h")
	fmt.Println(i) // -1
	i = strings.LastIndex(s, "H")
	fmt.Println(i) // 14
	i = strings.LastIndex(s, "")
	fmt.Println(i) // 20

	i = strings.IndexRune(s, '\n')
	fmt.Println(i) // -1
	i = strings.IndexRune(s, '界')
	fmt.Println(i) // 9
	i = strings.IndexRune(s, 0)
	fmt.Println(i) // -1

	i = strings.IndexAny(s, "abc")
	fmt.Println(i) // -1
	i = strings.IndexAny(s, "dof")
	fmt.Println(i) // 1
	i = strings.IndexAny(s, "")
	fmt.Println(i) // -1

	i = strings.LastIndexAny(s, "abc")
	fmt.Println(i) // -1
	i = strings.LastIndexAny(s, "def")
	fmt.Println(i) // 15
	i = strings.LastIndexAny(s, "")
	fmt.Println(i) // -1
}
