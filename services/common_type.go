package services

const (
	invalidError int = iota + 1
	systemError
)

type ServiceError struct {
	Code int
	Err  error
	Msg  string
}

func setInvalidReturn(err error, userMsg string) ServiceError {
	return ServiceError{
		Code: invalidError,
		Err:  err,
		Msg:  userMsg,
	}
}

func setSystemErrReturn(err error) ServiceError {
	return ServiceError{
		Code: systemError,
		Err:  err,
	}
}

func (se ServiceError) IsInvalidErr() bool {
	return se.Code == invalidError
}

func (se ServiceError) IsSystemErr() bool {
	return se.Code == systemError
}
