package utils

type Utils struct {
}

func (u *Utils) ErrorHandler(err interface{}) {
	if err != nil {
		panic(err)
	}
}
