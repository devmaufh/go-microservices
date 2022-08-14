package services

import (
	"github.com/devmaufh/go_books/domain/request_logs"
	"github.com/devmaufh/go_books/utils/errors"
)

func CreateRequestLog(requestLog request_logs.RequestLog) (*request_logs.RequestLog, *errors.RestError) {

	if err := requestLog.Save(); err != nil {
		return nil, errors.NewInternalServerError("The request log cannot be stored.")
	}

	return &requestLog, nil
}