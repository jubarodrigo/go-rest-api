package service

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/jubarodrigo/api-rest/database"
	"github.com/jubarodrigo/api-rest/model"
)

var people []model.Person

func CreatePerson(request *http.Request) error {


	client, err := database.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(database.DB).Collection(database.PERSON)

	_, err = collection.InsertOne(context.TODO(),json.NewDecoder(request.Body))
	if err != nil {
		return err
	}

	return nil

}

