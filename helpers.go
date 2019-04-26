package conbuilder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func clauseID(parent Condition) string {
	return fmt.Sprintf("%s_%d", parent.id, len(parent.Clauses))
}

func subKey(id string) string {
	return ":" + id
}

func valueOfType(value string, typ ...string) dynamodb.AttributeValue {
	if len(typ) == 0 {
		return dynamodb.AttributeValue{S: &value}
	}
	switch typ[0] {
	case "N", "n":
		return dynamodb.AttributeValue{N: &value}
	}
	return dynamodb.AttributeValue{S: &value}
}
