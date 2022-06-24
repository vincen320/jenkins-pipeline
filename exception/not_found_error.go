package exception

type NotFoundErr struct {
	ErrorString string
}

func NewNotFoundErr(err string) NotFoundErr {
	return NotFoundErr{
		ErrorString: err,
	}
}

//Impelement error, n NotFoundErr nya memang tidak memiliki pointer
func (n NotFoundErr) Error() string {
	return n.ErrorString
}
