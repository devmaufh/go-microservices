package request_logs

type RequestLog struct {
	Id        int64  `json:"id"`
	Timestamp string `json:"timestamp"`
	ClientIp  string `json:"client_ip"`
}
