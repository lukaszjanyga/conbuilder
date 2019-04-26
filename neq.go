package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type NEq struct {
	id    string
	Field string
	Value string
	Typ   []string
}

func (ne NEq) String() string {
	return ne.Field + " <> " + ne.Value
}

func (ne NEq) Subbed() string {
	return ne.Field + " <> " + subKey(ne.id)
}

func (ne NEq) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(ne.id)
	value := valueOfType(ne.Value, ne.Typ...)
	return map[string]*dynamodb.AttributeValue{
		key: &value,
	}
}
