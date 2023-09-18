package errno

import (
	"errors"
	"fmt"
)

type Err struct {
	Code int64
	Msg  string
}

func (err Err) Error() string {
	return fmt.Sprintf("[%d] %s", err.Code, err.Msg)
}

func NewErr(errno int64, a ...any) error {
	return &Err{
		Code: errno,
		Msg:  fmt.Sprintf(ErrnoFormatMap[errno], a...),
	}
}

func SimpleErr(errno int64) error {
	return &Err{
		Code: errno,
		Msg:  ErrnoMap[errno],
	}
}

func IsErr(err error, errno int64) bool {
	var e *Err
	if errors.As(err, &e) {
		return e.Code == errno
	}
	return false
}

func Errno(err error) int64 {
	var e *Err
	if errors.As(err, &e) {
		return e.Code
	}
	return ErrUnknown
}
