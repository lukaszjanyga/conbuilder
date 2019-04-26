package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"strings"
)

type In struct {
	id       string
	Field    string
	Operands []string
}

func (i In) String() string {
	return i.Field + " IN (" + strings.Join(i.Operands, ", ") + ")"
}

func (i In) Subbed() string {
	return i.Field + " IN (" + strings.Join(i.Operands, ", ") + ")"
}

func (i In) AV() map[string]*dynamodb.AttributeValue {
	return nil
}
