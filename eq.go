package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Eq struct {
	id    string
	Field string
	Value string
}

func (e Eq) String() string {
	return e.Field + " = " + e.Value
}

func (e Eq) Subbed() string {
	return e.Field + " = " + e.Value
}

func (e Eq) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
