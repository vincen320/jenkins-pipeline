package exception

type BadRequestErr struct {
	ErrorString string
}

func NewBadRequestErr(err string) BadRequestErr {
	return BadRequestErr{
		ErrorString: err,
	}
}

//Impelement error, n BadRequestErr nya memang tidak memiliki pointer
func (b BadRequestErr) Error() string {
	return b.ErrorString
}
