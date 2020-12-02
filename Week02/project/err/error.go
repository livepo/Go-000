package err

type Error struct {
	Msg string
}

func (e Error) Error() string {
	return e.Msg
}

func NewError(msg string) Error {
	return Error{Msg: msg}
}

var (
	Duplicate = NewError("Conflict")
	NotExist  = NewError("Not Found")
)
