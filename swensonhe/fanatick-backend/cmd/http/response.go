package main

import "strings"

const (
	CommonStatus200ResponseCode             = "success"
	CommonStatus200GETCallResponseMessage   = "Details fetched successfully"
	CommonStatus200POSTCallResponseMessage  = "Successfully Created"
	CommonStatus200PATCHCallResponseMessage = "Successfully Updated"
)

type CommonStatus200Response struct {
	Data    interface{} `json:"data"`
	Status  int         `json:"status" `
	Code    string      `json:"code"`
	Message string      `json:"message"`
}

// newCommonSuccessResponse returns an common response.
func newCommonSuccessResponse(data interface{}, messages ...string) CommonStatus200Response {
	return CommonStatus200Response{
		Data:    data,
		Status:  200,
		Code:    CommonStatus200ResponseCode,
		Message: strings.Join(messages, " "),
	}
}
