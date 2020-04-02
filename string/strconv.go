package test_strconv


// ============================================================================
//								integer to string
// ============================================================================
// common method
var n uint32 = 42
str := fmt.Sprint(n)
fmt.println(str)

// int => string
var n int = 22
str := strconv.Itoa(n)

// uint32 => string
var n uint32 = 42
str := strconv.FormatUint(uint64(n), 10)

// uint64 => string
var n uint64 = 42
str := strconv.FormatUint(n, 10)

// int32 => string
var n uint32 = 42
str := strconv.FormatInt(Int64(n), 10)

// int64 => string
var n uint64 = 42
str := strconv.FormatInt(n, 10)


