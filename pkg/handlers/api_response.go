package handlers

import (
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
	"github.com/vitorAzevedo09/go-serverless/pkg/user"
)

var ErrorMethodNotAllowed = "method not allowed"

type ErrorBody struct {
	ErrorMsg *string `json:"error,omitempty"`
}

func GetUser(req events.APIGatewayProxyRequest, tablename string, dynaClient dynamodbiface.DynamoDBAPI) (*events.APIGatewayProxyResponse, error) {
	email := req.QueryStringParameters["email"]
	if len(email) > 0 {
		result, err := user.FetchUser(email, tablename, dynaClient)
		if err != nil {
			return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
		}
		return apiResponse(http.StatusOK, result)
	}
	result, err := user.FetchUsers(tablename, dynaClient)
	if err != nil {
		return apiResponse(http.StatusBadRequest, ErrorBody{aws.String(err.Error())})
	}
	return apiResponse(http.StatusOK, result)
}

func CreateUser(req events.APIGatewayCustomAuthorizerRequest, tablename string, dynaClient dynamodbiface.DynamoDBAPI) {
}

func UpdateUser(req events.APIGatewayCustomAuthorizerRequest, tablename string, dynaClient dynamodbiface.DynamoDBAPI) {
}

func DeleteUser(req events.APIGatewayCustomAuthorizerRequest, tablename string, dynaClient dynamodbiface.DynamoDBAPI) {
}

func UnhandledUser() (*events.APIGatewayProxyResponse, error) {
	return apiResponse(http.StatusMethodNotAllowed, ErrorMethodNotAllowed)
}
