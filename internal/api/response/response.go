package response

type Response struct {
	Status     string `json:"status"`
	Error      string `json:"error,omitempty"`
	StatusCode int
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK() Response {
	return Response{
		Status: StatusOK,
	}
}

func Error(msg string, statusCode int) Response {
	return Response{
		Status:     StatusError,
		Error:      msg,
		StatusCode: statusCode,
	}
}
