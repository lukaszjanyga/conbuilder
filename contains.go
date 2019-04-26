package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Contains struct {
	id      string
	Path    string
	Operand string
}

func (c Contains) String() string {
	return "contains(" + c.Path + "," + c.Operand + ")"
}

func (c Contains) Subbed() string {
	return "contains(" + c.Path + "," + c.Operand + ")"
}

func (c Contains) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
