
// c++:
// reinterpret_cast意图执行低级转型，实际动作（及结果）可能取决于编辑器，这也就表示它不可移植。
// ref: https://www.cnblogs.com/chenyangchun/p/6795923.html

// safe
func (s OtherTest) TestCkv(ctx *context.Context) error {
	var buffer bytes.Buffer
	b1 := []byte("@tfcm")
	buffer.Write(b1)
	key := make([]byte, 8, 8)
	binary.LittleEndian.PutUint64(key, uint64(10062))
	buffer.Write(key)
	var realkey string = buffer.String()
	return nil
}

//  unsafe!!! The following cast is unsafed; code are not portable
func GetExperimentKey(id uint64) string {
	key := "@ept" + string((*((*[unsafe.Sizeof(id)]byte)(unsafe.Pointer(&id))))[:])
	return key
}

