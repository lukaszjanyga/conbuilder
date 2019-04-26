package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type LTE struct {
	id    string
	Field string
	Value string
}

func (lte LTE) String() string {
	return lte.Field + " <= " + lte.Value
}

func (lte LTE) Subbed() string {
	return lte.Field + " <= " + lte.Value
}

func (lte LTE) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
