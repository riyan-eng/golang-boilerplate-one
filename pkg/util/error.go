package util

func PanicIfNeeded(err interface{}) {
	if err != nil {
		panic(err)
	}
}

type ValidationError struct {
	Message string
}

func (validationError ValidationError) Error() string {
	return validationError.Message
}

type BadRequest struct {
	Message string
}

func (b BadRequest) Error() string {
	return b.Message
}
