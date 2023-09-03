package aws

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/fi3te/basicauth"
	"github.com/fi3te/sensor-am2302-data-store/pkg/config"
)

var cfg *config.Config
var db dynamodb.Client

func init() {
	var err error
	cfg, err = config.ReadConfig()
	if err != nil {
		panic(err)
	}
	db = createDynamoDbClient(context.TODO())
}

func HandleRequest(ctx context.Context, request events.LambdaFunctionURLRequest) (events.LambdaFunctionURLResponse, error) {
	authorized := basicauth.IsAuthorized(request.Headers, cfg.Username, cfg.Password)
	if !authorized {
		return responseWithStatus(http.StatusUnauthorized)
	}

	method := request.RequestContext.HTTP.Method
	if method == http.MethodPut {
		return handlePutRequest(context.Background(), request.Body)
	}
	return responseWithStatus(http.StatusBadRequest)
}

func handlePutRequest(ctx context.Context, requestBody string) (events.LambdaFunctionURLResponse, error) {
	var dataPoint DataPoint
	err := json.Unmarshal([]byte(requestBody), &dataPoint)
	if err != nil {
		return responseWithStatus(http.StatusBadRequest)
	}
	err = putItem(ctx, &db, cfg.TableName, dataPoint)
	if err != nil {
		log.Println(err)
		return responseWithStatus(http.StatusInternalServerError)
	}
	return responseWithStatus(http.StatusNoContent)
}

func responseWithStatus(code int) (events.LambdaFunctionURLResponse, error) {
	body, err := toJsonString(http.StatusText(code))
	return events.LambdaFunctionURLResponse{
		StatusCode: code,
		Body:       body,
	}, err
}

func toJsonString(v interface{}) (string, error) {
	b, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(b[:]), nil
}
