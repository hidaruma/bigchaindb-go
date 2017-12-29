package backend


type BackEndError struct {
	Msg string
	Code int
}

func (err *BackEndError) Error() string {
	return fmt.Sprintf("err %s [code=%d]", err.Msg, err.Code)
}

type ConnectionError struct {
	Msg string
	Code int
}

func (err *ConnectionError) Error() string {
	
}

type OperationError struct {
	Msg string
	Code int
}

type DuplicateKeyError struct {
	Msg string
	Code int
}