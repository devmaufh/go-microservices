package request_logs

import (
	"github.com/devmaufh/go_books/database/postgres/users_db"
	"github.com/devmaufh/go_books/utils/date_utils"
	"github.com/devmaufh/go_books/utils/errors"
)

const (
	queryInsertRequestLog = "INSERT INTO request_logs(timestamp,client_ip) VALUES(?,?) RETURNING ID"
)

func (RequestLog *RequestLog) Save() *errors.RestError {
	RequestLog.Timestamp = date_utils.GetNowString()
	result := users_db.Client.Raw(queryInsertRequestLog, RequestLog.Timestamp, RequestLog.ClientIp).Scan(&RequestLog)
	if result.Error != nil {
		return errors.NewInternalServerError(result.Error.Error())
	}
	return nil
}
