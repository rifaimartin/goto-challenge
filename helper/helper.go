package helper

import "github.com/go-playground/validator/v10"

//Response struct
type Response struct {
	Success 	bool    	   `json:"success"`
	Message string         `json:"message"`
	// Meta 	Meta     	   `json:"meta"`
	Data 	interface{}    `json:"data"`
}

type ResponseFail struct {
	Success 	bool    	   `json:"success"`
	Message string         `json:"message"`
	// Meta 	Meta     	   `json:"meta"`
	Error 	interface{}    `json:"error"`
}

// Meta Strucut
type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status" `
}

//APIResponse response data
func APIResponse(message string, code int, status bool, data interface{}) Response {
	// meta := Meta{
	// 	Message: message,
	// 	Code:    code,
	// }

	jsonResponse := Response{
		Success : status,
		Message : message,
		// Meta: meta,
		Data: data,
	}

	return jsonResponse
}

func APIResponseFail(message string, code int, status bool, data interface{}) ResponseFail {
	// meta := Meta{
	// 	Message: message,
	// 	Code:    code,
	// }

	var empity struct{}


	jsonResponse := ResponseFail{
		Success : status,
		Message : message,
		// Meta: meta,
		Error: empity,
	}

	return jsonResponse
}


//FormatValidationError for struct validation
func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}
