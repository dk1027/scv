package shared

func CHECK_ERR(err error) {
	if err != nil {
		panic(err)
	}
}

func CHECK_OK(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}
