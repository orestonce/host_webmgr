package ymdError

// 如果err不为nil，则触发panic
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
