package util

import (
	"errors"
	"poker/pkg/alert"
	service2 "poker/pkg/service"
)

func CheckErr(answer *service2.Answer, err error) {
	if err != nil {
		alert.Error(err)
	}
	if answer.Status != 0 {
		alert.Error(errors.New(answer.Msg))
	}
}
