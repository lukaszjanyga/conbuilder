package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Or struct {
	id string
}

func (o Or) String() string {
	return "OR"
}

func (o Or) Subbed() string {
	return o.String()
}

func (o Or) AV() map[string]*dynamodb.AttributeValue {
	return map[string]*dynamodb.AttributeValue{}
}
