package utils

import (
	"encoding/json"
	"github.com/eijiok/user-api/errors"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func IdToObjectId(id string) (primitive.ObjectID, *errors.HttpError) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, errors.NewHttpError(http.StatusBadRequest, "Error to serialize the id", &err)
	}
	return objectId, nil
}
