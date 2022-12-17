package utils

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func ReadInJsonToStruct(reader io.Reader, pointer any) error {
	bytes, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, pointer)
	if err != nil {
		return err
	}

	return nil
}

func ReadParam(request *http.Request, param string) string {
	return mux.Vars(request)[param]
}
