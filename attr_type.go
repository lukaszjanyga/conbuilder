package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type AttrType struct {
	id   string
	Path string
	Type string
}

func (at AttrType) String() string {
	return "attribute_type(" + at.Path + "," + at.Type + ")"
}

func (at AttrType) Subbed() string {
	return "attribute_type(" + at.Path + "," + subKey(at.id) + ")"
}

func (at AttrType) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(at.id)
	value := valueOfType(at.Type)
	return map[string]*dynamodb.AttributeValue{key: &value}
}

func (cf ConditionFunc) AttrType(path, attrType string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, AttrType{clauseID(c), path, attrType})
		return c
	}
}
