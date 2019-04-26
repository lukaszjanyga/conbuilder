package conbuilder

import (
	"fmt"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strconv"
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
	case "B", "b":
		val, _ := strconv.ParseBool(value)
		return dynamodb.AttributeValue{BOOL: &val}
	}
	return dynamodb.AttributeValue{S: &value}
}

func firstSafe(s []string) string {
	o := ""
	if len(s) > 0 {
		o = s[0]
	}
	return o
}
