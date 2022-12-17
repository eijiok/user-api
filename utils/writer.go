package utils

import (
	"encoding/json"
	"net/http"
)

func WriteToJson(writer http.ResponseWriter, object any) error {
	userJson, err := json.Marshal(object)
	if err != nil {
		return err
	}

	_, err = writer.Write(userJson)
	if err != nil {
		return err
	}

	return nil
}
