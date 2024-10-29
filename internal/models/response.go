package models

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Status      int         `json:"statusCode"`
	Data        interface{} `json:"data"`
	Message     string      `json:"message"`
	contentType string
	respWriter  http.ResponseWriter
}

func CreateResponse(rw http.ResponseWriter,
	codeHttp int,
	msg string,
	d interface{}) Response {
	return Response{
		Status:      codeHttp,
		respWriter:  rw,
		Data:        d,
		Message:     msg,
		contentType: "application/json",
	}
}

func (resp *Response) SendResult() {

	log.Println("[INFO]: Sending http result")
	resp.respWriter.Header().Set("Content-Type", resp.contentType)
	resp.respWriter.WriteHeader(resp.Status)

	output, _ := json.Marshal(&resp)
	fmt.Fprintln(resp.respWriter, string(output))
}
