package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type GTE struct {
	id    string
	Field string
	Value string
	Typ   []string
}

func (gte GTE) String() string {
	return gte.Field + " >= " + gte.Value
}

func (gte GTE) Subbed() string {
	return gte.Field + " >= " + subKey(gte.id)
}

func (gte GTE) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(gte.id)
	value := valueOfType(gte.Value, gte.Typ...)
	return map[string]*dynamodb.AttributeValue{key: &value}
}

func (cf ConditionFunc) GTE(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, GTE{clauseID(c), field, value, typ})
		return c
	}
}
