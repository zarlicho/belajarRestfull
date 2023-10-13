package helper

func PanicErrorIf(err error) {
	if err != nil {
		panic(err)
	}
}
