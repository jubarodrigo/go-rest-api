package service

import (
	"context"
	"encoding/json"
	"net/http"
	"github.com/jubarodrigo/api-rest/database"
	"github.com/jubarodrigo/api-rest/model"
)

var people model.Person

func CreatePerson(request *http.Request) interface{} {

	envs := database.LoadEnv()


	client, err := database.GetMongoClient()
	if err != nil {
		return err
	}

	collection := client.Database(envs["MG_DATABASE"]).Collection(database.PERSON)

	defer request.Body.Close()

	if err := json.NewDecoder(request.Body).Decode(&people); err != nil {
		return err
	}
	
	result, err := collection.InsertOne(context.TODO(),people)
	if err != nil {
		return err
	}

	return result


}

