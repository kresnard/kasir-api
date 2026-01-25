package helper

type ResponseSuccess struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SuccessResponse(message string, data interface{}) (res ResponseSuccess) {

	res.Status = "OK"
	res.Message = message
	res.Data = data

	return
}
