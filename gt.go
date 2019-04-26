package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GT struct {
	id    string
	Field string
	Value string
}

func (gt GT) String() string {
	return gt.Field + " > " + gt.Value
}

func (gt GT) Subbed() string {
	return gt.Field + " > " + gt.Value
}

func (gt GT) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
