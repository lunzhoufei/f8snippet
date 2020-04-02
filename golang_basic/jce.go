
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
	return nil, errors.New("Set data type no supported")

}

