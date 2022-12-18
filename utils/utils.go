package utils

import (
	"github.com/eijiok/user-api/interfaces"
	"log"
	"net/http"
)

func RequestResponseErrorHandler(reqRespErr interfaces.RequestResponseErrorFunc) interfaces.RequestResponseFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		err := reqRespErr(writer, request)
		if err != nil {
			writer.WriteHeader(err.StatusCode)
			_ = WriteToJson(writer, err)
			log.Print(err.Error())
		}
	}
}
