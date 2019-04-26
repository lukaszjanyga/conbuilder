package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type And struct {
	id string
}

func (a And) String() string {
	return "AND"
}

func (a And) Subbed() string {
	return a.String()
}

func (a And) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
