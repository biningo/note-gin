package utils

func ErrReport(err error) {
	if err != nil {
		panic(err)
	}
}
