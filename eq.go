package conbuilder

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Eq struct {
	id    string
	Field string
	Value string
	Typ   []string
}

func (e Eq) String() string {
	return e.Field + " = " + e.Value
}

func (e Eq) Subbed() string {
	return e.Field + " = " + subKey(e.id)
}

func (e Eq) AV() map[string]*dynamodb.AttributeValue {
	key := subKey(e.id)
	value := valueOfType(e.Value, e.Typ...)
	return map[string]*dynamodb.AttributeValue{key: &value}
}

func (cf ConditionFunc) Eq(field, value string, typ ...string) ConditionFunc {
	return func(id ...string) Condition {
		c := cf(id...)
		c.Clauses = append(c.Clauses, Eq{clauseID(c), field, value, typ})
		return c
	}
}
