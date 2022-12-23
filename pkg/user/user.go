package user

import (
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbiface"
)

func FetchUser(email string, tablename string, dynaClient dynamodbiface.DynamoDBAPI) {}

func FetchUsers(tablename string, dynaClient dynamodbiface.DynamoDBAPI) {}

func CreateUser() {}

func UpdateUser() {}

func DeleteUser() error {}
